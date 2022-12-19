package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

type DBConfig struct {
	Host     string
	Port     int64
	UserName string
	Password string
	DBName   string
}

type DBTableCreationTransaction struct {
	month string
	year  string
}

func (dbtct DBTableCreationTransaction) IntoSQL() string {
	// default transation schema
	tableName := dbtct.month + dbtct.year
	tableCreateString := "CREATE TABLE IF NOT EXISTS " + tableName
	tableColumnString := "( )" // TODO: finsh off the creation command
}

func generateTransactionTable() DBTableCreationTransaction {
	y, m, _ := time.Now().Date()

	yearString := strconv.Itoa(y)
	monthString := m.String()

	return DBTableCreationTransaction{
		year:  yearString,
		month: monthString,
	}
}

type DBManagerType struct {
	conn *sql.DB
}

var DBManager DBManagerType

func DBInit() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	// defer db.Close()

	// check db
	// err = db.Ping()
	// CheckError(err)

	// fmt.Println("Connected!")

	DBManager.conn = db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
