package model

type Categories struct {
	CategoryID   int    `json:"category_id" gorm:"primaryKey"`
	CategoryName string `json:"category_name" gorm:"not null;colum:category_name;size:255"`
}
