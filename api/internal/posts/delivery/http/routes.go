package http

import (
	"forum/api/internal/posts"

	"github.com/gorilla/mux"
)

func RegisterPostHandlers(r *mux.Router, u posts.PostsUseCase) {
	handler := NewPostHandlers(u)

	r.HandleFunc("/post", handler.CreatePost).Methods("POST")
	r.HandleFunc("/post/{id}", handler.GetPost).Methods("GET")
	r.HandleFunc("/post/{id}", handler.UpdatePost).Methods("PATCH")
	r.HandleFunc("/post/{id}", handler.DeletePost).Methods("DELETE")
	r.HandleFunc("/posts", handler.ListPosts).Methods("GET")
	r.HandleFunc("/posts/my", handler.ListMyPosts).Methods("GET")
}
