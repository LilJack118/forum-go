package models

import (
	"forum/api/pkg/validator"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	PostWithoutContent
	Content string `json:"content" bson:"content" validate:"required,lte=10000"`
}

func (post *Post) Validate() error {
	return validator.ValidateStruct(post)
}

func (post *Post) PrepareCreate() {
	post.ID = uuid.New()
	post.CreateAt = time.Now().UTC()

	post.Title = strings.TrimSpace(post.Title)
}

func (post *Post) SetUID(uidString string) error {

	uid, err := uuid.Parse(uidString)
	if err != nil {
		return err
	}

	post.UID = uid
	return nil
}

type PostWithoutContent struct {
	ID       uuid.UUID `json:"id" bson:"id" validate:"omitempty"`
	UID      uuid.UUID `json:"uid" bson:"uid" validate:"required"`
	Title    string    `json:"title" bson:"title" validate:"required,lte=300"`
	CreateAt time.Time `json:"created_at" bson:"created_at" validate:"omitempty"`
}
