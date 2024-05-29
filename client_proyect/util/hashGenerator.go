package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func HashPassword(password string) string {

	hasher := sha256.New()

	hasher.Write([]byte(password))

	return hex.EncodeToString(hasher.Sum(nil))

}

func CheckPasswordHash(password string, hash string) bool {

	fmt.Println("input contrasena:", HashPassword(password))
	fmt.Println("actual contrasena:", hash)

	return HashPassword(password) == hash

}
