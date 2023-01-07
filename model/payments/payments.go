package payments

type PaymentType uint

const (
	Cash PaymentType = iota + 1
	Card
	Account
	Returns
)
