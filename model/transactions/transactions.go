package transactions

import (
	"encoding/json"
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
	TxID        int64 `gorm:"primaryKey;autoIncrement;"`
	Date        string
	Time        string
	PaymentType payments.PaymentType //`gorm:"embedded"`
	TxList      string               //[]items.Item
	Total       string               //decimal.Decimal
	// CreatedAt   time.Time
}

func NewTrx(paymentMethod payments.PaymentType, list []items.Item, total decimal.Decimal) Transaction {
	timeNow := time.Now()

	// y, m, d := timeNow.Date()
	// date := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	date := timeNow.Format("01-02-2006")
	// dateLayout := "Mon Jan 02 2006 15:04:05 GMT-0700"

	time := timeNow.Format("15:04:05 PM")
	txId := timeNow.Unix()

	txlistJSON, _ := json.Marshal(list)

	totalString := total.StringFixed(2)

	trans := Transaction{
		TxID:        txId,
		Date:        date,
		Time:        time,
		PaymentType: paymentMethod,
		TxList:      string(txlistJSON), //list,
		Total:       totalString,        //total,
	}

	return trans
}
