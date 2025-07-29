package models

type Transaction struct {
	ID                   uint    `gorm:"primaryKey"`
	SourceAccountID      int     `json:"source_account_id"`
	DestinationAccountID int     `json:"destination_account_id"`
	Amount               float64 `json:"amount" gorm:"type:decimal(20,5)"`
}
