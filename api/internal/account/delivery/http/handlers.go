package http

import (
	"encoding/json"
	"forum/api/internal/account"
	"forum/api/internal/models"
	"forum/api/pkg/httpErrors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type accountHandlers struct {
	uc account.AccountUseCase
}

func NewAccountHandlers(u account.AccountUseCase) *accountHandlers {
	return &accountHandlers{u}
}

func (h *accountHandlers) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["uid"]

	account, code, err := h.uc.GetUserAccount(id)
	if err != nil {
		httpErrors.JSONError(w, err.Error(), code)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(account); err != nil {
		log.Print(err)
		httpErrors.JSONError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *accountHandlers) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["uid"]

	var input *models.UserEditableFields

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		httpErrors.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	code, err := h.uc.UpdateUserAccount(id, input)
	if err != nil {
		httpErrors.JSONError(w, err.Error(), code)
		return
	}
}

func (h *accountHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["uid"]

	code, err := h.uc.DeleteUserAccount(id)
	if err != nil {
		httpErrors.JSONError(w, err.Error(), code)
		return
	}
}
