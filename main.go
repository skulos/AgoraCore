package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/mattn/go-sqlite3"
// )

// func main() {

// 	db, err := sql.Open("sqlite3", ":memory:")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer db.Close()

// 	var version string
// 	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(version)
// }

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
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
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
