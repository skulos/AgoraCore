package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/mr-tron/base58"
	// "github.com/mr-tron/base58/base58"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "agoracore"
)

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	createStmt := `create table if not exists items(id serial primary key, name text not null, price int not null)`
	_, e := db.Exec(createStmt)
	CheckError(e)

	insertDynStmt := `insert into "items"("name", "price") values($1, $2)`
	rand.Seed(time.Now().Unix())

	for i := 0; i < 10000; i++ {
		time := time.Now().Unix() + (rand.Int63() % 1000)
		bytes := []byte(strconv.Itoa(int(time)))
		name := base58.Encode(bytes)
		price := rand.Int() % 500
		_, e = db.Exec(insertDynStmt, name, price)
		fmt.Println("i : ", i, "name(", name, ")", " price(", price, ")")
		CheckError(e)
	}

	// insert
	// hardcoded
	// insertStmt := `insert into "students"("id","name", "age") values(0,'John', 1)`
	// _, e = db.Exec(insertStmt)
	// CheckError(e)

	// // dynamic
	// insertDynStmt := `insert into "students"("id","name", "age") values($1, $2, $3)`
	// _, e = db.Exec(insertDynStmt, 1, "Jane", 2)
	// CheckError(e)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
