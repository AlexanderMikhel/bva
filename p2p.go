package bovasdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

type P2P struct {
	apiURL string
	secret string
	client *http.Client
}

func p2pNew(apiURL, secret string, client *http.Client) *P2P {
	return &P2P{apiURL: apiURL, secret: secret, client: client}
}

// CreateP2PTransaction создает платеж p2p и получает ссылку на пополнение.
func (p2p *P2P) CreateP2PTransaction(req P2PTransactionRequest) (*P2PTransactionResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	// Создаем HTTP запрос
	httpReq, err := http.NewRequest(http.MethodPost, p2p.apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Устанавливаем заголовки
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set(signatureHeader, calculateSignature(p2p.secret, jsonData))

	// Отправляем запрос
	resp, err := p2p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Обрабатываем ответ
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %v", resp.StatusCode)
	}

	var response P2PTransactionResponse
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &response, nil
}

func (p2p *P2P) MarkP2PTransactionPaid(transactionID string) (*P2PTransactionResponseShort, error) {
	url := fmt.Sprintf("%s/p2p_transactions/%s/paid", p2p.apiURL, transactionID)
	httpReq, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	resp, err := p2p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %v", resp.StatusCode)
	}

	var response P2PTransactionResponseShort
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &response, nil
}

// CancelP2PTransaction отменяет p2p транзакцию по её ID.
func (p2p *P2P) CancelP2PTransaction(transactionID string) (*P2PTransactionResponseShort, error) {
	url := fmt.Sprintf("%s/p2p_transactions/%s/cancel", p2p.apiURL, transactionID)
	httpReq, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	resp, err := p2p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %v", resp.StatusCode)
	}

	var response P2PTransactionResponseShort
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &response, nil
}

// GetP2PTransaction получает информацию о p2p транзакции по её ID.
func (p2p *P2P) GetP2PTransaction(transactionID string) (*P2PTransactionResponse, error) {
	url := fmt.Sprintf("%s/p2p_transactions/%s", p2p.apiURL, transactionID)
	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	resp, err := p2p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %v", resp.StatusCode)
	}

	var response P2PTransactionResponse
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &response, nil
}

// CreateP2PDispute создаем диспут по p2p транзакции.
func (p2p *P2P) CreateP2PDispute(req P2PDisputeRequest) (*P2PDisputeResponse, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField(transactionIdForm, req.TransactionID)
	_ = writer.WriteField(p2pDisputeAmountForm, req.Amount)

	part, err := writer.CreateFormFile(p2pDisputeProofImageForm, filepath.Base(req.file.Name()))
	if err != nil {
		return nil, fmt.Errorf("error creating form file with name: %s, err: %v", req.file.Name(), err)
	}
	defer req.file.Close()
	_, err = io.Copy(part, req.file)
	if err != nil {
		return nil, fmt.Errorf("error copying file: %v", err)
	}

	if err = writer.Close(); err != nil {
		return nil, fmt.Errorf("error closing writer: %v", err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/p2p_disputes/from_client", p2p.apiURL), body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	httpReq.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := p2p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response code: %v", resp.StatusCode)
	}

	var response P2PDisputeResponse
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &response, nil
}
