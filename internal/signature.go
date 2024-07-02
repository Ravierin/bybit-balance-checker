package internal

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSignature(secret, apiKey, timestamp, recvWindow, queryString string) string {
	message := timestamp + apiKey + recvWindow + queryString
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
