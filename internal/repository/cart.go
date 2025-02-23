package repository

import (
	"context"
	"database/sql"
	"errors"
	"super-indo-be/internal/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type ICartRepository interface {
	CreateCart(ctx context.Context, cart model.Cart) (uint64, error)
	GetCartByUserID(ctx context.Context, userID uint64) (*model.Cart, error)
	CreateOrUpdateCartItem(ctx context.Context, items []model.CartItem) error
	GetAllCartItems(ctx context.Context, cartID uint64) ([]model.CarItemDetail, error)
}

type cartRepository struct {
	db *sqlx.DB
}

func NewCartRepository(opt Option) ICartRepository {
	return &cartRepository{
		db: opt.DB,
	}
}

func (r *cartRepository) CreateCart(ctx context.Context, cart model.Cart) (id uint64, err error) {
	query, args, err := sq.Insert(model.Cart{}.TableName()).
		SetMap(sq.Eq{
			"user_id": cart.UserID,
		}).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return 0, err
	}

	err = r.db.QueryRowxContext(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *cartRepository) GetCartByUserID(ctx context.Context, userID uint64) (*model.Cart, error) {
	var cart model.Cart

	query, args, err := sq.Select("id", "user_id").
		From(model.Cart{}.TableName()).
		Where(sq.Eq{"user_id": userID, "deleted_at": nil}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.GetContext(ctx, &cart, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &cart, err
}

func (r *cartRepository) GetAllCartItems(ctx context.Context, cartID uint64) (cartItems []model.CarItemDetail, err error) {
	query, args, err := sq.Select(
		"ci.id", "ci.cart_id", "ci.product_id", "ci.quantity",
		"p.name as product_name", "p.price as product_price", "p.image as product_image",
		"c.name as category_name", "c.id as category_id",
	).
		From(model.CartItem{}.TableName() + " ci").
		Join("products p ON p.id = ci.product_id").
		Join("categories c ON c.id = p.category_id").
		Where(sq.Eq{"ci.cart_id": cartID, "ci.deleted_at": nil}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = r.db.SelectContext(ctx, &cartItems, query, args...)
	return cartItems, err
}

func (r *cartRepository) CreateOrUpdateCartItem(ctx context.Context, items []model.CartItem) error {
	// Start transaction
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	// Defer rollback in case of error - will be ignored if transaction is committed
	defer tx.Rollback()

	for _, item := range items {
		// Check if cart item exists
		query, args, err := sq.Select("id").
			From(model.CartItem{}.TableName()).
			Where(sq.Eq{
				"cart_id":    item.CartID,
				"product_id": item.ProductID,
				"deleted_at": nil,
			}).
			PlaceholderFormat(sq.Dollar).
			ToSql()

		if err != nil {
			return err
		}

		var existingID uint64
		err = tx.QueryRowxContext(ctx, query, args...).Scan(&existingID)

		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return err
		}

		if existingID > 0 {
			// Update existing cart item
			query, args, err = sq.Update(model.CartItem{}.TableName()).
				SetMap(sq.Eq{
					"quantity":   item.Quantity,
					"updated_at": sq.Expr("NOW()"),
				}).
				Where(sq.Eq{"id": existingID}).
				PlaceholderFormat(sq.Dollar).
				ToSql()
		} else {
			// Insert new cart item
			query, args, err = sq.Insert(model.CartItem{}.TableName()).
				SetMap(sq.Eq{
					"cart_id":    item.CartID,
					"product_id": item.ProductID,
					"quantity":   item.Quantity,
				}).
				PlaceholderFormat(sq.Dollar).
				ToSql()
		}

		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, query, args...)
		if err != nil {
			return err
		}
	}

	// Commit transaction
	return tx.Commit()
}
