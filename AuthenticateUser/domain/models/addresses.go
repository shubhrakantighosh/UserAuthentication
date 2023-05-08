package models

import "time"

type Addresses struct {
	Id        int
	PinCode   string
	City      string
	State     string
	UserId    int
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
	DeletedAt time.Time
	DeletedBy string
}
