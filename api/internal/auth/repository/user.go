package repository

import (
	"context"
	"forum/api/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type authRepository struct {
	db *mongo.Database
}

func NewAuthRepository(db *mongo.Database) *authRepository {
	return &authRepository{db}
}

func (repo *authRepository) getCollection() *mongo.Collection {
	return repo.db.Collection("users")
}

func (repo *authRepository) CreateUser(user *models.User) error {
	collection := repo.getCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	buser, err := bson.Marshal(user)
	if err != nil {
		return err
	}

	if _, err := collection.InsertOne(ctx, buser); err != nil {
		return err
	}

	return nil
}
