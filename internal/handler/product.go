package handler

import (
	"strconv"
	"super-indo-be/internal/config"
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

// Create implements IProductHandler.
func (h *product) Create(c *gin.Context) {
	var p payload.CreateProductRequest
	if err := c.ShouldBindJSON(&p); err != nil {
		util.ErrBindResponse(c, err)
		return
	}

	res, err := h.service.Create(c, p)
	if err != nil {
		util.ErrInternalResponse(c, err)
		return
	}

	util.GeneralSuccessResponse(c, "success create product", res)
}

// GetAll implements IProductHandler.
func (h *product) GetAll(c *gin.Context) {
	var p payload.GetProductListRequest
	if err := c.ShouldBindQuery(&p); err != nil {
		util.ErrBindResponse(c, err)
		return
	}

	res, totalData, err := h.service.GetAll(c, p)
	if err != nil {
		util.ErrInternalResponse(c, err)
		return
	}

	totalPage := (totalData + int64(p.Limit) - 1) / int64(p.Limit)
	util.GeneralSuccessListResponse(c, "success get all product", res, totalData, totalPage, int64(p.Page))
}

// GetByID implements IProductHandler.
func (h *product) GetByID(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		util.ErrBadRequestResponse(c, "Invalid product ID")
		return
	}

	res, err := h.service.GetByID(c, idUint)
	if err != nil {
		util.ErrInternalResponse(c, err)
		return
	}

	util.GeneralSuccessResponse(c, "success get product by id", res)
}
