package middleware

import (
	"context"
	"net/http"

	"go-demo/api/internal/core/session"
	"go-demo/api/internal/env"
)

type userKeyType string

const userKey userKeyType = "userClaims"

func WithAuth(cfg env.Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(cfg.CookieName)
		if err != nil || c.Value == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		claims, err := session.Parse(c.Value, cfg.JWTSecret)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ClaimsFromCtx(r *http.Request) *session.Claims {
	if v := r.Context().Value(userKey); v != nil {
		if c, ok := v.(*session.Claims); ok {
			return c
		}
	}
	return nil
}
