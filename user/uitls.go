package user

import (
	"crypto/sha256"
	"fmt"
)

func crypt(s string) string{
	var pwd = s
	crypt := sha256.New()
	crypt.Write([]byte(pwd))
	pwd = fmt.Sprintf("%x", crypt.Sum(nil))
	return pwd
}
