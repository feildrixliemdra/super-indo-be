package model

import "time"

type Cart struct {
	ID        uint64    `db:"id"`
	UserID    uint64    `db:"user_id"`
	ProductID uint64    `db:"product_id"`
	Quantity  uint64    `db:"quantity"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}

func (c Cart) TableName() string {
	return "carts"
}
