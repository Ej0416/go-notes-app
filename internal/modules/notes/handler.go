package notes

import (
	"log"
	"net/http"

	repo "github.com/Ej0416/go-note-app/internal/adapters/postgresql/sqlc"
	"github.com/Ej0416/go-note-app/internal/json"
	"github.com/Ej0416/go-note-app/internal/middleware"
	"github.com/Ej0416/go-note-app/internal/types"
	"github.com/Ej0416/go-note-app/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
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
			Error:   "invalild request body",
		})
		return
	}

	err := h.service.CreateNote(r.Context(), repo.CreateNoteParams{
		UserID: authUser.ID,
		Title:  req.Title,
		Body:   req.Body,
	})

	if err != nil {
		log.Printf("failed to create note: %s ", err)
		json.Write(w, http.StatusBadRequest, types.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.Write(w, http.StatusOK, types.APIResponse{
		Success: true,
		Data:    nil,
	})
}


func(h *handler) ListAllNotes(w http.ResponseWriter, r *http.Request) {
	authUser, ok := r.Context().Value(middleware.UserContextKey).(types.AuthUser)
	if !ok {
		json.Write(w, http.StatusUnauthorized, types.APIResponse{
			Success: false,
			Error:   "unauthorized",
		})
		return
	}

	q := r.URL.Query()
	limit, offset := utils.LimitAndOffsetConverter(q.Get("limit"), q.Get("offset"))

	params := repo.ListAllNotesParams{
		Limit: limit, 
		Offset: offset,
	}

	notes, err := h.service.ListAllNotes(r.Context(), params, authUser.ID)
		if err != nil {
		json.Write(w, http.StatusInternalServerError, types.APIResponse{
			Success: false,
			Error:   "failed to fetch all notes list",
		})
	}

	json.Write(w, http.StatusOK, notes)
} 

func(h *handler) ListUserNotes(w http.ResponseWriter, r *http.Request) {
	authUser, ok := r.Context().Value(middleware.UserContextKey).(types.AuthUser)
	if !ok {
		json.Write(w, http.StatusUnauthorized, types.APIResponse{
			Success: false,
			Error:   "unauthorized",
		})
		return
	}

	q := r.URL.Query()
	limit, offset := utils.LimitAndOffsetConverter(q.Get("limit"), q.Get("offset"))

	params := repo.ListUserNotesParams{
		UserID: authUser.ID,
		Limit: limit,
		Offset: offset,
	}

	notes, err := h.service.ListUserNotes(r.Context(), params)
	if err != nil {
		json.Write(w, http.StatusInternalServerError, types.APIResponse{
			Success: false,
			Error:   "failed to fetch all user`s notes list",
		})
	}

	json.Write(w, http.StatusOK, notes)
}

func(h *handler) EditNotes(w http.ResponseWriter, r *http.Request){
	authUser, ok := r.Context().Value(middleware.UserContextKey).(types.AuthUser)
	if !ok {
		json.Write(w, http.StatusUnauthorized, types.APIResponse{
			Success: false,
			Error:   "unauthorized",
		})
		return
	}

	var req EditNotesParams
	if err := json.Read(r, &req); err != nil {
		log.Printf("error parsing request params: %s ", err)
		json.Write(w, http.StatusBadRequest, types.APIResponse{
			Success: false,
			Error:   "invalild request body",
		})
		return
	}

	note, err := h.service.EditNotes(r.Context(), repo.EditNotesParams{
		Title: req.Title,
		Body: req.Body,
		ID: req.ID,
		UserID: authUser.ID,
	})

	if err != nil {
		log.Printf("failed to edit note: %s ", err)
		json.Write(w, http.StatusBadRequest, types.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.Write(w, http.StatusOK, note)
}

func(h *handler) GetNotesByID(w http.ResponseWriter, r *http.Request){
	authUser, ok := r.Context().Value(middleware.UserContextKey).(types.AuthUser)
	if !ok {
		json.Write(w, http.StatusUnauthorized, types.APIResponse{
			Success: false,
			Error:   "unauthorized",
		})
		return
	}

	noteIDString := chi.URLParam(r, "id")
	var noteID pgtype.UUID
	if err := noteID.Scan(noteIDString); err != nil {
		log.Printf("invalid note uuid: %v", err)
		json.Write(w, http.StatusBadRequest, types.APIResponse{
			Success: false,
			Error:   "invalid note uuid",
		})
		return
	}

	note, err := h.service.GetNotesByID(r.Context(), noteID, authUser.ID)
	if err != nil {
		log.Printf("failed to get note: %s ", err)
		json.Write(w, http.StatusBadRequest, types.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.Write(w, http.StatusOK, note)
}

func(h *handler) DeleteNotes(w http.ResponseWriter, r *http.Request){
	authUser, ok := r.Context().Value(middleware.UserContextKey).(types.AuthUser)
	if !ok {
		json.Write(w, http.StatusUnauthorized, types.APIResponse{
			Success: false,
			Error:   "unauthorized",
		})
		return
	}

	noteIDString := chi.URLParam(r, "id")
	var noteID pgtype.UUID
	if err := noteID.Scan(noteIDString); err != nil {
		log.Printf("invalid note uuid: %v", err)
		json.Write(w, http.StatusBadRequest, types.APIResponse{
			Success: false,
			Error:   "invalid note uuid",
		})
		return
	}

	note, err := h.service.DeleteNotes(r.Context(), noteID, authUser.ID)
	if err != nil {
		log.Printf("failed to delete note: %s ", err)
		json.Write(w, http.StatusBadRequest, types.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.Write(w, http.StatusOK, note)
}