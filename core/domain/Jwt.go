package domain

type JwtResponse struct {
	Token      string
	Expiration int64
}

type JwtAuth interface {
	CreateToken(userId string) (JwtResponse, error)
	Verify(token string) bool
}
