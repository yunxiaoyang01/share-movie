package utils

import "testing"

func TestBcryptHash(t *testing.T) {
	password := "123456"
	hash := BcryptHash(password)
	t.Log(hash)
}
