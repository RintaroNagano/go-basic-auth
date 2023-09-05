package myhash

import (
	"crypto/sha256"
	"encoding/hex"
)

func PasswordToHash(password string) string {
	checkSum := sha256.Sum256([]byte(password))
	hashpass := hex.EncodeToString(checkSum[:])

	return hashpass
}
