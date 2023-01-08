package db

import (
	"github.com/skulos/AgoraCore/model/payments"
	"github.com/skulos/AgoraCore/model/transactions"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

const (
	January string = "January"
)

func DBInit() {
	// var err error
	var dsn string = "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable"

	db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// automigrate
	// Migrate the schema
	db.AutoMigrate(&transactions.Transaction{})

	db.Create(&transactions.Transaction{
		Date:        "07-01-2023",
		Time:        "06:10 PM",
		PaymentType: payments.Cash,
		TxList:      "[]",
		Total:       "0.00",
	})

}
