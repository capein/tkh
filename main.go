package main

import (
	"context"
	"net/http"
	"tkh/logger"
	"tkh/routes"
)

func main() {
	m := routes.GetRoutes(context.Background())
	s := &http.Server{
		Addr:    ":8080",
		Handler: m,
	}
	logger.Println("Starting server on :8080")
	err := s.ListenAndServeTLS("./certificate.pem", "./key.pem")
	if err != nil {
		logger.Println("error while starting the server", err)
		return
	}
}
