package user

import (
	"context"
	"log"

	repo "github.com/Ej0416/go-note-app/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListUsers(ctx context.Context, args repo.ListUsersParams) ([]repo.User, error) {
	users, err := s.repo.ListUsers(ctx, args)
	if err != nil {
		log.Printf("error in getting users list: %s", err)
		return []repo.User{}, err
	}

	return users, nil
}

func (s *svc) GetUserByID(ctx context.Context, id pgtype.UUID) (repo.GetUserByIDRow, error) {
	user, err := s.repo.GetUserByID(ctx, id)

	if err != nil {
		log.Printf("no user found: %s", err)
		return repo.GetUserByIDRow{}, err
	}

	return user, nil
}

func (s *svc) UpdateUserInfo(ctx context.Context, arg repo.UpdateUserInfoParams) (repo.User, error) {
	userUpdated, err := s.repo.UpdateUserInfo(ctx, arg)
	if err != nil {
		log.Printf("update user info failed: %s", err)
		return repo.User{}, err
	}

	return userUpdated, nil
}

func (s *svc) ChangeUserEmail(ctx context.Context, arg repo.ChangeUserEmailParams) (repo.User, error) {
	userChangedEmail, err := s.repo.ChangeUserEmail(ctx, arg)
	if err != nil {
		log.Printf("update user email failed: %s", err)
		return repo.User{}, err
	}

	return userChangedEmail, nil
}

func (s *svc) DeleteUser(ctx context.Context, id pgtype.UUID) (repo.User, error) {
	return repo.User{}, nil
}
