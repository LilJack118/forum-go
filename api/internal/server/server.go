package server

import (
	"context"
	"forum/api/config"
	"forum/api/internal/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	accounthttp "forum/api/internal/account/delivery/http"
	accountuc "forum/api/internal/account/usecase"
	authhttp "forum/api/internal/auth/delivery/http"
	authrepo "forum/api/internal/auth/repository"
	authuc "forum/api/internal/auth/usecase"
	postshttp "forum/api/internal/posts/delivery/http"
	postsrepo "forum/api/internal/posts/repository"
	postsuc "forum/api/internal/posts/usecase"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

// Server struct
type Server struct {
	httpServer *http.Server
	db         *mongo.Database
}

// NewServer New Server constructor
func NewServer(db *mongo.Database) *Server {
	return &Server{db: db}
}

func (s *Server) Run(port string) error {

	router := mux.NewRouter()
	router.Use(middleware.DefaultMiddleware)

	// init repositiories
	auth_repo := authrepo.NewAuthRepository(s.db)
	posts_repo := postsrepo.NewPostsRepository(s.db)

	// init use cases
	auth_uc := authuc.NewAuthUseCase(auth_repo)
	account_uc := accountuc.NewAccountUseCase(auth_repo)
	posts_uc := postsuc.NewPostsUseCase(posts_repo)

	// register routes
	auth_router := router.PathPrefix("/auth").Subrouter()
	authhttp.RegisterAuthNotProtectedRoutes(auth_router, auth_uc)

	auth_protected_router := router.PathPrefix("/auth").Subrouter()
	auth_protected_router.Use(middleware.AuthJWTMiddleware)
	authhttp.RegisterAuthProtectedRoutes(auth_protected_router, auth_uc)

	api_router := router.PathPrefix("/api").Subrouter()
	api_router.Use(middleware.AuthJWTMiddleware)

	accounthttp.RegisterAccountRoutes(api_router, account_uc)
	postshttp.RegisterPostHandlers(api_router, posts_uc)

	originsStr, _ := config.Config("ALLOWED_ORIGINS", "string")
	origins := strings.Split(originsStr.(string), ",")
	// set cors
	c := cors.New(cors.Options{
		AllowedOrigins: origins,
		AllowedHeaders: []string{"*"},
		Debug:          true,
	})

	handler := c.Handler(router)

	s.httpServer = &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
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
