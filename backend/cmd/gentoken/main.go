package main

import (
	"fmt"
	"log"

	"webifylab-backend/internal/config"
	"webifylab-backend/pkg/jwt"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load("../../.env")
	config.LoadConfig()

	// Generate token untuk user dummy
	userID := uuid.New()
	role := "super_admin" // Coba ganti jadi "editor" untuk test role check

	token, err := jwt.GenerateToken(userID, role)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("TOKEN:")
	fmt.Println(token)
}