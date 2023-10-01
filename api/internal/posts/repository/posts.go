package repository

import (
	"context"
	"forum/api/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
