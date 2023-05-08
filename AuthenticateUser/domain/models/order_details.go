package models

import "gorm.io/gorm"

type OrderDetails struct {
	gorm.Model
	ProductPrice map[int][]float64 `json:"product_price"`
	UserId       int               `json:"user_id"`
}
