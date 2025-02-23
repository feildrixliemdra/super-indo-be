package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	UserRepository     IUserRepository
	CategoryRepository ICategoryRepository
	ProductRepository  IProductRepository
	CartRepository     ICartRepository
}

type Option struct {
	DB *sqlx.DB
}

func InitiateRepository(opt Option) *Repository {
	return &Repository{
		UserRepository:     NewUserRepository(opt),
		CategoryRepository: NewCategoryRepository(opt),
		ProductRepository:  NewProductRepository(opt),
		CartRepository:     NewCartRepository(opt),
	}
}
