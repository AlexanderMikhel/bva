package bovasdk

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Mock API secret
const apiSecret = "mock_api_secret"

// Mock request and response data
var (
	mockTransactionID         = "mock_transaction_id"
	mockP2PTransactionRequest = P2PTransactionRequest{
		UserUUID:         "mock_user_uuid",
		MerchantID:       "mock_merchant_id",
		Amount:           100,
		CallbackURL:      "https://example.com/callback",
		RedirectURL:      "https://example.com/redirect",
		Currency:         "USD",
		PayeerIdentifier: "mock_identifier",
		PayeerIP:         "127.0.0.1",
		PayeerType:       "trust",
		PaymentMethod:    "card",
	}

	mockP2PTransactionResponse = P2PTransactionResponse{
		ResultCode: "ok",
		Payload: struct {
			ID                string               `json:"id"`
			MerchantID        string               `json:"merchant_id"`
			Currency          CurrencyEnum         `json:"currency"`
			FormURL           string               `json:"form_url"`
			State             TransactionStateEnum `json:"state"`
			CreatedAt         string               `json:"created_at"`
			UpdatedAt         string               `json:"updated_at"`
			CloseAt           string               `json:"close_at"`
			CallbackURL       string               `json:"callback_url"`
			RedirectURL       string               `json:"redirect_url"`
			Email             string               `json:"email"`
			CustomerName      string               `json:"customer_name"`
			Rate              string               `json:"rate"`
			Amount            string               `json:"amount"`
			FiatAmount        string               `json:"fiat_amount"`
			OldFiatAmount     string               `json:"old_fiat_amount"`
			ServiceCommission string               `json:"service_commission"`
			TotalAmount       string               `json:"total_amount"`
			PaymentMethod     PaymentMethodEnum    `json:"payment_method"`
			RecipientCard     struct {
				ID            string   `json:"id"`
				Number        string   `json:"number"`
				BankName      string   `json:"bank_name"`
				BankFullName  string   `json:"bank_full_name"`
				BankColors    struct{} `json:"bank_colors"`
				Brand         string   `json:"brand"`
				CardHolder    string   `json:"card_holder"`
				PaymentMethod string   `json:"payment_method"`
				UpdatedAt     string   `json:"updated_at"`
				CreatedAt     string   `json:"created_at"`
				SberpayURL    string   `json:"sberpay_url"`
			} `json:"resipient_card"`
		}{
			ID:                "mock_id",
			MerchantID:        "mock_merchant_id",
			Currency:          UZS,
			FormURL:           "https://example.com/form",
			State:             Successed,
			CreatedAt:         "2023-01-01T00:00:00Z",
			UpdatedAt:         "2023-01-01T00:00:00Z",
			CloseAt:           "2023-01-01T01:00:00Z",
			CallbackURL:       "https://example.com/callback",
			RedirectURL:       "https://example.com/redirect",
			Email:             "test@example.com",
			CustomerName:      "Test Customer",
			Rate:              "1.0",
			Amount:            "100.0",
			FiatAmount:        "100.0",
			OldFiatAmount:     "100.0",
			ServiceCommission: "1.0",
			TotalAmount:       "99.0",
			PaymentMethod:     Card,
			RecipientCard: struct {
				ID            string   `json:"id"`
				Number        string   `json:"number"`
				BankName      string   `json:"bank_name"`
				BankFullName  string   `json:"bank_full_name"`
				BankColors    struct{} `json:"bank_colors"`
				Brand         string   `json:"brand"`
				CardHolder    string   `json:"card_holder"`
				PaymentMethod string   `json:"payment_method"`
				UpdatedAt     string   `json:"updated_at"`
				CreatedAt     string   `json:"created_at"`
				SberpayURL    string   `json:"sberpay_url"`
			}{
				ID:            "mock_card_id",
				Number:        "123456******5678",
				BankName:      "Mock Bank",
				BankFullName:  "Mock Bank Full Name",
				BankColors:    struct{}{},
				Brand:         "visa",
				CardHolder:    "Test Customer",
				PaymentMethod: "card",
				UpdatedAt:     "2023-01-01T00:00:00Z",
				CreatedAt:     "2023-01-01T00:00:00Z",
				SberpayURL:    "https://example.com/sberpay",
			},
		},
	}
)

