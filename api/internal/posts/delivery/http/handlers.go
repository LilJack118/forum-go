package http

import (
	"forum/api/internal/posts"
	"net/http"
)

type postHandlers struct {
	uc posts.PostsUseCase
}

func NewPostHandlers(uc posts.PostsUseCase) *postHandlers {
	return &postHandlers{uc}
}

func (h *postHandlers) CreatePost(w http.ResponseWriter, r *http.Request) {

}

func (h *postHandlers) GetPost(w http.ResponseWriter, r *http.Request) {

}

func (h *postHandlers) UpdatePost(w http.ResponseWriter, r *http.Request) {

}

func (h *postHandlers) DeletePost(w http.ResponseWriter, r *http.Request) {

}

func (h *postHandlers) ListPosts(w http.ResponseWriter, r *http.Request) {

}
