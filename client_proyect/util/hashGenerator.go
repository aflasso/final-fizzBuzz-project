package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password string) string {

	hasher := sha256.New()

	hasher.Write([]byte(password))

	return hex.EncodeToString(hasher.Sum(nil))

}

func CheckPasswordHash(password string, hash string) bool {

	return HashPassword(password) == hash

}
