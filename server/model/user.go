package model

import "time"

type Users struct {
	UserID       uint      `json:"user_id" gorm:"primaryKey"`
	UserName     string    `json:"username" gorm:"not null;colum:user_name;size:255"`
	Password     string    `json:"password" gorm:"not null;colum:password;size:255"`
	Email        string    `json:"email" gorm:"not null;colum:email;size:255"`
	Role         string    `json:"role" gorm:"not null;colum:role;size:255"`
	CreatedBy    string    `json:"created_by" gorm:"not null;colum:created_by;size:255"`
	CreatedDate  time.Time `json:"created_date" gorm:"not null;colum:created_date"`
	ModifiedBy   string    `json:"modified_by" gorm:"not null;colum:modified_by;size:255"`
	ModifiedDate time.Time `json:"modified_date" gorm:"not null;colum:modified_date"`
}
