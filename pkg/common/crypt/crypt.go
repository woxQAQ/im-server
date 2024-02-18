package crypt

import (
	"fmt"
	"golang.org/x/crypto/scrypt"
)

func PasswordEncrypt(pass string, salt string) string {
	dk, _ := scrypt.Key([]byte(pass), []byte(salt), 32768, 0, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}
