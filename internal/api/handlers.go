package api

import (
	"encoding/json"
	"net/http"

	"urlshortener/internal/logic"
)

type Handler struct {
	service *logic.Service
}

func NewHandler(service *logic.Service) *Handler {
	return &Handler{service: service}
}

type CreateShortURLSchema struct {
	URL string `json:"url"`
}

type CreateShortURLResponseSchema struct {
	ShortCode string `json:"short_code"`
	LongURL   string `json:"long_url"`
}

func (h *Handler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateShortURLSchema
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortURL, err := h.service.CreateShortURL(req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreateShortURLResponseSchema{
		ShortCode: shortURL.ShortCode,
		LongURL:   shortURL.LongUrl,
	})
}

func (h *Handler) GetLongURL(w http.ResponseWriter, r *http.Request, shortCode string) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	longURL, err := h.service.GetLongURL(shortCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"long_url": longURL})
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request, shortCode string) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	longURL, err := h.service.GetLongURL(shortCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

