package hasher

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	// DefaultCost adalah cost factor untuk bcrypt (10 adalah balance antara security & performance)
	DefaultCost = 10
)

// HashPassword meng-hash password plain text menggunakan bcrypt
func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// CheckPassword memverifikasi password plain text terhadap hash bcrypt
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ValidatePasswordStrength memvalidasi kekuatan password
// Returns error jika password tidak memenuhi kriteria
func ValidatePasswordStrength(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	if len(password) > 72 {
		return errors.New("password must not exceed 72 characters (bcrypt limit)")
	}
	return nil
}