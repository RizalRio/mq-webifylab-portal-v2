package main

import (
	"encoding/json"
	"fmt"

	"webifylab-backend/pkg/validator"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func main() {
	validator.Init()

	// Test valid request
	validReq := LoginRequest{
		Email:    "admin@webifylab.my.id",
		Password: "admin123",
	}
	if errors := validator.Validate(validReq); errors == nil {
		fmt.Println("✅ Valid request passed")
	} else {
		fmt.Println("❌ Valid request failed:", errors)
	}

	// Test invalid email
	invalidEmail := LoginRequest{
		Email:    "not-an-email",
		Password: "admin123",
	}
	if errors := validator.Validate(invalidEmail); errors != nil {
		fmt.Println("✅ Invalid email berhasil dideteksi:")
		prettyPrint(errors)
	}

	// Test missing fields
	missingFields := LoginRequest{}
	if errors := validator.Validate(missingFields); errors != nil {
		fmt.Println("✅ Missing fields berhasil dideteksi:")
		prettyPrint(errors)
	}

	// Test password too short
	shortPass := LoginRequest{
		Email:    "admin@webifylab.my.id",
		Password: "123",
	}
	if errors := validator.Validate(shortPass); errors != nil {
		fmt.Println("✅ Short password berhasil dideteksi:")
		prettyPrint(errors)
	}
}

func prettyPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
}