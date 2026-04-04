package user

import (
	"context"

	repo "github.com/Ej0416/go-note-app/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service interface {
	ListUsers(ctx context.Context, args repo.ListUsersParams) ([]repo.User, error)
	GetUserByID(ctx context.Context, id pgtype.UUID) (repo.GetUserByIDRow, error)
	UpdateUserInfo(ctx context.Context, arg repo.UpdateUserInfoParams) (repo.User, error)
	ChangeUserEmail(ctx context.Context, arg repo.ChangeUserEmailParams) (repo.User, error)
	DeleteUser(ctx context.Context, id pgtype.UUID) (repo.User, error)
}

type svc struct {
	repo repo.Querier
}

type handler struct {
	service Service
}

// update user request
type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// change user email request

type ChangeUserEmailRequest struct {
	Email string `json:"email"`
}