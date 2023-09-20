package user

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

const (
	cryptFormat = "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
)

func (u *userRepo) GenerateUserHash(password string) (hash string, err error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	argonHash := argon2.IDKey([]byte(password), salt, u.time, u.memory, u.threads, u.keyLen)
	b64Hash := u.encrypt(argonHash)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)

	encodedHash := fmt.Sprintf(cryptFormat, argon2.Version, u.memory, u.time, u.threads, b64Salt, b64Hash)

	return encodedHash, nil
}

func (u *userRepo) encrypt(text []byte) string {
	nonce := make([]byte, u.gcm.NonceSize())

	cipherText := u.gcm.Seal(nonce, nonce, text, nil)

	return base64.StdEncoding.EncodeToString(cipherText)
}
