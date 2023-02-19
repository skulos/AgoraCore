package types

type Type interface {
	FromString(t string) (Typetype, error)
	ToString() string
}

type Typetype string

// Data Types

// Account Types
type PaymentType = Typetype

func (p PaymentType) ToString() string {
	return string(p)
}

func (p PaymentType) FromString(t string) (PaymentType, error) {
	switch t {
		case
	}
}

func Test() {
	var interfaces Type = Card
}

const (
	Cash    PaymentType = "Cash"
	Card    PaymentType = "Card"
	Account PaymentType = "Account"
	Returns PaymentType = "Returns"
)

type AccountType string

const (
	Assets      AccountType = "Assets"      // + | -
	Capitals    AccountType = "Capitals"    // - | +
	Expenses    AccountType = "Expenses"    // + | -
	Income      AccountType = "Income"      // - | +
	Liabilities AccountType = "Liabilities" // - | +
)
