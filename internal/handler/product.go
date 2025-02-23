package handler

import (
	"errors"
	"strconv"
	"super-indo-be/internal/config"
	"super-indo-be/internal/errorcustom"
	"super-indo-be/internal/payload"
	"super-indo-be/internal/service"
	"super-indo-be/internal/util"

	"github.com/gin-gonic/gin"
)

type IProductHandler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
}

type product struct {
	cfg     *config.Config
	service service.IProductService
}

func NewProductHandler(cfg *config.Config, service service.IProductService) IProductHandler {
	return &product{cfg: cfg, service: service}
}

// Create godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags Product
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT token" default "Bearer <token>"
// @Param request body payload.CreateProductRequest true "Create Product Request"
// @Success 200 {object} payload.Response{data=payload.CreateProductResponse} "Success create product"
// @Failure 400 {object} payload.Response "Bad request"
// @Failure 500 {object} payload.Response "Internal server error"
// @Router /v1/products [post]
func (h *product) Create(c *gin.Context) {
	var p payload.CreateProductRequest
	if err := c.ShouldBindJSON(&p); err != nil {
		util.ErrBindResponse(c, err)
		return
	}

	res, err := h.service.Create(c, p)
	if err != nil {
		if errors.Is(err, errorcustom.ErrCategoryNotFound) {
			util.ErrBadRequestResponse(c, "Category not found")
			return
		}

		util.ErrInternalResponse(c, err)
		return
	}

	util.GeneralSuccessResponse(c, "success create product", res)
}

// GetAll godoc
// @Summary Get all products
// @Description Get all products with pagination
// @Tags Product
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT token" default "Bearer <token>"
// @Param page query int false "Page" default(1)
// @Param limit query int false "Limit" default(10)
// @Success 200 {object} payload.Response{data=[]payload.GetProductListResponse} "Success get all product"
// @Failure 400 {object} payload.Response "Bad request"
// @Failure 500 {object} payload.Response "Internal server error"
// @Router /v1/products [get]
func (h *product) GetAll(c *gin.Context) {
	var p payload.GetProductListRequest
	if err := c.ShouldBindQuery(&p); err != nil {
		util.ErrBindResponse(c, err)
		return
	}

	if p.Page < 1 {
		p.Page = 1
	}

	if p.Limit < 1 {
		p.Limit = 10
	}

	res, totalData, err := h.service.GetAll(c, p)
	if err != nil {
		util.ErrInternalResponse(c, err)
		return
	}

	totalPage := (totalData + int64(p.Limit) - 1) / int64(p.Limit)
	util.GeneralSuccessListResponse(c, "success get all product", res, totalData, totalPage, int64(p.Page))
}

// GetByID godoc
// @Summary Get a product by ID
// @Description Get a product by ID
// @Tags Product
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT token" default "Bearer <token>"
// @Param id path string true "Product ID"
// @Success 200 {object} payload.Response{data=payload.GetProductDetailResponse} "Success get product by id"
// @Failure 400 {object} payload.Response "Bad request"
// @Failure 500 {object} payload.Response "Internal server error"
// @Router /v1/products/{id} [get]
func (h *product) GetByID(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		util.ErrBadRequestResponse(c, "Invalid product ID")
		return
	}

	res, err := h.service.GetByID(c, idUint)
	if err != nil {
		if errors.Is(err, errorcustom.ErrProductNotFound) {
			util.ErrBadRequestResponse(c, "Product not found")
			return
		}

		util.ErrInternalResponse(c, err)
		return
	}

	util.GeneralSuccessResponse(c, "success get product by id", res)
}
