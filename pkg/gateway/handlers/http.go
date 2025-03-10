package handlers

import "net/http"

type Handler struct {
	// gateway
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/sensors", h.HandleGetSensors)
}

func (h *Handler) HandleGetSensors(w http.ResponseWriter, r *http.Request) {
}
