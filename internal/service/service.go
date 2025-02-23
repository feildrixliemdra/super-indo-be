package service

import (
	"super-indo-be/internal/config"
	"super-indo-be/internal/repository"
)

type Service struct {
	UserService     IUserService
	CategoryService ICategoryService
	ProductService  IProductService
	CartService     ICartService
}

type Option struct {
	Repository *repository.Repository
}

func InitiateService(cfg *config.Config, repository *repository.Repository) *Service {
	return &Service{
		UserService:     NewUserService(repository.UserRepository),
		CategoryService: NewCategoryService(repository.CategoryRepository),
		ProductService:  NewProductService(repository.ProductRepository),
		CartService:     NewCartService(repository.CartRepository),
	}
}
