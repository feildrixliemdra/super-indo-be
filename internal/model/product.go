package model

import "time"

type Product struct {
	ID          uint64    `db:"id"`
	Name        string    `db:"name"`
	CategoryID  uint64    `db:"category_id"`
	Price       uint64    `db:"price"`
	Stock       uint64    `db:"stock"`
	Description string    `db:"description"`
	Image       string    `db:"image"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	DeletedAt   time.Time `db:"deleted_at"`
}

func (p Product) TableName() string {
	return "products"
}
