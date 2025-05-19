package routes

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"tkh/middlewares"
	"tkh/order"
	"tkh/product"
)

func GetRoutes(ctx context.Context) *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc("/api/product", product.List).Methods(http.MethodGet)
	p := m.PathPrefix("/api/product").Subrouter()
	p.HandleFunc("/", product.List).Methods(http.MethodGet)
	p.HandleFunc("/{id}", product.Get).Methods(http.MethodGet)
	m.HandleFunc("/api/order", middlewares.Auth(http.HandlerFunc(order.Create)).ServeHTTP).Methods(http.MethodPost)
	o := m.PathPrefix("/api/order").Subrouter()
	o.HandleFunc("/", order.Create).Methods(http.MethodPost)
	o.Use(middlewares.Auth)
	return m
}
