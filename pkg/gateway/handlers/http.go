package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type GatewayHandler struct {
	// service handlers
}

func NewGatewayHandler() *GatewayHandler {
	return &GatewayHandler{}
}

func (h *GatewayHandler) RegisterRoutes(r chi.Router) {
	r.Route("/api", func(r chi.Router) {
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
}
