package models

type Account struct {
	AccountID      int     `json:"account_id" gorm:"primaryKey"`
	InitialBalance float64 `json:"initial_balance" gorm:"type:decimal(20,5)"`
}
