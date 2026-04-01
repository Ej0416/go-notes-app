package auth

import (
	"context"
	"log"

	repo "github.com/Ej0416/go-note-app/internal/adapters/postgresql/sqlc"
	"github.com/Ej0416/go-note-app/internal/utils"
)

func NewService(repo repo.Querier) Service {
	return &svc{ repo: repo }
}

func(s *svc) AddUsers(ctx context.Context, args repo.AddUsersParams) error {
	// hashed password
	hashed, err := utils.HashPassword(args.PasswordHash)

	if err != nil {
		log.Printf("error hashing password: %s", err)
		return err
	}

	return s.repo.AddUsers(ctx, repo.AddUsersParams{
		Email: args.Email,
		FirstName: args.FirstName,
		LastName: args.LastName,
		PasswordHash: hashed,
	})
}

func(s *svc) GetUserAuth(ctx context.Context, email, password string) (repo.GetUserAuthRow, error) {
	return s.repo.GetUserAuth(ctx, email)
}