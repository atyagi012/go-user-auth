package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	First_name string `json:"first_name" validate:"required  min=2 max=100"`
	Last_name  string `json:"last_name" validate:"required min=2 max=100"`
	Email      string `gorm:"unique" json:"email" validate:"email,required"`
	Phone      string `gorm:"unique" json:"phone" validate:"required"`
	Password   string `json:"password" validate:"required min=6"`
}
