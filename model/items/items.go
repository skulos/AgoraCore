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
	Description string
	UnitCost    decimal.Decimal
	Price       decimal.Decimal
}

// type Stock struct{}
