package repository

import "go.mongodb.org/mongo-driver/mongo"

type PostsRepository struct {
	db *mongo.Database
}

func NewPostsRepository(db *mongo.Database) *PostsRepository {
	return &PostsRepository{db}
}
