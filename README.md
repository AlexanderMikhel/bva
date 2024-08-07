# BovaApi SDK для Go

## Установка

Установите последнюю версию SDK с помощью команды:

```go get github.com/yourusername/bovaapi```

## Использование

Создание экземпляра BovaApi
Для создания экземпляра BovaApi используйте билдера BovaApiBuilder:

```go
sdkBuilder := bovaapi.NewBovaApiBuilder().
ApiURL("https://google.com").
Secret("your_api_secret")

sdk, err := sdkBuilder.Build()
if err != nil {
    log.Fatalf("Error building SDK: %v", err)
}

```

## P2P Транзакции

### Создание P2P транзакции

```go
p2pRequest := bovaapi.P2PTransactionRequest{
UserUUID:         "364dbfc8-ae50-492f-bdd9-748edd84d5c9",
MerchantID:       "test7",
BankName:         "sberbank",
Amount:           500,
CallbackURL:      "https://webhook.site/callback",
RedirectURL:      "https://ya.ru/",
Email:            "test@mail.ru",
CustomerName:     "Ivan Vasiliev",
Currency:         "rub",
PayeerIdentifier: "payeer_identifier123",
PayeerIP:         "127.0.0.1",
PayeerCardNumber: "1234567890123456",
PayeerType:       "trust",
Lifetime:         1000,
PaymentMethod:    "card",
}

p2pResponse, err := sdk.P2P.CreateP2PTransaction(p2pRequest)
if err != nil {
    log.Fatalf("Error creating P2P transaction: %v", err)
}
```

### Получение информации о P2P транзакции

```go
transactionID := "9bb5f95f36e1e40d6b1376ed6ce5048172ebfdb7"
p2pResponse, err := sdk.P2P.GetP2PTransaction(transactionID)
if err != nil {
    log.Fatalf("Error getting P2P transaction: %v", err)
}
```

### Отмена P2P транзакции

```go
cancelResponse, err := sdk.P2P.CancelP2PTransaction(transactionID)
if err != nil {
    log.Fatalf("Error canceling P2P transaction: %v", err)
}
```

### Пометка P2P транзакции как оплаченной

```go
paidResponse, err := sdk.P2P.MarkP2PTransactionPaid(transactionID)
if err != nil {
    log.Fatalf("Error marking P2P transaction as paid: %v", err)
}
```

## Массовые Транзакции

### Создание массовой транзакции

```go
massTransactionRequest := bovaapi.MassTransactionRequest{
UserUUID:      "364dbfc8-ae50-492f-bdd9-748edd84d5c9",
ToCard:        "4111111111111111",
Amount:        200,
CallbackURL:   "https://webhook.site/callback",
MerchantID:    "test7",
Currency:      "rub",
PaymentMethod: "card",
Lifetime:      3600,
}

massTransactionResponse, err := sdk.MassTransaction.CreateMassTransaction(massTransactionRequest)
if err != nil {
    log.Fatalf("Error creating mass transaction: %v", err)
}
```

### Получение информации о массовой транзакции

```go
massTransactionID := "mock_transaction_id"
massTransactionResponse, err := sdk.MassTransaction.GetMassTransaction(massTransactionID)
if err != nil {
    log.Fatalf("Error getting mass transaction: %v", err)
}
```

## Опциональные настройки
### Логгирование

По умолчанию библиотека создает свой логгер с логами в формате json, логгер логгирует все входящие и исходящие запросы в
stdout,
Вы можете создать и настроить свой логгер реализовав интерфейс:

```go
type Logger interface {
	Enabled() bool

	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}
```

где параметр Enabled - отвечает за то, надо ли логгировать запросы и ответы или нет.
Вашу реализацию необзодимо подложить в NewBovaApiBuilder при сборке:

```go
sdkBuilder := bovasdk.NewBovaApiBuilder().
ApiURL("https://google.com").
Secret("your_api_secret").
Logger(myCustomLogger)


sdk, err := sdkBuilder.Build()
if err != nil {
    log.Fatalf("Error building SDK: %v", err)
}
```

### Клиент для запросов

Кроме того, вы можете полностю передать свою структуру http.Client со своими настройками(включая логгирование, таймауты и т.д.) для
выполнения запросов

```go
sdkBuilder := bovasdk.NewBovaApiBuilder().
ApiURL("https://google.com").
Secret("your_api_secret").
Client(myCustomCLient)


sdk, err := sdkBuilder.Build()
if err != nil {
    log.Fatalf("Error building SDK: %v", err)
}
```