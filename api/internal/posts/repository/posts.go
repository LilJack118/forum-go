package repository

import (
	"context"
	"errors"
	"forum/api/internal/models"
	"forum/api/pkg/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type postsRepository struct {
	db *mongo.Database
}

func NewPostsRepository(db *mongo.Database) *postsRepository {
	return &postsRepository{db}
}

func (repo *postsRepository) getCollection() *mongo.Collection {
	return repo.db.Collection("posts")
}

func (repo *postsRepository) CreatePost(post *models.Post) error {
	collection := repo.getCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bpost, err := bson.Marshal(post)
	if err != nil {
		return err
	}

	if _, err := collection.InsertOne(ctx, bpost); err != nil {
		return err
	}

	return nil
}

func (repo *postsRepository) GetPost(id string) (*models.Post, int, error) {
	post_id, err := uuid.Parse(id)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	var post *models.Post
	collection := repo.getCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "id", Value: post_id}}
	if err := collection.FindOne(ctx, filter).Decode(&post); err != nil {
		return nil, http.StatusBadRequest, err
	}

	return post, 0, nil
}

func (repo *postsRepository) UpdatePost(id_s string, uid_s string, fields *models.PostEditableFields) (int, error) {
	id, err := uuid.Parse(id_s)
	if err != nil {
		return http.StatusBadRequest, err
	}

	uid, err := uuid.Parse(uid_s)
	if err != nil {
		return http.StatusBadRequest, err
	}

	collection := repo.getCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bfields, err := utils.StructToBson(fields)
	if err != nil {
		return http.StatusBadRequest, err
	}

	filter := bson.D{primitive.E{Key: "id", Value: id}, primitive.E{Key: "uid", Value: uid}}
	update := bson.D{primitive.E{Key: "$set", Value: bfields}}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return http.StatusBadRequest, err
	}

	if result.MatchedCount == 0 {
		return http.StatusNotFound, errors.New("post with specified id and user id does not exist")
	}

	return 0, nil

}

func (repo *postsRepository) DeletePost(id_s string, uid_s string) (int, error) {
	id, err := uuid.Parse(id_s)
	if err != nil {
		return http.StatusBadRequest, err
	}

	uid, err := uuid.Parse(uid_s)
	if err != nil {
		return http.StatusBadRequest, err
	}

	collection := repo.getCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "id", Value: id}, primitive.E{Key: "uid", Value: uid}}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return http.StatusBadRequest, err
	}

	if result.DeletedCount == 0 {
		return http.StatusNotFound, errors.New("post with specified id and user id does not exist")
	}

	return 0, nil
}

func (repo *postsRepository) countPosts(filter bson.D) (int, error) {
	collection := repo.getCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if filter == nil {
		filter = bson.D{}
	}

	num, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return int(num), err
}

func (repo *postsRepository) getPosts(filter bson.D, page int, limit int) (*[]models.Post, error) {
	collection := repo.getCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// aggregate
	aggregateStage := bson.D{
		primitive.E{Key: "$project", Value: bson.D{
			primitive.E{Key: "id", Value: "$id"},
			primitive.E{Key: "uid", Value: "$uid"},
			primitive.E{Key: "title", Value: "$title"},
			primitive.E{Key: "created_at", Value: "$created_at"},
			// get only first 60 characters of content
			primitive.E{
				Key: "content",
				Value: bson.D{primitive.E{
					Key:   "$substrBytes",
					Value: bson.A{"$content", 0, 60},
				}},
			},
		}},
	}

	skipStage := bson.D{
		primitive.E{Key: "$skip", Value: (page - 1) * limit},
	}

	limitStage := bson.D{
		primitive.E{Key: "$limit", Value: limit},
	}

	sortStage := bson.D{primitive.E{
		Key:   "$sort",
		Value: bson.D{primitive.E{Key: "created_at", Value: -1}},
	}}

	pipeline := mongo.Pipeline{sortStage, skipStage, limitStage, aggregateStage}

	if filter != nil {
		matchStage := bson.D{
			primitive.E{Key: "$match", Value: filter},
		}
		pipeline = append([]bson.D{matchStage}, pipeline...)
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var results []models.Post
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return &results, nil
}

func (repo *postsRepository) listPostsWithFilter(filter bson.D, page int, limit int) (*models.PostsPage, error) {

	posts, err := repo.getPosts(filter, page, limit)
	if err != nil {
		return nil, err
	}

	postsNum, err := repo.countPosts(filter)
	if err != nil {
		return nil, err
	}

	postsPage := &models.PostsPage{
		Page:     page,
		Limit:    limit,
		PostsNum: postsNum,
		Posts:    posts,
	}

	return postsPage, nil
}

func (repo *postsRepository) ListPosts(page int, limit int) (*models.PostsPage, error) {
	return repo.listPostsWithFilter(nil, page, limit)
}

func (repo *postsRepository) ListUserPosts(uid uuid.UUID, page int, limit int) (*models.PostsPage, error) {
	filter := bson.D{primitive.E{Key: "uid", Value: uid}}
	return repo.listPostsWithFilter(filter, page, limit)
}
