package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

func ComputeHMACSHA1(data []byte, secretKey string) string {
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
