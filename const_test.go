package bovasdk

import (
	"math/rand"
	"time"
)

// Mock API secret
const apiUrl = "https://sandbox.bovatech.cc"
const apiSecret = "1cec9fe9f2e0e49ac80de6774eae074429a16816"
const userUUID = "a53fb67d-d807-4055-b7b3-56aafd88ff16"

func getRandMerchantID() string {
	length := 5
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(charset))]
	}
	return string(b)
}
