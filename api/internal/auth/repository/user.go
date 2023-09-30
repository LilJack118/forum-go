package repository

import "go.mongodb.org/mongo-driver/mongo"

type authRepository struct {
	db *mongo.Client
}

func NewAuthRepository(db *mongo.Client) *authRepository {
	return &authRepository{db}
}
