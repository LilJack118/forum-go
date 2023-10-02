package posts

import (
	"forum/api/internal/models"

	"github.com/google/uuid"
)

type PostsRepository interface {
	CreatePost(post *models.Post) error
	GetPost(id string) (*models.Post, int, error)
	UpdatePost(id string, uid string, fields *models.PostEditableFields) (int, error)
	DeletePost(id string, uid string) (int, error)
	ListPosts(page int, limit int) (*models.PostsPage, error)
	ListUserPosts(uid uuid.UUID, pageS int, limitS int) (*models.PostsPage, error)
}
