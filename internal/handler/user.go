package handler

import (
	"errors"
	"super-indo-be/internal/config"
	"super-indo-be/internal/errorcustom"
	"super-indo-be/internal/service"
	"super-indo-be/internal/util"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type IUserHandler interface {
	GetDetail(c *gin.Context)
}

type user struct {
	cfg         *config.Config
	UserService service.IUserService
}

func NewUserHandler(cfg *config.Config, userService service.IUserService) IUserHandler {
	return &user{
		cfg:         cfg,
		UserService: userService,
	}
}

// GetDetail godoc
// @Summary Get user detail
// @Description Get authenticated user's detail information
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "Bearer JWT token" default "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG5kb2VAZ21haWwuY29tIiwiZXhwIjoxNzQwOTI1MjEwLCJ1c2VyX2lkIjo0fQ.pqtisYNaYFZYJF49SM2jiBBAsAKFY6ovCJeokcUZcjk"
// @Success 200 {object} payload.Response{data=payload.GetUserDetailData} "Success get user detail"
// @Failure 400 {object} payload.Response "Bad request"
// @Failure 500 {object} payload.Response "Internal server error"
// @Router /v1/users/detail [get]
func (h *user) GetDetail(c *gin.Context) {
	token := util.ExtractToken(c)
	user, err := util.ParseJWT(token, h.cfg.JWT.SecretKey)
	if err != nil {
		util.ErrBadRequestResponse(c, err.Error())
		return
	}

	result, err := h.UserService.GetByID(c, user.UserID)
	if err != nil {
		if errors.Is(err, errorcustom.ErrUserNotFound) {
			log.Warnf("user id not found %v", user.UserID)
			util.ErrBadRequestResponse(c, err.Error())

			return
		}

		util.ErrInternalResponse(c, err)

		return
	}

	util.GeneralSuccessResponse(c, "success get user detail", result)

}
