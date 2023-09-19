package user

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

func (u *userRepo) GenerateUserHash(password string) (hash string, err error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}	
	argonHash := argon2.IDKey([]byte(password), salt, u.time, u.memory, u.threads, u.keyLen)
	b64Hash := u.encrypt
}

func (u *userRepo) encrypt(text []byte) string {
	nonce := 
}
