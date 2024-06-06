package user

import (
	"Ecommerce-Go/cmd/service/auth"
	"Ecommerce-Go/types"
	"Ecommerce-Go/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload

	err := utils.ParseJson(r, payload)
	if err != nil {
		utils.WirteError(w, http.StatusBadRequest, err)
		return
	}

	_, err = h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WirteError(w, http.StatusBadRequest, fmt.Errorf("user with this email %s already exists", payload.Email))
		return
	}

	hashPasswrord, err := auth.HashPassword(payload.Password)

	if err != nil {
		utils.WirteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashPasswrord,
	})

	if err != nil {
		utils.WirteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WirteJson(w, http.StatusOK, nil)
}
