package notes

import (
	"context"

	repo "github.com/Ej0416/go-note-app/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service interface {
	CreateNote(ctx context.Context, arg repo.CreateNoteParams) error
	ListAllNotes(ctx context.Context, arg repo.ListAllNotesParams, id pgtype.UUID) ([]repo.Note, error)
	ListUserNotes(ctx context.Context, arg repo.ListUserNotesParams) ([]repo.Note, error)
	EditNotes(ctx context.Context, arg repo.EditNotesParams, userID pgtype.UUID) (repo.Note, error)
	GetNotesByID(ctx context.Context, noteID pgtype.UUID, userID pgtype.UUID) (repo.Note, error)
	DeleteNotes(ctx context.Context, noteID pgtype.UUID, userID pgtype.UUID) (repo.Note, error)
}

type svc struct {
	repo repo.Querier
}

type handler struct {
	service Service
}
