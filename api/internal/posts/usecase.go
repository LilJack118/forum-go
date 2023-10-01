package posts

import "forum/api/internal/models"

type PostsUseCase interface {
	CreatePost(post *models.Post) (*models.PostWithoutContent, int, error)
}
