package notes

import (
	"context"
	"errors"
	"strings"

	repo "github.com/Ej0416/go-note-app/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) CreateNote(ctx context.Context, arg repo.CreateNoteParams) error {
	_, err := s.repo.GetUserByID(ctx, arg.UserID)
	if err != nil {
		return err
	}

	if len(strings.TrimSpace(arg.Title)) == 0 {
		return errors.New("title is required")
	}

	if len(arg.Title) > 50 {
		return errors.New("title too long")
	}

	if len(strings.TrimSpace(arg.Body)) == 0 {
		return errors.New("body is required")
	}

	if len(arg.Body) > 200 {
		return errors.New("body too long")
	}

	return s.repo.CreateNote(ctx, arg)
}

func (s *svc) ListAllNotes(ctx context.Context, arg repo.ListAllNotesParams, id pgtype.UUID) ([]repo.Note, error) {
	_, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return []repo.Note{}, err
	}

	return s.repo.ListAllNotes(ctx, arg)
}

func (s *svc) ListUserNotes(ctx context.Context, arg repo.ListUserNotesParams) ([]repo.Note, error) {
	_, err := s.repo.GetUserByID(ctx, arg.UserID)
	if err != nil {
		return []repo.Note{}, err
	}


	return s.repo.ListUserNotes(ctx, arg)
}

func (s *svc) EditNotes(ctx context.Context, arg repo.EditNotesParams, userID pgtype.UUID) (repo.Note, error) {
	_, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return repo.Note{}, err
	}

	return s.repo.EditNotes(ctx, arg)
}

func (s *svc) GetNotesByID(ctx context.Context, id pgtype.UUID) (repo.Note, error) {
	return s.repo.GetNotesByID(ctx, id)
}

func (s *svc) DeleteNotes(ctx context.Context) (repo.Note, error) {
	return s.repo.DeleteNotes(ctx)
}
