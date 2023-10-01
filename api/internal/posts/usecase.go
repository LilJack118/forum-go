package posts

import "forum/api/internal/models"

type PostsUseCase interface {
	CreatePost(post *models.Post) (*models.PostWithoutContent, int, error)
	GetPost(id string) (*models.Post, int, error)
	UpdatePost(id string, uid string, fields *models.PostEditableFields) (int, error)
	DeletePost(id string, uid string) (int, error)
	ListPosts(pageS string, limitS string) (*models.PostsPage, error)
}
