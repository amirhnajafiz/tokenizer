package internal

import (
	"os"

	"github.com/amirhnajafiz/encrypto"
)

const (
	private = "xx#40op3^i)jn2jf"
)

// get encryption private key
func getKey() string {
	key := os.Getenv("TK_PRIVATE")
	if len(key) == 0 {
		key = private
	}

	return key
}

// Code converts a cipher into coded value
func Code(cipher string) (string, error) {
	return encrypto.Encrypt(getKey(), cipher)
}

// DeCode converts a coded value into cipher
func DeCode(token string) (string, error) {
	return encrypto.Decrypt(getKey(), token)
}
