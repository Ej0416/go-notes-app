package notes

import (
	"log"
	"net/http"

	repo "github.com/Ej0416/go-note-app/internal/adapters/postgresql/sqlc"
	"github.com/Ej0416/go-note-app/internal/json"
	"github.com/Ej0416/go-note-app/internal/middleware"
	"github.com/Ej0416/go-note-app/internal/types"
)

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	authUser, ok := r.Context().Value(middleware.UserContextKey).(types.AuthUser)
	if !ok {
		json.Write(w, http.StatusUnauthorized, types.APIResponse{
			Success: false,
			Error:   "unauthorized",
		})
		return
	}

	var req repo.CreateNoteParams
	if err := json.Read(r, &req); err != nil {
		log.Printf("error parsing request params: %s ", err)
		json.Write(w, http.StatusBadRequest, types.APIResponse{
			Success: false,
			Error: "invalild request body",
		})
		return
	}

	err := h.service.CreateNote(r.Context(), repo.CreateNoteParams{
		UserID: authUser.ID,
		Title: req.Title,
		Body: req.Body,
	})

	if err != nil {
		log.Printf("failed to create note: %s ", err)
		json.Write(w, http.StatusBadRequest, types.APIResponse{
			Success: false,
			Error: "failed to create note",
		})
		return
	}

	json.Write(w, http.StatusOK, types.APIResponse{
		Success: true,
		Data: nil,
	})
}