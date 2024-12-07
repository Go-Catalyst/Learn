package main

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	// ID          uint       `gorm:"primaryKey"` // gorm handles!
	Username    string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	Email       string `gorm:"unique;not null"` // is the uniqueness needed?
	PhoneNumber string `gorm:"unique;not null"` // is the uniqueness needed?
	Role        string `gorm:"not null"`        // should we store Permissions?
	//CreatedAt   time.Time  // gorm handles!
	//UpdatedAt   time.Time  // gorm handles!
	IsActive    bool       `gorm:"default:true"` // to be able to keep track of former users that otherwise would be deleted
	LastLoginAt *time.Time // necessary?
}
