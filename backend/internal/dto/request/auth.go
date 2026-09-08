package request

// LoginRequest adalah DTO untuk request login
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// ForgotPasswordRequest adalah DTO untuk forgot password
type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// ResetPasswordRequest adalah DTO untuk reset password
type ResetPasswordRequest struct {
	Token               string `json:"token" validate:"required"`
	Password            string `json:"password" validate:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
}

// RefreshTokenRequest adalah DTO untuk refresh token
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}