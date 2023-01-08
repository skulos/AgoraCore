package items

import "github.com/shopspring/decimal"

type ItemInterface interface {
	// ID function to get ID allocation of item
	ID() int64

	// Descripton reading and writing
	GetDescription() string
	SetDescription(description string) error

	// UnitPrice
	GetUnitCost() decimal.Decimal
	SetUnitCost(price decimal.Decimal) error
}

type Item struct {
	ID          int64
	QTY         int64
	Description string
	UnitCost    string //decimal.Decimal
	Price       string //decimal.Decimal
}

func NewItem(id int64, qty int64, description string, unit decimal.Decimal, price decimal.Decimal) Item {

	uc := unit.StringFixed(2)
	p := price.StringFixed(2)

	return Item{
		ID:          id,
		QTY:         qty,
		Description: description,
		UnitCost:    uc,
		Price:       p,
	}
}

// type Stock struct{}
