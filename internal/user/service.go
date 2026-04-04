package user

import (
	"context"

	repo "github.com/Ej0416/go-note-app/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *svc) NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) GetUserByID(ctx context.Context, id pgtype.UUID) (repo.GetUserByIDRow, error) {
	return repo.GetUserByIDRow{}, nil
}

func (s *svc) UpdateUserInfo(ctx context.Context, arg repo.UpdateUserInfoParams) (repo.User, error) {
	return repo.User{}, nil
}

func (s *svc) ChangeUserEmail(ctx context.Context, arg repo.ChangeUserEmailParams) (repo.User, error) {
	return repo.User{}, nil
}

func (s *svc) DeleteUser(ctx context.Context, id pgtype.UUID) (repo.User, error) {
	return repo.User{}, nil
}
