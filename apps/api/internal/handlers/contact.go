package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"webifylab/api/internal/models"
	"webifylab/api/internal/services"
)

type ContactHandler struct {
	service  services.ContactService
	validate *validator.Validate
}

func NewContactHandler(service services.ContactService) *ContactHandler {
	return &ContactHandler{
		service:  service,
		validate: validator.New(),
	}
}

func (h *ContactHandler) SubmitContact(w http.ResponseWriter, r *http.Request) {
	var contact models.ContactSubmission

	// Parse JSON
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, `{"error": "Invalid JSON format"}`, http.StatusBadRequest)
		return
	}

	// Validate required fields
	if err := h.validate.Struct(contact); err != nil {
		http.Error(w, `{"error": "Validation failed: `+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	// Gather metadata
	ip := r.RemoteAddr
	userAgent := r.UserAgent()
	contact.IPAddress = &ip
	contact.UserAgent = &userAgent

	// Process submission
	if err := h.service.ProcessContactSubmission(r.Context(), &contact); err != nil {
		http.Error(w, `{"error": "Failed to process submission"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Contact submission received successfully",
	})
}
