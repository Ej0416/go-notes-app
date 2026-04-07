package notes

import (
	"context"

	repo "github.com/Ej0416/go-note-app/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) CreateNote(ctx context.Context, arg repo.CreateNoteParams) error {
	return s.repo.CreateNote(ctx, arg)
}

func (s *svc) ListAllNotes(ctx context.Context, arg repo.ListAllNotesParams) ([]repo.Note, error) {
	return s.repo.ListAllNotes(ctx, arg)
}

func (s *svc) ListUserNotes(ctx context.Context, arg repo.ListUserNotesParams) ([]repo.Note, error) {
	return s.repo.ListUserNotes(ctx, arg)
}

func (s *svc) EditNotes(ctx context.Context, arg repo.EditNotesParams) (repo.Note, error) {
	return s.repo.EditNotes(ctx, arg)
}

func (s *svc) GetNotesByID(ctx context.Context, id pgtype.UUID) (repo.Note, error) {
	return s.repo.GetNotesByID(ctx, id)
}

func (s *svc) DeleteNotes(ctx context.Context) (repo.Note, error) {
	return s.repo.DeleteNotes(ctx)
}
