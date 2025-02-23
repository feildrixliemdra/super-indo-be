package model

import "time"

type CartItem struct {
	ID        uint64    `db:"id"`
	CartID    uint64    `db:"cart_id"`
	ProductID uint64    `db:"product_id"`
	Quantity  int       `db:"quantity"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}

type CarItemDetail struct {
	ID           uint64    `db:"id"`
	CartID       uint64    `db:"cart_id"`
	ProductID    uint64    `db:"product_id"`
	Quantity     int       `db:"quantity"`
	ProductName  string    `db:"product_name"`
	ProductPrice int       `db:"product_price"`
	ProductImage string    `db:"product_image"`
	CategoryID   uint64    `db:"category_id"`
	CategoryName string    `db:"category_name"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func (c CartItem) TableName() string {
	return "cart_items"
}
