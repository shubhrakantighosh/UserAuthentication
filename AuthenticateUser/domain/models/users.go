package models

import (
	"gopkg.in/guregu/null.v4"
)

type Users struct {
	Id           int    `gorm:"primaryKey"`
	Username     string `gorm:"unique"`
	Password     string
	Mail         string `gorm:"uniqueIndex"`
	MobileNumber string `gorm:"uniqueIndex"`
	CreatedAt    null.Time
	CreatedBy    null.String
	UpdatedAt    null.Time
	UpdatedBy    null.String
	DeletedAt    null.Time
	DeletedBy    null.String
}
