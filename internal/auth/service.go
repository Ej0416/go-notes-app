package auth

import (
	"context"

	repo "github.com/Ej0416/go-note-app/internal/adapters/postgresql/sqlc"
)

type Service interface {
	AddUsers(ctx context.Context, args repo.AddUsersParams) error
	GetUserAuth(ctx context.Context, email string) (repo.GetUserAuthRow, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{ repo: repo }
}

func(s *svc) AddUsers(ctx context.Context, args repo.AddUsersParams) error {
	return s.repo.AddUsers(ctx, args)
}

func(s *svc) GetUserAuth(ctx context.Context, email string) (repo.GetUserAuthRow, error) {
	return s.repo.GetUserAuth(ctx, email)
}