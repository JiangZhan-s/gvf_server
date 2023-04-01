package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// EncodeMd5 md5加密
func EncodeMd5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
