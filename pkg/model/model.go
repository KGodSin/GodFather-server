package model

import (
	"time"
)

type User struct {
	ID        uint
	UserID    string `gorm:"primaryKey"`
	Password  string
	Role      string `gorm:"type:ENUM('admin', 'user')"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID        uint
	Name      string `gorm:"primaryKey"`
	Price     int
	Count     int
	CreatedAt time.Time
}

type status string

type ApprovalInformation struct {
	ID uint
	// Status       string `gorm:"type:ENUM('confirm', 'wait')"`
	Status       status `json:"status" sql:"type:ENUM('confirm', 'wait')"`
	ApprovalDate time.Time
	Count        int
	UserID       string `gorm:"foreignKey"`
	ProductID    string `gorm:"foreignKey"`
	CreatedAt    time.Time
}
