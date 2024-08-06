package bovasdk

import (
	"fmt"
	"net/http"
	"time"
)

// BovaApi представляет структуру SDK с заголовками и URL.
type BovaApi struct {
	apiURL          string
	secret          string
	logger          Logger
	client          *http.Client
	P2P             *P2P
	MassTransaction *MassTransaction
}

// BovaApiBuilder помогает построить экземпляр BovaApi.
type BovaApiBuilder struct {
	apiURL string
	secret string
	client *http.Client
	logger Logger
}

// NewBovaApiBuilder создает новый экземпляр BovaApiBuilder.
func NewBovaApiBuilder() *BovaApiBuilder {
	return &BovaApiBuilder{}
}

// ApiURL устанавливает URL API.
func (b *BovaApiBuilder) ApiURL(apiURL string) *BovaApiBuilder {
	b.apiURL = apiURL
	return b
}

// Secret устанавливает заголовок Secret.
func (b *BovaApiBuilder) Secret(secret string) *BovaApiBuilder {
	b.secret = secret
	return b
}

// client устанавливает клиента для проведения настроек
func (b *BovaApiBuilder) Client(client *http.Client) *BovaApiBuilder {
	b.client = client
	return b
}

// client устанавливает логгер для логгирования запросов
func (b *BovaApiBuilder) Logger(logger Logger) *BovaApiBuilder {
	b.logger = logger
	return b
}

// Build строит и возвращает экземпляр BovaApi.
func (b *BovaApiBuilder) Build() (*BovaApi, error) {
	if b.secret == "" {
		return nil, fmt.Errorf("secret is required")
	}

	if b.apiURL == "" {
		return nil, fmt.Errorf("api URL is required")
	}

	var err error
	if b.logger == nil {
		//по дефолту логгер активен в режиме info
		b.logger, err = NewLogger(true, "info")
		if err != nil {
			return nil, fmt.Errorf("cant build default logger: %s", err.Error())
		}
	}

	if b.client == nil {
		//по дефолту создается кастомный клиент с логом
		b.client = &http.Client{
			Transport: NewLoggingRoundTripper(b.logger, http.DefaultTransport),
			Timeout:   30 * time.Second,
		}
	}

	return &BovaApi{
		apiURL:          b.apiURL,
		secret:          b.secret,
		client:          b.client,
		logger:          b.logger,
		P2P:             p2pNew(b.apiURL, b.secret, b.client),
		MassTransaction: massTransactionNew(b.apiURL, b.secret, b.client),
	}, nil
}
