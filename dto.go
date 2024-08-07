package bovasdk

import "os"

// P2PTransactionRequest представляет тело запроса для создания P2P транзакции.
type P2PTransactionRequest struct {
	UserUUID         string `json:"user_uuid"`
	MerchantID       string `json:"merchant_id"`
	BankName         string `json:"bank_name"`
	Amount           int    `json:"amount"`
	CallbackURL      string `json:"callback_url"`
	RedirectURL      string `json:"redirect_url"`
	Email            string `json:"email"`
	CustomerName     string `json:"customer_name"`
	Currency         string `json:"currency"`
	PayeerIdentifier string `json:"payeer_identifier"`
	PayeerIP         string `json:"payeer_ip"`
	PayeerCardNumber string `json:"payeer_card_number"`
	PayeerType       string `json:"payeer_type"` // "trust" or "ftd"
	Lifetime         int    `json:"lifetime"`
	PaymentMethod    string `json:"payment_method"`
}

// P2PTransactionResponse представляет тело ответа API для создания P2P транзакции.
type P2PTransactionResponse struct {
	ResultCode string `json:"result_code"`
	Payload    struct {
		ID                string `json:"id"`
		MerchantID        string `json:"merchant_id"`
		Currency          string `json:"currency"`
		FormURL           string `json:"form_url"`
		State             string `json:"state"`
		CreatedAt         string `json:"created_at"`
		UpdatedAt         string `json:"updated_at"`
		CloseAt           string `json:"close_at"`
		CallbackURL       string `json:"callback_url"`
		RedirectURL       string `json:"redirect_url"`
		Email             string `json:"email"`
		CustomerName      string `json:"customer_name"`
		Rate              string `json:"rate"`
		Amount            string `json:"amount"`
		FiatAmount        string `json:"fiat_amount"`
		OldFiatAmount     string `json:"old_fiat_amount"`
		ServiceCommission string `json:"service_commission"`
		TotalAmount       string `json:"total_amount"`
		PaymentMethod     string `json:"payment_method"`
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
	} `json:"payload"`
}

// P2PTransactionResponseShort представляет тело ответа API для отмены или факта оплаты p2p транзакции.
type P2PTransactionResponseShort struct {
	ResultCode string `json:"result_code"`
	Payload    struct {
		ID            string `json:"id"`
		MerchantID    string `json:"merchant_id"`
		Amount        string `json:"amount"`
		OldAmount     string `json:"old_amount"`
		FormURL       string `json:"form_url"`
		State         string `json:"state"`
		CreatedAt     string `json:"created_at"`
		UpdatedAt     string `json:"updated_at"`
		CallbackURL   string `json:"callback_url"`
		RecipientCard struct {
			ID         int      `json:"id"`
			Number     string   `json:"number"`
			BankName   string   `json:"bank_name"`
			BankColors struct{} `json:"bank_colors"`
			Brand      string   `json:"brand"`
			UpdatedAt  string   `json:"updated_at"`
			CreatedAt  string   `json:"created_at"`
		} `json:"resipient_card"`
	} `json:"payload"`
}

// P2PDisputeRequest представляет тело запроса для создания диспута по p2p транзакции.
type P2PDisputeRequest struct {
	TransactionID string
	Amount        string
	file          *os.File
}

// P2PDisputeResponse представляет тело ответа API для создания диспута по p2p транзакции.
type P2PDisputeResponse struct {
	ID               int    `json:"id"`
	State            string `json:"state"`
	ProofImage       string `json:"proof_image"`
	ProofImage2      string `json:"proof_image2"`
	P2PTransactionID string `json:"p2p_transaction_id"`
	Repeated         bool   `json:"repeated"`
	Amount           int    `json:"amount"`
	UpdatedAt        string `json:"updated_at"`
	CreatedAt        string `json:"created_at"`
	RecipientCard    struct {
		ID         int      `json:"id"`
		Number     string   `json:"number"`
		BankName   string   `json:"bank_name"`
		BankColors struct{} `json:"bank_colors"`
		Brand      string   `json:"brand"`
		UpdatedAt  string   `json:"updated_at"`
		CreatedAt  string   `json:"created_at"`
	} `json:"resipient_card"`
}

// MassTransactionRequest представляет тело запроса для создания массовой транзакции.
type MassTransactionRequest struct {
	UserUUID      string `json:"user_uuid"`
	ToCard        string `json:"to_card"`
	Amount        int    `json:"amount"`
	CallbackURL   string `json:"callback_url"`
	MerchantID    string `json:"merchant_id"`
	Currency      string `json:"currency"`
	PaymentMethod string `json:"payment_method"`
	Lifetime      int    `json:"lifetime,omitempty"`
}

// MassTransactionResponse представляет тело ответа API для создания массовой транзакции.
type MassTransactionResponse struct {
	ResultCode string `json:"result_code"`
	Payload    struct {
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
	} `json:"payload"`
}
