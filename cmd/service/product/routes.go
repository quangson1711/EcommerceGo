package product

import (
	"Ecommerce-Go/types"
	"Ecommerce-Go/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProducts).Methods("GET")
	router.HandleFunc("/product", h.handleGetProductByID).Methods("GET")
	router.HandleFunc("/product", h.handleCreateProduct).Methods("POST")
}

func (h *Handler) handleGetProducts(writer http.ResponseWriter, request *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WirteError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WirteJson(writer, http.StatusOK, ps)
}

func (h *Handler) handleCreateProduct(writer http.ResponseWriter, request *http.Request) {

}

func (h *Handler) handleGetProductByID(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	idStr := query.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WirteError(writer, http.StatusBadRequest, err)
		return
	}

	p, err := h.store.GetProductByID(id)
	if err != nil {
		utils.WirteError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WirteJson(writer, http.StatusOK, p)
}
