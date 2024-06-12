package cart

import (
	"Ecommerce-Go/cmd/service/auth"
	"Ecommerce-Go/types"
	"Ecommerce-Go/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
)

type Handle struct {
	store        types.OrderStore
	productStore types.ProductStore
	userStore    types.UserStore
}

func NewHandle(store types.OrderStore, productStore types.ProductStore, userStore types.UserStore) *Handle {
	return &Handle{store, productStore, userStore}
}

func (h *Handle) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", auth.WithJWTAuth(h.handleCheckout, h.userStore)).Methods(http.MethodPost)
}

func (h *Handle) handleCheckout(w http.ResponseWriter, r *http.Request) {
	// TODO userID
	userID := auth.GetUserIDFromContext(r.Context())
	var payload types.CartCheckoutPayload

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WirteError(w, http.StatusBadRequest, err)
		return
	}

	// validate request
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WirteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
	}

	// get products
	productIDs, err := getCartItemIDs(payload.Items)

	if err != nil {
		utils.WirteError(w, http.StatusBadRequest, err)
		return
	}

	ps, err := h.productStore.GetProductByIDs(productIDs)

	if err != nil {
		utils.WirteError(w, http.StatusBadRequest, err)
		return
	}

	orderID, totalPrice, err := h.createOrder(ps, payload.Items, userID)

	response := types.NewCheckoutCartResponseBody(totalPrice, orderID)

	utils.WirteJson(w, http.StatusOK, response)

}
