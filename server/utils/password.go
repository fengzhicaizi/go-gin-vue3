package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

func DefaultPassword() string {
	var psw = "123456"
	m := md5.New()
	io.WriteString(m, psw)
	return fmt.Sprintf("%x", m.Sum(nil))
}
