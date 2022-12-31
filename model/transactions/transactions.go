package transactions

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/skulos/AgoraCore/model/items"
	"github.com/skulos/AgoraCore/model/payments"
)

// type TxList []items.Item

// // Also called the Stock Delta
// type TxContext struct {
// }

/*
ID
Date
Time
PaymentType
List of Items
Total
*/
// DateTime time.Time
type Transaction struct {
	TxID        int64
	Date        string
	Time        string
	PaymentType payments.PaymentType
	TxList      []items.Item // TxList
	Total       decimal.Decimal
}

func NewTrx(paymentMethod payments.PaymentType, list []items.Item, total decimal.Decimal) Transaction {
	timeNow := time.Now()

	date := timeNow.Format("01-02-2006")
	time := timeNow.Format("15:04:05")
	txId := timeNow.Unix()

	trans := Transaction{
		TxID:        txId,
		Date:        date,
		Time:        time,
		PaymentType: paymentMethod,
		TxList:      list,
		Total:       total,
	}

	return trans
}
