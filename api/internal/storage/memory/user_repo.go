package memory

import (
	"go-demo/api/internal/core/models"
	"go-demo/api/internal/core/security"
)

type UserRepo struct {
	users map[string]models.User // keyed by email
}

func NewUserRepo() *UserRepo {
	// seed one demo user: test@example.com / Passw0rd!
	hash, _ := security.Hash("Passw0rd!")
	u := models.User{ID: "u_1", Email: "test@example.com", Password: hash}
	return &UserRepo{users: map[string]models.User{u.Email: u}}
}

func (r *UserRepo) GetByEmail(email string) (models.User, bool) {
	u, ok := r.users[email]
	return u, ok
}
