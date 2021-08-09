package domain

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiredAt   int64  `json:"expired_at"`
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type IAuthUseCase interface {
	LoginUser(user *AuthRequest) (*AuthResponse, error)
}
