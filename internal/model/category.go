package model

import "time"

type Category struct {
	ID          uint64    `db:"id"`
	Name        string    `db:"name"`
	Code        string    `db:"code"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	DeletedAt   time.Time `db:"deleted_at"`
}

func (c Category) TableName() string {
	return "categories"
}
