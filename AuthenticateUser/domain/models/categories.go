package models

import "time"

type Categories struct {
	Id           int
	CategoryName string
	ProductIds   []int
	CreatedAt    time.Time
	CreatedBy    string
	UpdatedAt    time.Time
	UpdatedBy    string
	DeletedAt    time.Time
	DeletedBy    string
}
