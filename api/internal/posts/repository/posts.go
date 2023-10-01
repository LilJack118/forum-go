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
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (repo *postsRepository) applyPagination(page int, limit int, opts *options.FindOptions) {
	opts.SetSkip(int64((page - 1) * limit))
	opts.SetLimit(int64(limit))
}

func (repo *postsRepository) countPosts(filter bson.D) (int, error) {
	collection := repo.getCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	num, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return int(num), err
}

func (repo *postsRepository) getPosts(filter bson.D, opts *options.FindOptions) (*[]models.Post, error) {
	collection := repo.getCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var results []models.Post
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return &results, nil
}

func (repo *postsRepository) ListPosts(page int, limit int) (*models.PostsPage, error) {
	opts := options.Find().SetSort(primitive.E{Key: "created_at", Value: -1})
	repo.applyPagination(page, limit, opts)

	posts, err := repo.getPosts(bson.D{}, opts)
	if err != nil {
		return nil, err
	}

	postsNum, err := repo.countPosts(bson.D{})
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
