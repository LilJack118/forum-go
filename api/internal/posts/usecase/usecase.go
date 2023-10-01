package usecase

import "forum/api/internal/posts"

type postsUseCase struct {
	repo posts.PostsRepository
}

func NewPostsUseCase(repo posts.PostsRepository) *postsUseCase {
	return &postsUseCase{repo}
}
