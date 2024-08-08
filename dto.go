package bovasdk

import (
	"mime/multipart"
)

// P2PTransactionRequest представляет тело запроса для создания P2P транзакции.
type P2PTransactionRequest struct {
	UserUUID         string            `json:"user_uuid"`
	MerchantID       string            `json:"merchant_id"`
	PayeerIdentifier string            `json:"payeer_identifier"`
	PayeerIP         string            `json:"payeer_ip"`
	PayeerType       string            `json:"payeer_type"`
	Currency         CurrencyEnum      `json:"currency"`
	PaymentMethod    PaymentMethodEnum `json:"payment_method"`
	Amount           int               `json:"amount"`
	CallbackURL      string            `json:"callback_url"`
	RedirectURL      string            `json:"redirect_url"`
	//not required
	Email            *string `json:"email"`
	CustomerName     *string `json:"customer_name"`
	PayeerCardNumber *string `json:"payeer_card_number"`
}

// NewP2PTransactionRequest создает новый экземпляр P2PTransactionRequest с обязательными параметрами.
func NewP2PTransactionRequest(userUUID, merchantID, payeerIdentifier, payeerIP, payeerType, callbackURL, redirectURL string, currency CurrencyEnum, paymentMethod PaymentMethodEnum, amount int) *P2PTransactionRequest {
	return &P2PTransactionRequest{
		UserUUID:         userUUID,
		MerchantID:       merchantID,
		PayeerIdentifier: payeerIdentifier,
		PayeerIP:         payeerIP,
		PayeerType:       payeerType,
		Currency:         currency,
		PaymentMethod:    paymentMethod,
		Amount:           amount,
		CallbackURL:      callbackURL,
		RedirectURL:      redirectURL,
	}
}

// WithEmail задает email и возвращает обновленный запрос
func (p *P2PTransactionRequest) WithEmail(email string) *P2PTransactionRequest {
	p.Email = &email
	return p
}

// WithCustomerName задает имя клиента и возвращает обновленный запрос
func (p *P2PTransactionRequest) WithCustomerName(customerName string) *P2PTransactionRequest {
	p.CustomerName = &customerName
	return p
}

// WithPayeerCardNumber задает номер карты и возвращает обновленный запрос
func (p *P2PTransactionRequest) WithPayeerCardNumber(payeerCardNumber string) *P2PTransactionRequest {
	p.PayeerCardNumber = &payeerCardNumber
	return p
}

// P2PTransactionResponse представляет тело ответа API для создания P2P транзакции.
type P2PTransactionResponse struct {
	ResultCode string `json:"result_code"`
	Payload    struct {
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
	} `json:"payload"`
}

// P2PDisputeRequest представляет тело запроса для создания диспута по p2p транзакции.
type sdkFile struct {
	file multipart.File
	Name string
}

type P2PDisputeRequest struct {
	TransactionID string
	Amount        string
	ProofImage    sdkFile
	//not required
	ProofImage2 *sdkFile
}

func NewP2PDisputeRequest(TransactionID, Amount, fileName string, file multipart.File) *P2PDisputeRequest {
	return &P2PDisputeRequest{TransactionID: TransactionID, Amount: Amount, ProofImage: sdkFile{Name: fileName, file: file}}
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
	UserUUID      string            `json:"user_uuid"`
	MerchantID    string            `json:"merchant_id"`
	Amount        int               `json:"amount"`
	CallbackURL   string            `json:"callback_url"`
	Currency      CurrencyEnum      `json:"currency"`
	PaymentMethod PaymentMethodEnum `json:"payment_method"`

	ToCard      *string `json:"to_card"`
	SbpBankName *string `json:"sbp_bank_name"`
}

// NewMassTransactionRequest создает новый экземпляр MassTransactionRequest с обязательными параметрами.
func NewMassTransactionRequest(userUUID, merchantID, callbackURL string, amount int, currency CurrencyEnum, paymentMethod PaymentMethodEnum) *MassTransactionRequest {
	return &MassTransactionRequest{
		UserUUID:      userUUID,
		MerchantID:    merchantID,
		Amount:        amount,
		CallbackURL:   callbackURL,
		Currency:      currency,
		PaymentMethod: paymentMethod,
	}
}

// WithToCard задает карту получателя и возвращает обновленный запрос.
func (m *MassTransactionRequest) WithToCard(toCard string) *MassTransactionRequest {
	m.ToCard = &toCard
	return m
}

// WithSbpBankName задает название банка для SBP и возвращает обновленный запрос.
func (m *MassTransactionRequest) WithSbpBankName(sbpBankName string) *MassTransactionRequest {
	m.SbpBankName = &sbpBankName
	return m
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
