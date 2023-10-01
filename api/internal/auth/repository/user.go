package repository

import (
	"context"
	"errors"
	"forum/api/internal/models"
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (repo *authRepository) getUser(filter bson.D) (*models.User, error) {
	var user *models.User
	collection := repo.getCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *authRepository) GetUserByEmail(email string) (*models.User, error) {
	filter := bson.D{primitive.E{Key: "email", Value: email}}
	return repo.getUser(filter)
}

func (repo *authRepository) GetUserByID(id string) (*models.User, error) {

	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{primitive.E{Key: "id", Value: uuid}}
	return repo.getUser(filter)
}

func (repo *authRepository) UpdateUser(id string, fields *models.UserEditableFields) (int, error) {

	if err := fields.Validate(); err != nil {
		return http.StatusBadRequest, err
	}

	collection := repo.getCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	bfields, err := bson.Marshal(fields)
	if err != nil {
		return http.StatusBadRequest, err
	}

	var updateFields bson.M
	if err := bson.Unmarshal(bfields, &updateFields); err != nil {
		return http.StatusBadRequest, err
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return http.StatusBadRequest, err
	}

	filter := bson.D{primitive.E{Key: "id", Value: uid}}
	update := bson.D{primitive.E{Key: "$set", Value: updateFields}}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return http.StatusBadRequest, err
	}

	if result.MatchedCount == 0 {
		return http.StatusNotFound, errors.New("user with specified id does not exist")
	}

	return 0, nil
}

func (repo *authRepository) DeleteUser(id string) (int, error) {

	collection := repo.getCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	uid, err := uuid.Parse(id)
	if err != nil {
		return http.StatusBadRequest, err
	}

	result, err := collection.DeleteOne(ctx, bson.D{primitive.E{Key: "id", Value: uid}})
	if err != nil {
		return http.StatusBadRequest, err
	}

	if result.DeletedCount == 0 {
		return http.StatusNotFound, errors.New("user with specified id does not exist")
	}

	return 0, nil
}
