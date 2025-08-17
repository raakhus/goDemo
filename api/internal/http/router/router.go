package router

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"go-demo/api/internal/env"
	"go-demo/api/internal/http/handlers"
	"go-demo/api/internal/http/middleware"
	"go-demo/api/internal/storage/memory"
)

func New(cfg env.Config, repo *memory.UserRepo) http.Handler {
	r := chi.NewRouter()

	auth := handlers.NewAuthHandler(cfg, repo)

	r.Get("/api/healthz", handlers.Health)
	r.Post("/api/auth/login", auth.Login)

	r.Group(func(protected chi.Router) {
		protected.Use(func(next http.Handler) http.Handler { return middleware.WithAuth(cfg, next) })
		protected.Get("/api/auth/me", auth.Me)
		protected.Post("/api/auth/logout", auth.Logout)
		protected.Get("/api/secret", auth.Secret)
	})

	return r
}
