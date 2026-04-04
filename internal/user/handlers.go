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
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(jwt.MapClaims)
	if !ok {
		json.Write(w, http.StatusUnauthorized, types.APIResponse{
			Success: false,
			Error: "unauthorized",
		})
		return
	}

	_, ok = claims["user_id"].(string)
	if !ok {
		json.Write(w, http.StatusUnauthorized, types.APIResponse{
			Success: false,
			Error: "invalid user_id",
		})
		return
	}

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
		Limit: int32(limit),
		Offset: int32(offset),
	}

	users, err := h.service.ListUsers(r.Context(), params)
	if err != nil {
		json.Write(w, http.StatusInternalServerError, types.APIResponse{
			Success: false,
			Error: err.Error(),
		})
	}

	json.Write(w, http.StatusOK, types.APIResponse{
		Success: true,
		Data: users,
	})
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(jwt.MapClaims)

	if !ok {
		json.Write(w, http.StatusUnauthorized, types.APIResponse{
			Success: false,
			Error: "unauthorized",
		})
		return
	}

	_, ok = claims["user_id"].(string)
	if !ok {
		json.Write(w, http.StatusUnauthorized, types.APIResponse{
			Success: false,
			Error: "invalid user id",
		})
		return
	}

	userIDString := chi.URLParam(r, "id")

	var userID pgtype.UUID
	if err := userID.Scan(userIDString); err != nil {
		log.Printf("invalid uuid: %v", err)
		json.Write(w, http.StatusBadRequest, types.APIResponse{
			Success: false,
			Error: "invalid uuid",
		})
		return
	}

	user, err := h.service.GetUserByID(r.Context(), userID)
	if err != nil {
		log.Printf("error getting user: %s", err)
		json.Write(w, http.StatusBadRequest, types.APIResponse{
			Success: false,
			Error: "user not found",
		})
	}

	json.Write(w, http.StatusOK, types.APIResponse{
		Success: true,
		Data: user,
	})
}

func (h *handler) UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
}

func (h *handler) ChangeUserEmail(w http.ResponseWriter, r *http.Request) {
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
}
