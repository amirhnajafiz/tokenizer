package internal

import "github.com/amirhnajafiz/encrypto"

const (
	private = "xx#40op3^i)jn2jf"
)

// Code converts a cipher into coded value
func Code(cipher string) (string, error) {
	return encrypto.Encrypt(private, cipher)
}

// DeCode converts a coded value into cipher
func DeCode(token string) (string, error) {
	return encrypto.Decrypt(private, token)
}
