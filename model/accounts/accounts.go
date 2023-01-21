package accounts

type AccountType string

const (
	Assets      AccountType = "Assets"      // + | -
	Capitals    AccountType = "Capitals"    // - | +
	Expenses    AccountType = "Expenses"    // + | -
	Income      AccountType = "Income"      // - | +
	Liabilities AccountType = "Liabilities" // - | +
)

type Accounts struct {
	ID   int64 `gorm:"primaryKey;autoIncrement;"`
	Name string
	Type AccountType
	// AccountNumber string
}

func (Accounts) TableName() string {
	return "list_of_accounts"
}
