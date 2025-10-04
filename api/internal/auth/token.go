package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func NewSessionToken() (token string, hash []byte, err error) {
	b := make([]byte, 16)
	if _, err = rand.Read(b); err != nil {
		return "", nil, err
	}
	token = hex.EncodeToString(b)
	h := sha256.Sum256([]byte(token))
	return token, h[:], nil
}
