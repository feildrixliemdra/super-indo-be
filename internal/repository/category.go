package repository

import (
	"context"
	"super-indo-be/internal/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type ICategoryRepository interface {
	Create(ctx context.Context, category model.Category) (id uint64, err error)
	GetAll(ctx context.Context) ([]model.Category, error)
	GetBy(ctx context.Context, category model.Category) (*model.Category, error)
}

type category struct {
	DB *sqlx.DB
}

// GetBy implements ICategoryRepository.
func (c *category) GetBy(ctx context.Context, category model.Category) (*model.Category, error) {
	var result model.Category

	q := sq.Select("id", "name", "code", "description", "created_at", "updated_at").
		From(model.Category{}.TableName()).
		Where(sq.Eq{"deleted_at": nil}).
		PlaceholderFormat(sq.Dollar)

	if category.ID > 0 {
		q = q.Where(sq.Eq{"id": category.ID})
	}

	if category.Code != "" {
		q = q.Where(sq.Eq{"code": category.Code})
	}

	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	err = c.DB.GetContext(ctx, &result, query, args...)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func NewCategoryRepository(opt Option) ICategoryRepository {
	return &category{
		DB: opt.DB,
	}
}

// Create implements ICategoryRepository.
func (c *category) Create(ctx context.Context, category model.Category) (id uint64, err error) {
	query, args, err := sq.Insert(model.Category{}.TableName()).
		SetMap(
			sq.Eq{
				"name":        category.Name,
				"code":        category.Code,
				"description": category.Description,
			},
		).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, err
	}

	result, err := c.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	id = uint64(lastInsertID)

	return id, nil
}

// GetAll implements ICategoryRepository.
func (c *category) GetAll(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category

	query, args, err := sq.Select("id", "name", "code", "description", "created_at", "updated_at").
		From(model.Category{}.TableName()).
		Where(sq.Eq{"deleted_at": nil}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = c.DB.SelectContext(ctx, &categories, query, args...)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
