package user

import (
	"log"
	"net/http"
	"strconv"

	repo "github.com/Ej0416/go-note-app/internal/adapters/postgresql/sqlc"
	"github.com/Ej0416/go-note-app/internal/json"
	"github.com/Ej0416/go-note-app/internal/middleware"
	"github.com/Ej0416/go-note-app/internal/types"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	// get query params limit and offset
	q := r.URL.Query()

	limitStr := q.Get("limit")
	offsetStr := q.Get("offset")

	// set defaults
	limit := 10
	offset := 0

	if limitStr != "" {
		if v, err := strconv.Atoi(limitStr); err == nil {
			limit = v
		}
	}

	if offsetStr != "" {
		if v, err := strconv.Atoi(offsetStr); err == nil {
			offset = v
		}
	}

	params := repo.ListUsersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	users, err := h.service.ListUsers(r.Context(), params)
	if err != nil {
		json.Write(w, http.StatusInternalServerError, types.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	json.Write(w, http.StatusOK, types.APIResponse{
		Success: true,
		Data:    users,
	})
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userIDString := chi.URLParam(r, "id")
	var userID pgtype.UUID
	if err := userID.Scan(userIDString); err != nil {
		log.Printf("invalid uuid: %v", err)
		json.Write(w, http.StatusBadRequest, types.APIResponse{
			Success: false,
			Error:   "invalid uuid",
		})
		return
	}

	user, err := h.service.GetUserByID(r.Context(), userID)
	if err != nil {
		log.Printf("error getting user: %s", err)
		json.Write(w, http.StatusBadRequest, types.APIResponse{
			Success: false,
			Error:   "user not found",
		})
	}

	json.Write(w, http.StatusOK, types.APIResponse{
		Success: true,
		Data:    user,
	})
}

func (h *handler) UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	authUser, ok := r.Context().Value(middleware.UserContextKey).(types.AuthUser)
	if !ok {
		json.Write(w, http.StatusUnauthorized, types.APIResponse{
			Success: false,
			Error:   "unauthorized",
		})
		return
	}

	var req UpdateUserRequest
	if err := json.Read(r, &req); err != nil {
		log.Printf("error parsing request params: %s ", err)
		json.Write(w, http.StatusBadRequest, types.APIResponse{
			Success: false,
			Error:   "invalild request body",
		})
		return
	}

	userUpdated, err := h.service.UpdateUserInfo(r.Context(), repo.UpdateUserInfoParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		ID:        authUser.ID,
	})

	if err != nil {
		log.Printf("failed to update use: %s ", err)
		json.Write(w, http.StatusBadRequest, types.APIResponse{
			Success: false,
			Error:   "failed to update use",
		})
		return
	}

	json.Write(w, http.StatusOK, types.APIResponse{
		Success: true,
		Data:    userUpdated,
	})
}

func (h *handler) ChangeUserEmail(w http.ResponseWriter, r *http.Request) {
	authUser, ok := r.Context().Value(middleware.UserContextKey).(types.AuthUser)
	if !ok {
		json.Write(w, http.StatusUnauthorized, types.APIResponse{
			Success: false,
			Error:   "unauthorized",
		})
		return
	}

	var req ChangeUserEmailRequest
	if err := json.Read(r, &req); err != nil {
		log.Printf("error parsing request params: %s ", err)
		json.Write(w, http.StatusBadRequest, types.APIResponse{
			Success: false,
			Error:   "invalild request body",
		})
		return
	}

	userChangedEmail, err := h.service.ChangeUserEmail(r.Context(), repo.ChangeUserEmailParams{
		Email: req.Email,
		ID:    authUser.ID,
	})

	if err != nil {
		log.Printf("failed to update user email: %s ", err)
		json.Write(w, http.StatusBadRequest, types.APIResponse{
			Success: false,
			Error:   "failed to update user email",
		})
		return
	}

	json.Write(w, http.StatusOK, types.APIResponse{
		Success: true,
		Data:    userChangedEmail,
	})
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	authUser, ok := r.Context().Value(middleware.UserContextKey).(types.AuthUser)
	if !ok {
		json.Write(w, http.StatusUnauthorized, types.APIResponse{
			Success: false,
			Error:   "unauthorized",
		})
		return
	}

	userDeleted, err := h.service.DeleteUser(r.Context(), authUser.ID)
	if err != nil {
		log.Printf("failed to delete user email: %s ", err)
		json.Write(w, http.StatusBadRequest, types.APIResponse{
			Success: false,
			Error:   "failed to delete user email",
		})
		return
	}

	json.Write(w, http.StatusOK, types.APIResponse{
		Success: true,
		Data:    userDeleted,
	})
}