// Helper function to create a mock server
func createMockServerP2P(t *testing.T, statusCode int, response interface{}) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		if response != nil {
			json.NewEncoder(w).Encode(response)
		}
	})
	return httptest.NewServer(handler)
}

// TestCreateP2PTransaction tests the CreateP2PTransaction method
func TestCreateP2PTransaction(t *testing.T) {
	server := createMockServerP2P(t, http.StatusOK, mockP2PTransactionResponse)
	defer server.Close()

	p2p := p2pNew(server.URL, apiSecret, server.Client())
	resp, err := p2p.CreateP2PTransaction(mockP2PTransactionRequest)
	if err != nil {
		t.Fatalf("CreateP2PTransaction() error = %v", err)
	}

	if resp.ResultCode != mockP2PTransactionResponse.ResultCode {
		t.Errorf("CreateP2PTransaction() ResultCode = %v, want %v", resp.ResultCode, mockP2PTransactionResponse.ResultCode)
	}
}

// TestCreateP2PTransaction_Error tests the CreateP2PTransaction method for error response
func TestCreateP2PTransaction_Error(t *testing.T) {
	server := createMockServerP2P(t, http.StatusBadRequest, nil)
	defer server.Close()

	p2p := p2pNew(server.URL, apiSecret, server.Client())
	_, err := p2p.CreateP2PTransaction(mockP2PTransactionRequest)
	if err == nil {
		t.Fatalf("CreateP2PTransaction() expected error, got nil")
	}
}

// TestMarkP2PTransactionPaid tests the MarkP2PTransactionPaid method
func TestMarkP2PTransactionPaid(t *testing.T) {
	server := createMockServerP2P(t, http.StatusOK, mockP2PTransactionResponse)
	defer server.Close()

	p2p := p2pNew(server.URL, apiSecret, server.Client())
	resp, err := p2p.MarkP2PTransactionPaid(mockTransactionID)
	if err != nil {
		t.Fatalf("MarkP2PTransactionPaid() error = %v", err)
	}

	if resp.ResultCode != mockP2PTransactionResponse.ResultCode {
		t.Errorf("MarkP2PTransactionPaid() ResultCode = %v, want %v", resp.ResultCode, mockP2PTransactionResponse.ResultCode)
	}
}

// TestMarkP2PTransactionPaid_Error tests the MarkP2PTransactionPaid method for error response
func TestMarkP2PTransactionPaid_Error(t *testing.T) {
	server := createMockServerP2P(t, http.StatusBadRequest, nil)
	defer server.Close()

	p2p := p2pNew(server.URL, apiSecret, server.Client())
	_, err := p2p.MarkP2PTransactionPaid(mockTransactionID)
	if err == nil {
		t.Fatalf("MarkP2PTransactionPaid() expected error, got nil")
	}
}

// TestCancelP2PTransaction tests the CancelP2PTransaction method
func TestCancelP2PTransaction(t *testing.T) {
	server := createMockServerP2P(t, http.StatusOK, mockP2PTransactionResponse)
	defer server.Close()

	p2p := p2pNew(server.URL, apiSecret, server.Client())
	resp, err := p2p.CancelP2PTransaction(mockTransactionID)
	if err != nil {
		t.Fatalf("CancelP2PTransaction() error = %v", err)
	}

	if resp.ResultCode != mockP2PTransactionResponse.ResultCode {
		t.Errorf("CancelP2PTransaction() ResultCode = %v, want %v", resp.ResultCode, mockP2PTransactionResponse.ResultCode)
	}
}

