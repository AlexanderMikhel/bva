package bovasdk

import (
	"crypto/sha1"
	"encoding/hex"
)

// calculateSignature вычисляет SHA1 подпись для тела запроса и api_secret.
func calculateSignature(apiSecret string, body []byte) string {
	hash := sha1.New()
	hash.Write([]byte(apiSecret + string(body)))
	return hex.EncodeToString(hash.Sum(nil))
}

// verifySignature проверяет подпись, вычисленную для тела ответа и api_secret.
func verifySignature(apiSecret string, body []byte, receivedSignature string) bool {
	expectedSignature := calculateSignature(apiSecret, body)
	return expectedSignature == receivedSignature
}
