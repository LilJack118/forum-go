package repository

import (
	"context"
	"forum/api/internal/models"
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
