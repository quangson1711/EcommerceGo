package user

import (
	"Ecommerce-Go/cmd/service/auth"
	"Ecommerce-Go/types"
	"Ecommerce-Go/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
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
	var payload types.LoginUserPayload

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WirteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WirteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// check if the user exists
	u, err := h.store.GetUserByEmail(payload.Email)

	if err != nil {
		utils.WirteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	if !auth.CompareHashAndPassword(u.Password, payload.Password) {
		utils.WirteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	token, err := auth.CreateJWT(u.ID)

	if err != nil {
		utils.WirteError(w, http.StatusInternalServerError, err)
		return
	}

	response := types.NewLoginResponseBody(token)

	utils.WirteJson(w, http.StatusOK, response)

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WirteError(w, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WirteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// check if the user exists
	_, err := h.store.GetUserByEmail(payload.Email)
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

	utils.WirteJson(w, http.StatusCreated, nil)
}
