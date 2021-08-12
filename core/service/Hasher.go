package service

import (
	"golangx/core/domain"

	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct {
	salt int
}

func NewBcryptHasher(salt int) domain.Hasher {
	return &BcryptHasher{salt: salt}
}

func (b *BcryptHasher) Hash(s string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(s), b.salt)
	return string(res), err
}

func (b *BcryptHasher) Verify(hashed, original string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(original))
	return err == nil
}
