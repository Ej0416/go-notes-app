package auth

import (
	"log"
	"net/http"

	repo "github.com/Ej0416/go-note-app/internal/adapters/postgresql/sqlc"
	"github.com/Ej0416/go-note-app/internal/json"
)

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// access the service params
	var req repo.AddUsersParams

	// read json fom http and parses it to the stuct data (req -> service params)
	if err := json.Read(r, &req); err != nil {
		log.Printf("error in adding user: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// pass the data to the service layer and do error handling
	err := h.service.AddUsers(r.Context(),req)
	if err != nil {
		log.Printf("error creating user: %s", err)
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	json.Write(w, 200, nil)
}



func(h *handler) LoginUser(w http.ResponseWriter, r *http.Request){

}