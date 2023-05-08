package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderDetails map[int][]int `json:"order_details"`
	Total        float64       `json:"total"`
}
