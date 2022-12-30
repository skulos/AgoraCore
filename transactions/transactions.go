package transactions

import (
	"github.com/shopspring/decimal"
	"github.com/skulos/AgoraCore/items"
)

type PaymentType uint

const (
	Cash PaymentType = iota + 1
	Card
	Account
	Returns
)

type TxList []items.Item

// Also called the Stock Delta
type TxContext struct {
}

/*
ID
Date
Time
PaymentType
List of Items
Total
*/
type Transactions struct {
	TxID string
	// DateTime time.Time
	Date        string
	Time        int64
	PaymentType PaymentType
	TxList      TxList
	Total       decimal.Decimal
}

type TxContextSaleSummary struct {
	Sales decimal.Decimal // numbers
	Card  decimal.Decimal
	Cash  decimal.Decimal
}

type TxContextReturnsSummary struct {
	Returned decimal.Decimal
	Card     decimal.Decimal
	Cash     decimal.Decimal
}