// TestCancelP2PTransaction_Error tests the CancelP2PTransaction method for error response
func TestCancelP2PTransaction_Error(t *testing.T) {
	server := createMockServerP2P(t, http.StatusBadRequest, nil)
	defer server.Close()

	p2p := p2pNew(server.URL, apiSecret, server.Client())
	_, err := p2p.CancelP2PTransaction(mockTransactionID)
	if err == nil {
		t.Fatalf("CancelP2PTransaction() expected error, got nil")
	}
}

// TestGetP2PTransaction tests the GetP2PTransaction method
func TestGetP2PTransaction(t *testing.T) {
	server := createMockServerP2P(t, http.StatusOK, mockP2PTransactionResponse)
	defer server.Close()

	p2p := p2pNew(server.URL, apiSecret, server.Client())
	resp, err := p2p.GetP2PTransaction(mockTransactionID)
	if err != nil {
		t.Fatalf("GetP2PTransaction() error = %v", err)
	}

	if resp.ResultCode != mockP2PTransactionResponse.ResultCode {
		t.Errorf("GetP2PTransaction() ResultCode = %v, want %v", resp.ResultCode, mockP2PTransactionResponse.ResultCode)
	}
}

// TestGetP2PTransaction_Error tests the GetP2PTransaction method for error response
func TestGetP2PTransaction_Error(t *testing.T) {
	server := createMockServerP2P(t, http.StatusBadRequest, nil)
	defer server.Close()

	p2p := p2pNew(server.URL, apiSecret, server.Client())
	_, err := p2p.GetP2PTransaction(mockTransactionID)
	if err == nil {
		t.Fatalf("GetP2PTransaction() expected error, got nil")
	}
}

// TestCreateP2PDispute tests the CreateP2PDispute method
func TestCreateP2PDispute(t *testing.T) {
	// Mock file for upload
	fileContent := []byte("mock file content")
	_ = bytes.NewReader(fileContent)
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(fileContent); err != nil {
		t.Fatal(err)
	}
	if _, err := tmpfile.Seek(0, 0); err != nil {
		t.Fatal(err)
	}

	mockP2PDisputeRequest := NewP2PDisputeRequest("mock_transaction_id", "1000", 1000, tmpfile)
	server := createMockServerP2P(t, http.StatusOK, P2PDisputeResponse{
		ID:               1,
		State:            "opened",
		ProofImage:       "https://example.com/proof_image.jpg",
		P2PTransactionID: "mock_transaction_id",
		Amount:           1000,
	})
	defer server.Close()

	p2p := p2pNew(server.URL, apiSecret, server.Client())
	resp, err := p2p.CreateP2PDispute(*mockP2PDisputeRequest)
	if err != nil {
		t.Fatalf("CreateP2PDispute() error = %v", err)
	}

	if resp.State != "opened" {
		t.Errorf("CreateP2PDispute() State = %v, want %v", resp.State, "opened")
	}
}

// TestCreateP2PDispute_Error tests the CreateP2PDispute method for error response
func TestCreateP2PDispute_Error(t *testing.T) {
	// Mock file for upload
	fileContent := []byte("mock file content")
	_ = bytes.NewReader(fileContent)
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(fileContent); err != nil {
		t.Fatal(err)
	}
	if _, err := tmpfile.Seek(0, 0); err != nil {
		t.Fatal(err)
	}

	mockP2PDisputeRequest := NewP2PDisputeRequest("mock_transaction_id", "1000", 1000, tmpfile)
	server := createMockServerP2P(t, http.StatusOK, P2PDisputeResponse{
		ID:               1,
		State:            "opened",
		ProofImage:       "https://example.com/proof_image.jpg",
		P2PTransactionID: "mock_transaction_id",
		Amount:           1000,
	})
	defer server.Close()

	p2p := p2pNew(server.URL, apiSecret, server.Client())
	_, err = p2p.CreateP2PDispute(*mockP2PDisputeRequest)
	if err == nil {
		t.Fatalf("CreateP2PDispute() expected error, got nil")
	}
}
