package service

import (
	"aic-be-playground/core/domain"

	"golang.org/x/crypto/bcrypt"
)

type bcryptHasher struct {
	salt int
}

func NewbcryptHasherHasher(salt int) domain.Hasher {
	return &bcryptHasher{salt: salt}
}

func (b *bcryptHasher) Hash(s string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(s), b.salt)
	return string(res), err
}

func (b *bcryptHasher) Verify(hashed, original string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(original))
}