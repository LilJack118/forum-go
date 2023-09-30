package repository

import (
	"forum/api/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type authRepository struct {
	db *mongo.Database
}

func NewAuthRepository(db *mongo.Database) *authRepository {
	return &authRepository{db}
}

func (repo *authRepository) CreateUser(user *models.User) error {

	return nil
}
