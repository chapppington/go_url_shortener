package api

import (
	"net/http"
	"strings"

	"urlshortener/internal/logic"
)

func NewRouter(service *logic.Service) http.Handler {
	handler := NewHandler(service)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/create", handler.CreateShortURL)
	mux.HandleFunc("/api/v1/", func(w http.ResponseWriter, r *http.Request) {
		shortCode := strings.TrimPrefix(r.URL.Path, "/api/v1/")
		if shortCode == "" {
			http.Error(w, "Short code is required", http.StatusBadRequest)
			return
		}
		handler.GetLongURL(w, r, shortCode)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Error(w, "Short code is required", http.StatusBadRequest)
			return
		}
		shortCode := strings.TrimPrefix(r.URL.Path, "/")
		handler.Redirect(w, r, shortCode)
	})

	return mux
}
