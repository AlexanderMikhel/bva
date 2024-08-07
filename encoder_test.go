package bovasdk

import "testing"

// TestCalculateSignature tests the calculateSignature function
func TestCalculateSignature(t *testing.T) {
	body := []byte(`{"user_uuid":"364dbfc8-ae50-492f-bdd9-748edd84d5c9","amount":300,"callback_url":"https://example.com/callback"}`)
	expectedSignature := "1d41b723b630e0cd790e553b12293995f24a1dd8"
	signature := calculateSignature(apiSecret, body)
	if signature != expectedSignature {
		t.Errorf("calculateSignature() = %v, want %v", signature, expectedSignature)
	}
}

// TestVerifySignature tests the verifySignature function
func TestVerifySignature(t *testing.T) {
	body := []byte(`{"user_uuid":"364dbfc8-ae50-492f-bdd9-748edd84d5c9","amount":300,"callback_url":"https://example.com/callback"}`)
	expectedSignature := calculateSignature(apiSecret, body)
	if !verifySignature(apiSecret, body, expectedSignature) {
		t.Errorf("verifySignature() = false, want true")
	}

	if verifySignature(apiSecret, body, "invalid_signature") {
		t.Errorf("verifySignature() = true, want false")
	}
}
