package bovasdk

type CurrencyEnum string
type PaymentMethodEnum string
type TransactionStateEnum string

const (
	RUB CurrencyEnum = "rub"
	UZS CurrencyEnum = "uzs"

	Card    PaymentMethodEnum = "card"
	SberPay PaymentMethodEnum = "sberpay"
	Sbp     PaymentMethodEnum = "sbp"

	WaitingPayment            TransactionStateEnum = "waiting_payment"
	Paid                      TransactionStateEnum = "paid"
	Failed                    TransactionStateEnum = "failed"
	ClosedFailed              TransactionStateEnum = "closed_failed"
	RepeatedClosedFailed      TransactionStateEnum = "repeated_closed_failed"
	Successed                 TransactionStateEnum = "successed"
	AcceptedSuccessed         TransactionStateEnum = "accepted_successed"
	RepeatedAcceptedSuccessed TransactionStateEnum = "repeated_accepted_successed"
	Reviewing                 TransactionStateEnum = "reviewing"
	RepeatedReviewing         TransactionStateEnum = "repeated_reviewing"
)
