package auth

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Data LoginResponseData `json:"data"`
}

type LoginResponseData struct {
	Token string `json:"token"`
}
