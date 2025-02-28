package utils

import (
	"crypto/md5"
	"fmt"
)

func GetMd5Str(str string) string {
	strToByte := []byte(str)
	return fmt.Sprintf("%x", md5.Sum(strToByte))
}
