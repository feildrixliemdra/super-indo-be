package handler

import (
	"errors"
	"super-indo-be/internal/config"
	"super-indo-be/internal/errorcustom"
	"super-indo-be/internal/payload"
	"super-indo-be/internal/service"
	"super-indo-be/internal/util"

	"github.com/gin-gonic/gin"
)

type ICategoryHandler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
}

type category struct {
	cfg     *config.Config
	service service.ICategoryService
}

func NewCategoryHandler(cfg *config.Config, service service.ICategoryService) ICategoryHandler {
	return &category{cfg: cfg, service: service}
}

// Create godoc
// @Summary Create a new category
// @Description Create a new category with the given name
// @Tags Category
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT token" default "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG5kb2VAZ21haWwuY29tIiwiZXhwIjoxNzQwOTI1MjEwLCJ1c2VyX2lkIjo0fQ.pqtisYNaYFZYJF49SM2jiBBAsAKFY6ovCJeokcUZcjk"
// @Param request body payload.CreateCategoryRequest true "Create Category Request"
// @Success 200 {object} payload.Response{data=payload.CreateCategoryResponse} "Success create category"
// @Failure 400 {object} payload.Response "Bad request"
// @Failure 500 {object} payload.Response "Internal server error"
// @Router /v1/categories [post]
func (h *category) Create(c *gin.Context) {
	var p payload.CreateCategoryRequest
	if err := c.ShouldBindJSON(&p); err != nil {
		util.ErrBindResponse(c, err)
		return
	}

	res, err := h.service.Create(c, p)
	if err != nil {
		if errors.Is(err, errorcustom.ErrCategoryAlreadyExists) {
			util.ErrBadRequestResponse(c, err.Error())
			return
		}

		util.ErrInternalResponse(c, err)
		return
	}

	util.GeneralSuccessResponse(c, "success create category", res)
}

// GetAll godoc
// @Summary Get all categories
// @Description Get all categories with pagination
// @Tags Category
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT token" default "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG5kb2VAZ21haWwuY29tIiwiZXhwIjoxNzQwOTI1MjEwLCJ1c2VyX2lkIjo0fQ.pqtisYNaYFZYJF49SM2jiBBAsAKFY6ovCJeokcUZcjk"
// @Success 200 {object} payload.Response{data=[]payload.GetCategoryListResponse} "Success get all category"
// @Failure 400 {object} payload.Response "Bad request"
// @Failure 500 {object} payload.Response "Internal server error"
// @Router /v1/categories [get]
func (h *category) GetAll(c *gin.Context) {
	res, err := h.service.GetAll(c)
	if err != nil {
		util.ErrInternalResponse(c, err)
		return
	}

	util.GeneralSuccessResponse(c, "success get all category", res)
}
