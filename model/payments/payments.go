package payments

type PaymentType string

const (
	Cash    PaymentType = "Cash"
	Card    PaymentType = "Card"
	Account PaymentType = "Account"
	Returns PaymentType = "Returns"
)
