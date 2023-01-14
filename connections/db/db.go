package db

import (
	"strconv"
	"time"

	"github.com/skulos/AgoraCore/model/payments"
	"github.com/skulos/AgoraCore/model/transactions"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

const (
	January   string = "January"
	February  string = "February"
	March     string = "March"
	May       string = "May"
	June      string = "June"
	July      string = "July"
	August    string = "August"
	September string = "September"
	October   string = "October"
	November  string = "November"
	December  string = "December"
)

var months [12]string = [12]string{
	"january",
	"february",
	"march",
	"april",
	"may",
	"june",
	"july",
	"august",
	"september",
	"october",
	"november",
	"december",
}

// https://gorm.io/docs/conventions.html#Temporarily-specify-a-name
func TableNames() {

	yearInt := time.Now().Year()
	yearString := strconv.Itoa(yearInt) + "_"

	for i := 0; i < 12; i++ {
		tableName := yearString + months[i]
		db.Table(tableName).AutoMigrate(&transactions.Transaction{})
	}

}

func DBInit() {
	// var err error
	var dsn string = "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable"

	db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// automigrate
	// Migrate the schema
	// db.AutoMigrate(&transactions.Transaction{})

	db.Create(&transactions.Transaction{
		Date:        "07-01-2023",
		Time:        "06:10 PM",
		PaymentType: payments.Cash,
		TxList:      "[]",
		Total:       "0.00",
	})

	TableNames()
	// db.Delete()

}
