package main

import (
	"encoding/json"
	"log"
	"time"
)

// type AccountType struct {
// 	Id int
// 	Amount int
// }

// func (at AccountType) Id() int {}
// func (at AccountType) Debit(amount int) error {}
// func (at AccountType) Credit(amount int) error {}
// func (at AccountType) Balance() int {}
// func (at AccountType) Summery() string {}

type BalanceCalcFunctions func(int, int) int

var (
	DebitBalanceType BalanceCalcFunctions = func(d int, c int) int {
		return d - c
	}

	CreditBalanceType BalanceCalcFunctions = func(d int, c int) int {
		return c - d
	}
)

// type Account interface {
// 	AccId() int
// 	Debit(entry AccountEntries)
// 	Credit(entry AccountEntries)
// 	Balance() int
// 	Summery() string
// }

type AccountType struct {
	Id            int
	DebitEntries  []AccountEntries
	CreditEntries []AccountEntries
	// balanceCalc   func(int, int) int
	BalanceCalc BalanceCalcFunctions
}

type AccountEntries struct {
	Date        string
	Description string
	Reference   string
	Account     int
}

func (at *AccountType) AccId() int {
	return at.Id
}

func (at *AccountType) Debit(entry AccountEntries) {
	at.DebitEntries = append(at.DebitEntries, entry)
}

func (at *AccountType) Credit(entry AccountEntries) {
	at.CreditEntries = append(at.CreditEntries, entry)
}

func (at *AccountType) Balance() int {
	// var totalDebit int = 0
	// var totalCredit int = 0
	totalDebit := 0
	totalCredit := 0

	for _, entry := range at.DebitEntries {
		totalDebit = totalDebit + entry.Account
	}

	for _, entry := range at.CreditEntries {
		totalCredit = totalCredit + entry.Account
	}

	log.Println("Balances:\nD = ", totalDebit, "\nC = ", totalCredit)
	amount := at.BalanceCalc(totalDebit, totalCredit)

	return amount
}

func (at *AccountType) Summery() string {
	jsonStr, _ := json.Marshal(at)
	return string(jsonStr)
}

func main() {
	generalLedger := make(map[int]AccountType)

	bankAcc := AccountType{
		Id:          0,
		BalanceCalc: DebitBalanceType,
	}

	loanAcc := AccountType{
		Id:          1,
		BalanceCalc: CreditBalanceType,
	}

	generalLedger[bankAcc.AccId()] = bankAcc
	generalLedger[loanAcc.AccId()] = loanAcc

	date := time.Now().Format("02-01-2006")

	entry1 := AccountEntries{
		Date:        date,
		Description: "Entry 1",
		Account:     1000,
	}
	entry2 := AccountEntries{
		Date:        date,
		Description: "Entry 2",
		Account:     100,
	}
	entry3 := AccountEntries{
		Date:        date,
		Description: "Entry 3",
		Account:     250,
	}

	bankAcc.Debit(entry1)
	bankAcc.Credit(entry2)
	bankAcc.Debit(entry3)

	log.Println("BankAcc:\n", bankAcc)
	balance := bankAcc.Balance()

	log.Println("Balance:\n", balance)

}
