package models

import (
	"gopkg.in/guregu/null.v4"
)

type UserSessions struct {
	Id        int `gorm:"primaryKey"`
	StartTime null.Time
	EndTime   null.Time
	UserId    int
}
