package models

import "gorm.io/gorm"

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
