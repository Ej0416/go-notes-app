package types

import "github.com/jackc/pgx/v5/pgtype"

type APIResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type AuthUser struct {
	ID pgtype.UUID `json:"user_id"`
}