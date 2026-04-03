package auth

import (
	"context"
	"log"
	"time"

	repo "github.com/Ej0416/go-note-app/internal/adapters/postgresql/sqlc"
	"github.com/Ej0416/go-note-app/internal/env"
	"github.com/Ej0416/go-note-app/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

var jwtSecret = []byte(env.GetString("JWT_SECRET", "te-mp-or-ar-ry-ke-y!"))

func (s *svc) AddUsers(ctx context.Context, args repo.AddUsersParams) error {
	// hashed password
	hashed, err := utils.HashPassword(args.PasswordHash)

	if err != nil {
		log.Printf("error hashing password: %s", err)
		return err
	}

	return s.repo.AddUsers(ctx, repo.AddUsersParams{
		Email:        args.Email,
		FirstName:    args.FirstName,
		LastName:     args.LastName,
		PasswordHash: hashed,
	})
}

func (s *svc) GetAuthToken(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetUserAuth(ctx, email)

	if err != nil {
		return "", err
	}

	// check password
	if err := utils.CheckPassword(password, user.PasswordHash); err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), //24 hour expiry
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
