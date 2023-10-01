package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type accountRepository struct {
	db *mongo.Database
}

func NewAccountRepository(db *mongo.Database) *accountRepository {
	return &accountRepository{db}
}
