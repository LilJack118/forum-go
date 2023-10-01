package models

import (
	"forum/api/pkg/validator"

	"github.com/google/uuid"
)

type Post struct {
	ID       uuid.UUID `json:"id" bson:"id" validate:"omitempty"`
	UID      uuid.UUID `json:"uid" bson:"uid" validate:"required"`
	Title    string    `json:"title" bson:"title" validate:"required,lte=300"`
	Content  string    `json:"content" bson:"content" validate:"required"`
	CreateAt string    `json:"created_at" bson:"created_at" validate:"omitempty"`
}

func (post *Post) Validate() error {
	return validator.ValidateStruct(post)
}

type PostWithoutContent struct {
	ID       uuid.UUID `json:"id" bson:"id" validate:"omitempty"`
	UID      uuid.UUID `json:"uid" bson:"uid" validate:"required"`
	Title    string    `json:"title" bson:"title" validate:"required,lte=300"`
	CreateAt string    `json:"created_at" bson:"created_at" validate:"omitempty"`
}
