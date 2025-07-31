package service

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string, salt []byte) []byte {
	return argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
}

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 32) // 16 bytes is minimum, 32 is better
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}
