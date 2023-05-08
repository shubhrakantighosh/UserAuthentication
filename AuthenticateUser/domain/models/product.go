package models

type Product struct {
	ID       int32    `db:"id"`
	Price    float64  `db:"price"`
	Images   []string `db:"images"`
	Category string   `db:"category"`
}
