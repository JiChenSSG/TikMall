package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func String2MD5(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}
