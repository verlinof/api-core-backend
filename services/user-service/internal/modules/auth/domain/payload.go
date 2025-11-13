package domain

type VerifyTokenPayload struct {
	Token string `json:"token" validate:"required"`
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshTokenPayload struct {
	Token        string `json:"token" validate:"required"`
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type RegisterPayload struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type GenerateTokenPayload struct {
	UserID string `json:"userId" validate:"required"`
}
