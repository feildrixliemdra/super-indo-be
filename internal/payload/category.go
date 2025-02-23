package payload

import (
	"super-indo-be/internal/model"
	"time"
)

type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Code        string `json:"code" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (c *CreateCategoryRequest) ToModel() model.Category {
	return model.Category{
		Name:        c.Name,
		Code:        c.Code,
		Description: c.Description,
	}
}

type CreateCategoryResponse struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type GetCategoryListResponse struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
