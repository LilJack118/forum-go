package main

import (
	"forum/api/src/handlers"
	"forum/api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router = router.PathPrefix("/api").Subrouter()
	router.Use(middlewares.DefaultMiddleware)

	// auth endpoints
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	http.ListenAndServe(":80", router)
}
