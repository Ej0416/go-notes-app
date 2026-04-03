package auth

import (
	"context"

	repo "github.com/Ej0416/go-note-app/internal/adapters/postgresql/sqlc"
)

// auth service types and interface
type Service interface {
	AddUsers(ctx context.Context, args repo.AddUsersParams) error
	GetAuthToken(ctx context.Context, email, password string) (string, error)
}

type svc struct {
	repo repo.Querier
}

// auth handler types
type handler struct {
	service Service
}

// extra types for auth
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type APIResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}
