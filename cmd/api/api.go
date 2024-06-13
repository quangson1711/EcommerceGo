package api

import (
	"Ecommerce-Go/cmd/service/cart"
	"Ecommerce-Go/cmd/service/order"
	"Ecommerce-Go/cmd/service/product"
	"Ecommerce-Go/cmd/service/user"
	"Ecommerce-Go/middleware"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	//Log all request response
	// Define your routes
	router.HandleFunc("/", HomeHandler).Methods("GET")

	// Add the logging middleware
	router.Use(middleware.LoggingMiddleware)

	subrouter := router.PathPrefix("/api").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)

	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)

	productHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(s.db)
	cartHandler := cart.NewHandle(orderStore, productStore, userStore, s.db)
	cartHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

}
