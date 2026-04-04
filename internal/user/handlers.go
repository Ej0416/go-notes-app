package user

import (
	"net/http"
)

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) ListUsers(w http.ResponseWriter, r *http.Request) {
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
}

func (h *handler) UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
}

func (h *handler) ChangeUserEmail(w http.ResponseWriter, r *http.Request) {
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
}
