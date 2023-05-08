package models

import (
	"net/url"
	"time"
)

type ProductImages struct {
	Id        int
	Url       url.URL
	ProductId int
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
	DeletedAt time.Time
	DeletedBy string
}
