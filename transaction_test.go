package bovasdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock data for MassTransaction
var (
	mockMassTransactionRequest = MassTransactionRequest{
		UserUUID:      "mock_user_uuid",
		Amount:        100,
		CallbackURL:   "https://example.com/callback",
		MerchantID:    "mock_merchant_id",
		Currency:      "USD",
		PaymentMethod: "card",
	}
	mockMassTransactionResponse = MassTransactionResponse{
		ResultCode: "ok",
		Payload: struct {
			ID                string `json:"id"`
			Amount            string `json:"amount"`
			CommissionType    string `json:"commission_type"`
			ServiceCommission string `json:"service_commission"`
			Rate              string `json:"rate"`
			FiatAmount        string `json:"fiat_amount"`
			OldFiatAmount     string `json:"old_fiat_amount"`
			State             string `json:"state"`
			Currency          string `json:"currency"`
			RLine             string `json:"r_line"`
			CreatedAt         string `json:"created_at"`
			UpdatedAt         string `json:"updated_at"`
			TotalAmount       string `json:"total_amount"`
			RecipientCard     string `json:"recipient_card"`
		}{
			ID:                "mock_id",
			Amount:            "100.0",
			CommissionType:    "fixed",
			ServiceCommission: "1.0",
			Rate:              "1.0",
			FiatAmount:        "100.0",
			OldFiatAmount:     "100.0",
			State:             "created",
			Currency:          "USD",
			RLine:             "",
			CreatedAt:         "2023-01-01T00:00:00Z",
			UpdatedAt:         "2023-01-01T00:00:00Z",
			TotalAmount:       "99.0",
			RecipientCard:     "411111******1111",
		},
	}
)

// TestCreateMassTransaction tests the CreateMassTransaction method
func TestCreateMassTransaction(t *testing.T) {
	server := createMockServerP2P(t, http.StatusOK, mockMassTransactionResponse)
	defer server.Close()

	mt := massTransactionNew(server.URL, apiSecret, server.Client())
	resp, err := mt.CreateMassTransaction(mockMassTransactionRequest)
	if err != nil {
		t.Fatalf("CreateMassTransaction() error = %v", err)
	}

	if resp.ResultCode != mockMassTransactionResponse.ResultCode {
		t.Errorf("CreateMassTransaction() ResultCode = %v, want %v", resp.ResultCode, mockMassTransactionResponse.ResultCode)
	}
}

// TestCreateMassTransaction_Error tests the CreateMassTransaction method for error response
func TestCreateMassTransaction_Error(t *testing.T) {
	server := createMockServerP2P(t, http.StatusBadRequest, nil)
	defer server.Close()

	mt := massTransactionNew(server.URL, apiSecret, server.Client())
	_, err := mt.CreateMassTransaction(mockMassTransactionRequest)
	if err == nil {
		t.Fatalf("CreateMassTransaction() expected error, got nil")
	}
}

// TestGetMassTransaction tests the GetMassTransaction method
func TestGetMassTransaction(t *testing.T) {
	server := createMockServerP2P(t, http.StatusOK, mockMassTransactionResponse)
	defer server.Close()

	mt := massTransactionNew(server.URL, apiSecret, server.Client())
	resp, err := mt.GetMassTransaction("mock_transaction_id")
	if err != nil {
		t.Fatalf("GetMassTransaction() error = %v", err)
	}

	if resp.ResultCode != mockMassTransactionResponse.ResultCode {
		t.Errorf("GetMassTransaction() ResultCode = %v, want %v", resp.ResultCode, mockMassTransactionResponse.ResultCode)
	}
}

// TestGetMassTransaction_Error tests the GetMassTransaction method for error response
func TestGetMassTransaction_Error(t *testing.T) {
	server := createMockServerP2P(t, http.StatusBadRequest, nil)
	defer server.Close()

	mt := massTransactionNew(server.URL, apiSecret, server.Client())
	_, err := mt.GetMassTransaction("mock_transaction_id")
	if err == nil {
		t.Fatalf("GetMassTransaction() expected error, got nil")
	}
}

// Функции и структуры, которые уже есть в файле sdk_test.go:

// Helper function to create a mock server
func createMockServerTransaction(t *testing.T, statusCode int, response interface{}) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		if response != nil {
			json.NewEncoder(w).Encode(response)
		}
	})
	return httptest.NewServer(handler)
}
