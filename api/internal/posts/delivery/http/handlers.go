package http

import (
	"encoding/json"
	"forum/api/internal/models"
	"forum/api/internal/posts"
	"forum/api/pkg/httpErrors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type postHandlers struct {
	uc posts.PostsUseCase
}

func NewPostHandlers(uc posts.PostsUseCase) *postHandlers {
	return &postHandlers{uc}
}

func (h *postHandlers) CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPost models.Post

	vars := mux.Vars(r)

	if err := json.NewDecoder(r.Body).Decode(&newPost); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	newPost.SetUID(vars["uid"])

	newPostResponse, code, err := h.uc.CreatePost(&newPost)
	if err != nil {
		httpErrors.JSONError(w, err.Error(), code)
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newPostResponse); err != nil {
		httpErrors.JSONError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *postHandlers) GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	post_id := vars["id"]

	post, code, err := h.uc.GetPost(post_id)
	if err != nil {
		httpErrors.JSONError(w, err.Error(), code)
	}

	if err := json.NewEncoder(w).Encode(post); err != nil {
		log.Print(err)
		httpErrors.JSONError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *postHandlers) UpdatePost(w http.ResponseWriter, r *http.Request) {

}

func (h *postHandlers) DeletePost(w http.ResponseWriter, r *http.Request) {

}

func (h *postHandlers) ListPosts(w http.ResponseWriter, r *http.Request) {

}
