package model

import "time"

type Transactions struct {
	TransactionID   int       `json:"transaction_id" gorm:"primaryKey"`
	UserID          int       `json:"user_id" gorm:"not null;column:user_id"`
	BookID          int       `json:"book_id" gorm:"not null;column:book_id"`
	TransactionDate time.Time `json:"transaction_date" gorm:"not null;colum:transaction_date"`
	ReturnDate      time.Time `json:"return_date" gorm:"colum:return_date"`
	Status          string    `json:"status" gorm:"not null;colum:status;size:50"`
	CreatedBy       string    `json:"created_by" gorm:"not null;colum:created_by;size:255"`
	CreatedDate     time.Time `json:"created_date" gorm:"not null;colum:created_date"`
	ModifiedBy      string    `json:"modified_by" gorm:"not null;colum:modified_by;size:255"`
	ModifiedDate    time.Time `json:"modified_date" gorm:"not null;colum:modified_date"`
}
