package domain

type Hasher interface {
	Hash(s string) (string, error)
	Verify(hashed, original string) bool
}
