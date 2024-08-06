package bovasdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MassTransaction struct {
	apiURL string
	secret string
	client *http.Client
}

func massTransactionNew(apiURL, secret string, client *http.Client) *MassTransaction {
	return &MassTransaction{apiURL: apiURL, secret: secret, client: client}
}

// CreateMassTransaction создает заявку на выплату на карту.
func (mt *MassTransaction) CreateMassTransaction(req MassTransactionRequest) (*MassTransactionResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/mass_transactions", mt.apiURL), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set(signatureHeader, calculateSignature(mt.secret, jsonData))

	resp, err := mt.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %v", resp.StatusCode)
	}

	var response MassTransactionResponse
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &response, nil
}

// GetMassTransaction получает информацию о транзакции по её ID.
func (mt *MassTransaction) GetMassTransaction(transactionID string) (*MassTransactionResponse, error) {
	url := fmt.Sprintf("%s/mass_transactions/%s", mt.apiURL, transactionID)
	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	resp, err := mt.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %v", resp.StatusCode)
	}

	var response MassTransactionResponse
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &response, nil
}
