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
