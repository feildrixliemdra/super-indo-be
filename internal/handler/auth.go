package handler

import (
	"super-indo-be/internal/config"
	"super-indo-be/internal/constant"
	"super-indo-be/internal/payload"
	"super-indo-be/internal/service"
	"super-indo-be/internal/util"

	"github.com/gin-gonic/gin"
)

type IAuthHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type auth struct {
	cfg         *config.Config
	userService service.IUserService
}

func NewAuthHandler(cfg *config.Config, userService service.IUserService) IAuthHandler {
	return &auth{
		cfg:         cfg,
		userService: userService,
	}
}

// Login implements IAuthHandler.
func (a *auth) Login(c *gin.Context) {
	var p payload.LoginRequest
	if err := c.ShouldBindJSON(&p); err != nil {
		util.ErrBindResponse(c, err)
		return
	}

	user, err := a.userService.GetByEmail(c, p.Email)
	if err != nil {
		util.ErrBadRequestResponse(c, err.Error())
		return
	}

	token, err := util.GenerateJWT(util.JWTUser{
		UserID: user.ID,
		Email:  user.Email,
	}, a.cfg.JWT.SecretKey)

	if err != nil {
		util.ErrInternalResponse(c, err)
		return
	}

	util.GeneralSuccessResponse(c, "login success", gin.H{
		"token": token,
	})
}

// Register implements IAuthHandler.
func (a *auth) Register(c *gin.Context) {
	var p payload.RegisterRequest
	if err := c.ShouldBindJSON(&p); err != nil {
		util.ErrBindResponse(c, err)
		return
	}

	user, err := a.userService.Create(c, payload.CreateUserRequest{
		Name:     p.Name,
		Email:    p.Email,
		Password: p.Password,
	})
	if err != nil {
		if err == constant.ErrEmailAlreadyRegistered {
			util.ErrBadRequestResponse(c, err.Error())
			return
		}

		util.ErrInternalResponse(c, err)
		return
	}

	token, err := util.GenerateJWT(util.JWTUser{
		UserID: user.ID,
		Email:  user.Email,
	}, a.cfg.JWT.SecretKey)

	if err != nil {
		util.ErrInternalResponse(c, err)
		return
	}

	util.GeneralSuccessResponse(c, "register success", gin.H{
		"token": token,
	})
}
