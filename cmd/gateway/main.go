package main

import (
	"log"
	"net/http"

	"github.com/skni-kod/iot-monitor-backend/pkg/gateway/handlers"
)

const (
	httpAddr = ":8080"
)

func main() {
	mux := http.NewServeMux()
	handler := handlers.NewHandler()
	handler.registerRoutes(mux)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}
}
