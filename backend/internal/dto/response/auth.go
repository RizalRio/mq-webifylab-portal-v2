package response

import "time"

// LoginResponse adalah DTO untuk response login
type LoginResponse struct {
	AccessToken  string     `json:"access_token"`
	RefreshToken string     `json:"refresh_token"`
	ExpiresIn    int64      `json:"expires_in"` // dalam detik
	User         UserResponse `json:"user"`
}

// UserResponse adalah DTO untuk user data (tanpa password)
type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// RefreshTokenResponse adalah DTO untuk response refresh token
type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}