package constants

import (
	"time"
)

const EXPIRATION_TIME = 5 * time.Minute

var const_JWT_KEY = []byte("sample_key")

func Get_const_JWT_KEY() []byte {
	return const_JWT_KEY
}
