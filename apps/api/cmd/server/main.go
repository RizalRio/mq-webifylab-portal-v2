package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	
	"webifylab/api/db"
	"webifylab/api/internal/config"
	"webifylab/api/internal/handlers"
	"webifylab/api/internal/repositories"
	"webifylab/api/internal/services"
)

func main() {
	// 1. Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. Setup Database
	pool, err := db.NewPostgresPool(cfg)
	if err != nil {
		log.Printf("Warning: Failed to connect to database. Starting without DB connection. Error: %v", err)
		// For local dev, we might continue without DB if we just want to test ping
	} else {
		log.Println("Successfully connected to database")
		defer pool.Close()
	}

	// 3. Initialize Repo, Service, Handlers
	var contactHandler *handlers.ContactHandler
	if pool != nil {
		contactRepo := repositories.NewContactRepository(pool)
		contactService := services.NewContactService(contactRepo, cfg)
		contactHandler = handlers.NewContactHandler(contactService)
	}

	// 4. Setup Router
	r := chi.NewRouter()
	
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	
	// Basic CORS setup for local development
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			
			next.ServeHTTP(w, r)
		})
	})

	r.Get("/api/health", handlers.HealthCheck)
	
	if contactHandler != nil {
		r.Post("/api/contact", contactHandler.SubmitContact)
	} else {
		r.Post("/api/contact", func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, `{"error": "Database not configured"}`, http.StatusInternalServerError)
		})
	}

	// 5. Start Server
	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Starting server on port %s in %s mode", port, cfg.Environment)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
