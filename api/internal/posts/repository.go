package posts

import "forum/api/internal/models"

type PostsRepository interface {
	CreatePost(post *models.Post) error
	GetPost(id string) (*models.Post, int, error)
	UpdatePost(id string, uid string, fields *models.PostEditableFields) (int, error)
}
