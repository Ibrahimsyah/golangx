package domain

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type IAuthUseCase interface {
	LoginUser(user *AuthRequest) (*AuthResponse, error)
}
