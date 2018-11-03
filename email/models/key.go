package models

import (
	"encoding/base64"

	"golang.org/x/crypto/scrypt"
)

var salt = "dksjhfjnsldmf;s0-asdfhw3efms"

//密码加密
func Getkey(password string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
	return base64.StdEncoding.EncodeToString(dk)
}
