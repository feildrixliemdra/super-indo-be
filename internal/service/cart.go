package service

import (
	"context"
	"super-indo-be/internal/errorcustom"
	"super-indo-be/internal/model"
	"super-indo-be/internal/payload"
	"super-indo-be/internal/repository"

	log "github.com/sirupsen/logrus"
)

type ICartService interface {
	CreateOrUpdateCartItem(ctx context.Context, userID uint64, request []payload.CreateCartItemRequest) error
	GetAll(ctx context.Context, userID uint64) (payload.GetAllCartItemResponse, error)
}

type cartService struct {
	repo repository.ICartRepository
}

func NewCartService(repo repository.ICartRepository) ICartService {
	return &cartService{
		repo: repo,
	}
}

func (s *cartService) CreateOrUpdateCartItem(ctx context.Context, userID uint64, request []payload.CreateCartItemRequest) error {
	cart, err := s.repo.GetCartByUserID(ctx, userID)
	if err != nil {
		log.Errorf("error get cart by user id: %v", err)
		return err
	}

	// If cart doesn't exist, create new cart first
	if cart == nil {
		cartID, err := s.repo.CreateCart(ctx, model.Cart{
			UserID: userID,
		})
		if err != nil {
			log.Errorf("error create cart: %v", err)
			return err
		}
		cart = &model.Cart{
			ID:     cartID,
			UserID: userID,
		}
	}

	items := make([]model.CartItem, len(request))
	for i, item := range request {
		items[i] = model.CartItem{
			CartID:    cart.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
	}

	return s.repo.CreateOrUpdateCartItem(ctx, items)
}

func (s *cartService) GetAll(ctx context.Context, userID uint64) (res payload.GetAllCartItemResponse, err error) {
	cart, err := s.repo.GetCartByUserID(ctx, userID)
	if err != nil {
		log.Errorf("error get cart by user id: %v", err)
		return res, err
	}

	if cart == nil {
		return res, errorcustom.ErrCartNotFound
	}

	cartItems, err := s.repo.GetAllCartItems(ctx, cart.ID)
	if err != nil {
		log.Errorf("error get all cart items: %v", err)
		return res, err
	}

	res.CartID = cart.ID
	for _, v := range cartItems {
		res.Items = append(res.Items, payload.Item{
			ID:           v.ID,
			ProductID:    v.ProductID,
			Quantity:     v.Quantity,
			ProductName:  v.ProductName,
			ProductPrice: v.ProductPrice,
			ProductImage: v.ProductImage,
			CategoryName: v.CategoryName,
			CategoryID:   v.CategoryID,
		})
	}

	return res, nil
}
