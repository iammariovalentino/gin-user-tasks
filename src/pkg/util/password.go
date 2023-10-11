package util

import (
	"crypto/md5"
	"fmt"
	"io"
)

// HashPassword hash password
func HashPassword(pwd string) string {
	h := md5.New()
	io.WriteString(h, pwd)
	password := fmt.Sprintf("%x", h.Sum(nil))
	return password
}
