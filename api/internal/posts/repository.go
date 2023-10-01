package posts

import "forum/api/internal/models"

type PostsRepository interface {
	CreatePost(post *models.Post) error
}
