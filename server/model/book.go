package model

import "time"

type Books struct {
	BookID          int       `json:"book_id" gorm:"primaryKey"`
	Title           string    `json:"title" gorm:"not null;column:title;size:255"`
	Author          string    `json:"author" gorm:"not null;column:author;size:100"`
	PublishYear     int       `json:"published_year" gorm:"not null;column:published_year"`
	ISBN            string    `json:"isbn" gorm:"not null;column:isbn;size:20"`
	TotalCopies     int       `json:"total_copies" gorm:"not null;column:total_copies"`
	AvailableCopies int       `json:"available_copies" gorm:"not null;column:available_copies"`
	CategoryID      int       `json:"category_id" gorm:"not null;column:category_id"`
	CreatedBy       string    `json:"created_by" gorm:"not null;colum:created_by;size:255"`
	CreatedDate     time.Time `json:"created_date" gorm:"not null;colum:created_date"`
	ModifiedBy      string    `json:"modified_by" gorm:"not null;colum:modified_by;size:255"`
	ModifiedDate    time.Time `json:"modified_date" gorm:"not null;colum:modified_date"`
}

type ListBooks struct {
	BookID          int    `json:"book_id" gorm:"primaryKey"`
	Title           string `json:"title" gorm:"not null;column:title;size:255"`
	Author          string `json:"author" gorm:"not null;column:author;size:100"`
	PublishYear     int    `json:"published_year" gorm:"not null;column:published_year"`
	ISBN            string `json:"isbn" gorm:"not null;column:isbn;size:20"`
	TotalCopies     int    `json:"total_copies" gorm:"not null;column:total_copies"`
	AvailableCopies int    `json:"available_copies" gorm:"not null;column:available_copies"`
	CategoryName    string `json:"category_name" gorm:"not null;column:category_name"`
}
