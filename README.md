BovaApi SDK для Go
BovaApi SDK предоставляет удобный способ интеграции с API BovaTech для работы с p2p и массовыми транзакциями. SDK
включает методы для создания и управления транзакциями, а также для создания диспутов.

Установка
Установите последнюю версию SDK с помощью команды:

```go get github.com/yourusername/bovaapi```

Использование
Создание экземпляра BovaApi
Для создания экземпляра BovaApi используйте билдера BovaApiBuilder:

```
package main

import (
"fmt"
"log"

	"github.com/yourusername/bovaapi"
)

func main() {
sdkBuilder := bovaapi.NewBovaApiBuilder().
ApiURL("https://bovatech.cc/v1").
Secret("your_api_secret")

	sdk, err := sdkBuilder.Build()
	if err != nil {
		log.Fatalf("Error building SDK: %v", err)
	}

	fmt.Println("SDK успешно создан!")
}
```

P2P Транзакции
Создание P2P транзакции
```
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

fmt.Printf("P2P Transaction Response: %+v\n", p2pResponse)
```

Получение информации о P2P транзакции
```
transactionID := "9bb5f95f36e1e40d6b1376ed6ce5048172ebfdb7"
p2pResponse, err := sdk.P2P.GetP2PTransaction(transactionID)
if err != nil {
log.Fatalf("Error getting P2P transaction: %v", err)
}

fmt.Printf("P2P Transaction Details: %+v\n", p2pResponse)
```

Отмена P2P транзакции
```
cancelResponse, err := sdk.P2P.CancelP2PTransaction(transactionID)
if err != nil {
log.Fatalf("Error canceling P2P transaction: %v", err)
}

fmt.Printf("Cancel P2P Transaction Response: %+v\n", cancelResponse)
```

Пометка P2P транзакции как оплаченной
```
paidResponse, err := sdk.P2P.MarkP2PTransactionPaid(transactionID)
if err != nil {
log.Fatalf("Error marking P2P transaction as paid: %v", err)
}

fmt.Printf("Mark P2P Transaction Paid Response: %+v\n", paidResponse)
```

Массовые Транзакции
Создание массовой транзакции

```
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

fmt.Printf("Mass Transaction Response: %+v\n", massTransactionResponse)
```

Получение информации о массовой транзакции
```
massTransactionID := "mock_transaction_id"
massTransactionResponse, err := sdk.MassTransaction.GetMassTransaction(massTransactionID)
if err != nil {
	log.Fatalf("Error getting mass transaction: %v", err)
}

fmt.Printf("Mass Transaction Details: %+v\n", massTransactionResponse)
```