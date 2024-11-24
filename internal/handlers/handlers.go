package handlers

import (
	"net/http"

	"github.com/benleem/show-beam/internal/handlers/routes"
)

func Init() *http.ServeMux {
	mux := http.NewServeMux()
	indexHandler := routes.NewIndexHandler()
	mux.HandleFunc("GET /", indexHandler.Get)
	return mux
}
