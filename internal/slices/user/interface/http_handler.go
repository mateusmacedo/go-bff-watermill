package iface

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/mateusmacedo/bff-watermill/internal/slices/user/application"
)

type UserHTTPHandler struct {
	commandHandler *application.CreateUserCommandHandler
	queryHandler   *application.GetUserQueryHandler
}

func NewUserHTTPHandler(
	commandHandler *application.CreateUserCommandHandler,
	queryHandler *application.GetUserQueryHandler,
) *UserHTTPHandler {
	return &UserHTTPHandler{
		commandHandler: commandHandler,
		queryHandler:   queryHandler,
	}
}

func (h *UserHTTPHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var cmd application.CreateUserCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	id, err := h.commandHandler.Handle(cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func (h *UserHTTPHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	query := application.GetUserQuery{UserID: userID}
	user, err := h.queryHandler.Handle(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHTTPHandler) RegisterRoutes(router chi.Router) {
	router.Post("/users", h.HandleCreateUser)
	router.Get("/users/{userID}", h.HandleGetUser)
}
