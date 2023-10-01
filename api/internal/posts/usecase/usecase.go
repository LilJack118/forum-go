package usecase

import (
	"forum/api/internal/models"
	"forum/api/internal/posts"
	"net/http"
)

type postsUseCase struct {
	repo posts.PostsRepository
}

func NewPostsUseCase(repo posts.PostsRepository) *postsUseCase {
	return &postsUseCase{repo}
}

func (uc *postsUseCase) CreatePost(post *models.Post) (*models.PostWithoutContent, int, error) {

	if err := post.Validate(); err != nil {
		return nil, http.StatusBadRequest, err
	}

	post.PrepareCreate()

	if err := uc.repo.CreatePost(post); err != nil {
		return nil, http.StatusBadRequest, err
	}

	return &post.PostWithoutContent, 0, nil
}