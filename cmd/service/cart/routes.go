package cart

import (
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
}

func NewHandle(store types.OrderStore, productStore types.ProductStore) *Handle {
	return &Handle{store, productStore}
}

func (h *Handle) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("cart/checkout", h.handleCheckout).Methods(http.MethodPost)
}

func (h *Handle) handleCheckout(w http.ResponseWriter, r *http.Request) {
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

	//TODO
	h.createOrder(ps, payload.Items)

}
