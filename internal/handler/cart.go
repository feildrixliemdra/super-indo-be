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

type ICartHandler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
}

type cart struct {
	cfg     *config.Config
	service service.ICartService
}

func NewCartHandler(cfg *config.Config, service service.ICartService) ICartHandler {
	return &cart{cfg: cfg, service: service}
}

// Create godoc
// @Summary Create a new cart item
// @Description Create a new cart item
// @Tags Cart
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT token" default "Bearer <token>"
// @Param request body payload.CreateCartItemRequest true "Create Cart Item Request"
// @Success 200 {object} payload.Response "Success create cart item"
// @Failure 400 {object} payload.Response "Bad request"
// @Failure 500 {object} payload.Response "Internal server error"
// @Router /v1/carts [post]
func (h *cart) Create(c *gin.Context) {
	token := util.ExtractToken(c)
	user, err := util.ParseJWT(token, h.cfg.JWT.SecretKey)
	if err != nil {
		util.ErrBadRequestResponse(c, err.Error())
		return
	}

	var p payload.CreateCartItemRequest
	if err := c.ShouldBindJSON(&p); err != nil {
		util.ErrBindResponse(c, err)
		return
	}

	err = h.service.CreateOrUpdateCartItem(c, user.UserID, []payload.CreateCartItemRequest{p})
	if err != nil {
		util.ErrInternalResponse(c, err)
		return
	}

	util.GeneralSuccessResponse(c, "success create or update cart item", nil)
}

// GetAll godoc
// @Summary Get all cart items
// @Description Get all cart items
// @Tags Cart
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT token" default "Bearer <token>"
// @Success 200 {object} payload.Response{data=[]payload.GetAllCartItemResponse} "Success get all cart items"
// @Failure 400 {object} payload.Response "Bad request"
// @Failure 500 {object} payload.Response "Internal server error"
// @Router /v1/carts [get]
func (h *cart) GetAll(c *gin.Context) {
	token := util.ExtractToken(c)
	jwtUser, err := util.ParseJWT(token, h.cfg.JWT.SecretKey)
	if err != nil {
		util.ErrBadRequestResponse(c, err.Error())
		return
	}

	res, err := h.service.GetAll(c, jwtUser.UserID)
	if err != nil {
		if errors.Is(err, errorcustom.ErrCartNotFound) {
			util.ErrBadRequestResponse(c, err.Error())
			return
		}

		util.ErrInternalResponse(c, err)
		return
	}

	util.GeneralSuccessResponse(c, "success get all cart items", res)
}
