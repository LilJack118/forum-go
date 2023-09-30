package server

import (
	"context"
	"forum/api/internal/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	authhttp "forum/api/internal/auth/delivery/http"
	authrepo "forum/api/internal/auth/repository"
	authuc "forum/api/internal/auth/usecase"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

// Server struct
type Server struct {
	httpServer *http.Server
	db         *mongo.Client
}

// NewServer New Server constructor
func NewServer(db *mongo.Client) *Server {
	return &Server{db: db}
}

func (s *Server) Run(port string) error {

	router := mux.NewRouter()

	// init repositiories
	auth_repo := authrepo.NewAuthRepository(s.db)

	// init use cases
	auth_uc := authuc.NewAuthUseCase(auth_repo)

	// register routes
	auth_router := router.PathPrefix("/auth").Subrouter()
	auth_router.Use(middleware.DefaultMiddleware)
	authhttp.RegisterRoutes(auth_router, auth_uc)

	s.httpServer = &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// run httpserver in goroutine
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1) // create channel to receive os signals
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return s.httpServer.Shutdown(ctx)
}
