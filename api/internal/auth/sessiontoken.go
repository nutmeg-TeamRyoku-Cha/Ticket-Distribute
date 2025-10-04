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
	h := sha256.Sum256(b)
	return token, h[:], nil
}

func SessionHashFromToken(token string) ([]byte, error) {
	raw, err := hex.DecodeString(token)
	if err != nil {
		return nil, err
	}
	sum := sha256.Sum256(raw)
	return sum[:], nil
}
