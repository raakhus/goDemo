package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"go-demo/api/internal/core/security"
	"go-demo/api/internal/core/session"
	"go-demo/api/internal/env"
	"go-demo/api/internal/http/middleware"
	"go-demo/api/internal/storage/memory"
)

type AuthHandler struct {
	cfg  env.Config
	repo *memory.UserRepo
}

func NewAuthHandler(cfg env.Config, repo *memory.UserRepo) *AuthHandler {
	return &AuthHandler{cfg: cfg, repo: repo}
}

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest); return
	}

	u, ok := h.repo.GetByEmail(req.Email)
	if !ok || !security.Check(u.Password, req.Password) {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	j, err := session.Sign(u.ID, u.Email, h.cfg.JWTSecret, 24*time.Hour)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError); return
	}

	cookie := &http.Cookie{
		Name:     h.cfg.CookieName,
		Value:    j,
		Path:     "/",
		HttpOnly: true,
		Secure:   h.cfg.CookieSecure,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(24 * time.Hour),
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusNoContent)
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	claims := middleware.ClaimsFromCtx(r)
	if claims == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized); return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"id": claims.UserID, "email": claims.Email,
	})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:     h.cfg.CookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   h.cfg.CookieSecure,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	}
	http.SetCookie(w, c)
	w.WriteHeader(http.StatusNoContent)
}

func (h *AuthHandler) Secret(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "You made it."})
}
