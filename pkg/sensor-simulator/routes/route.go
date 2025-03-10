package routes

import (
	"fmt"
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "haloooo")
	})
	// r.HandleFunc("GET /sensors")
	// r.HandleFunc("GET /sensors/{id}")

	return r
}
