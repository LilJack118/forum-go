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

	createdPost, _ := post.WithoutContent()

	return createdPost, 0, nil
}

func (uc *postsUseCase) GetPost(id string) (*models.Post, int, error) {
	return uc.repo.GetPost(id)
}

func (uc *postsUseCase) UpdatePost(id string, uid string, fields *models.PostEditableFields) (int, error) {
	return uc.repo.UpdatePost(id, uid, fields)
}

func (uc *postsUseCase) DeletePost(id string, uid string) (int, error) {
	return uc.repo.DeletePost(id, uid)
}
