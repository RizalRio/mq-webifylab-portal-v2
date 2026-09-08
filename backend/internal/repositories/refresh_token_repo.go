package repositories

import (
	"errors"
	"time"

	"webifylab-backend/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

// NewRefreshTokenRepository membuat instance baru
func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

// Create menyimpan refresh token baru ke database
func (r *RefreshTokenRepository) Create(userID uuid.UUID, token string, expiresAt time.Time) error {
	refreshToken := models.RefreshToken{
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiresAt,
	}
	
	return r.db.Create(&refreshToken).Error
}

// GetByToken mencari refresh token berdasarkan token string
func (r *RefreshTokenRepository) GetByToken(token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	err := r.db.Where("token = ?", token).First(&refreshToken).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRefreshTokenNotFound
		}
		return nil, err
	}
	
	return &refreshToken, nil
}

// DeleteByUserID menghapus semua refresh token milik user (untuk logout semua device)
func (r *RefreshTokenRepository) DeleteByUserID(userID uuid.UUID) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.RefreshToken{}).Error
}

// DeleteByToken menghapus refresh token spesifik
func (r *RefreshTokenRepository) DeleteByToken(token string) error {
	return r.db.Where("token = ?", token).Delete(&models.RefreshToken{}).Error
}

// CleanupExpired menghapus refresh token yang sudah expired (dipanggil periodik)
func (r *RefreshTokenRepository) CleanupExpired() error {
	return r.db.Where("expires_at < ?", time.Now()).Delete(&models.RefreshToken{}).Error
}

var (
	ErrRefreshTokenNotFound = errors.New("refresh token not found")
)