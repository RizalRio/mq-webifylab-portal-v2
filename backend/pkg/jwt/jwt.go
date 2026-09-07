package jwt

import (
	"errors"
	"fmt"
	"time"

	"webifylab-backend/internal/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Claims mendefinisikan struktur payload JWT kita
type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken membuat access token baru
func GenerateToken(userID uuid.UUID, role string) (string, error) {
	// Parse durasi dari config (default 24h)
	expiry := config.AppConfig.JWTExpiry
	if expiry == 0 {
		expiry = 24 * time.Hour
	}

	claims := &Claims{
		UserID: userID.String(),
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "webifylab-api",
		},
	}

	// Sign token dengan HMAC-SHA256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	// Ambil secret dari config
	secret := []byte(config.AppConfig.JWTSecret)
	if len(secret) < 32 {
		return "", errors.New("JWT_SECRET must be at least 32 characters long")
	}

	return token.SignedString(secret)
}

// ValidateToken memverifikasi token dan mengembalikan claims
func ValidateToken(tokenString string) (*Claims, error) {
	secret := []byte(config.AppConfig.JWTSecret)

	// Parse token dengan validasi
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Pastikan signing method adalah HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	// Ekstrak claims jika token valid
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}