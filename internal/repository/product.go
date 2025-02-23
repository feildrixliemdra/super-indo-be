package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"super-indo-be/internal/errorcustom"
	"super-indo-be/internal/model"
	"super-indo-be/internal/payload"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type IProductRepository interface {
	Create(ctx context.Context, product model.Product) (uint64, error)
	GetAll(ctx context.Context, p payload.GetProductListRequest) (product []model.Product, totalData int64, err error)
	GetByID(ctx context.Context, id uint64) (*model.Product, error)
}

type product struct {
	DB *sqlx.DB
}

func NewProductRepository(opt Option) IProductRepository {
	return &product{
		DB: opt.DB,
	}
}

func (r *product) Create(ctx context.Context, product model.Product) (id uint64, err error) {

	q := sq.Insert(product.TableName()).
		SetMap(sq.Eq{
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
			"category_id": product.CategoryID,
			"image":       product.Image,
			"stock":       product.Stock,
		}).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := q.ToSql()
	if err != nil {
		return 0, err
	}

	err = r.DB.QueryRowxContext(ctx, query, args...).Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), `violates foreign key constraint "products_category_id_fkey" (SQLSTATE 23503)`) {
			return 0, errorcustom.ErrCategoryNotFound
		}
		return 0, err
	}

	return id, nil
}

func (r *product) GetAll(ctx context.Context, p payload.GetProductListRequest) (product []model.Product, totalData int64, err error) {

	// get product list
	q := sq.Select("id", "name", "description", "price", "category_id", "image", "stock").
		From(model.Product{}.TableName()).
		Where(sq.Eq{"deleted_at": nil}).
		Limit(p.Limit).
		Offset((p.Page - 1) * p.Limit).
		PlaceholderFormat(sq.Dollar)

	query, args, err := q.ToSql()
	if err != nil {
		return nil, 0, err
	}

	err = r.DB.SelectContext(ctx, &product, query, args...)
	if err != nil {
		return nil, 0, err
	}

	// Count total data
	countQuery, countArgs, err := sq.Select("COUNT(id)").
		From(model.Product{}.TableName()).
		Where(sq.Eq{"deleted_at": nil}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, 0, err
	}

	err = r.DB.GetContext(ctx, &totalData, countQuery, countArgs...)
	if err != nil {
		return nil, 0, err
	}

	return product, totalData, nil
}

func (r *product) GetByID(ctx context.Context, id uint64) (*model.Product, error) {

	var product model.Product
	q := sq.Select("id", "name", "description", "price", "category_id", "image", "stock").
		From(model.Product{}.TableName()).
		Where(sq.Eq{"id": id}).
		Where(sq.Eq{"deleted_at": nil}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	err = r.DB.GetContext(ctx, &product, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}
