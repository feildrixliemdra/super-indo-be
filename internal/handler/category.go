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

// Create implements ICategoryHandler.
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

// GetAll implements ICategoryHandler.
func (h *category) GetAll(c *gin.Context) {
	res, err := h.service.GetAll(c)
	if err != nil {
		util.ErrInternalResponse(c, err)
		return
	}

	util.GeneralSuccessResponse(c, "success get all category", res)
}
