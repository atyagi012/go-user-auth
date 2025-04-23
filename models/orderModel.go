package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Order_id       string `json:"order_id" validate:"required"`
	Payment_status string `json:"payment_status" validate:"required"`
	Payment_mode   string `json:"payment_mode" validate:"required"`
	//Product
}
