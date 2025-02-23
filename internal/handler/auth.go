package handler

import (
	"super-indo-be/internal/config"
	"super-indo-be/internal/errorcustom"
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

// Login godoc
// @Summary User login
// @Description Authenticate user and generate JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body payload.LoginRequest true "Login credentials"
// @Success 200 {object} payload.Response{data=payload.LoginResponse} "Login success with token"
// @Failure 400 {object} payload.Response "Bad request"
// @Failure 500 {object} payload.Response "Internal server error"
// @Router /v1/login [post]
func (h *auth) Login(c *gin.Context) {
	var p payload.LoginRequest
	if err := c.ShouldBindJSON(&p); err != nil {
		util.ErrBindResponse(c, err)
		return
	}

	user, err := h.userService.GetByEmail(c.Request.Context(), p.Email)
	if err != nil {
		util.ErrBadRequestResponse(c, err.Error())
		return
	}

	token, err := util.GenerateJWT(util.JWTUser{
		UserID: user.ID,
		Email:  user.Email,
	}, h.cfg.JWT.SecretKey)

	if err != nil {
		util.ErrInternalResponse(c, err)
		return
	}

	util.GeneralSuccessResponse(c, "login success", payload.LoginResponse{
		Token: token,
	})
}

// Register godoc
// @Summary User registration
// @Description Register new user and generate JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body payload.RegisterRequest true "Registration details"
// @Success 200 {object} payload.Response{data=payload.RegisterResponse} "Registration success with token"
// @Failure 400 {object} payload.Response "Bad request or email already registered"
// @Failure 500 {object} payload.Response "Internal server error"
// @Router /v1/register [post]
func (h *auth) Register(c *gin.Context) {
	var p payload.RegisterRequest
	if err := c.ShouldBindJSON(&p); err != nil {
		util.ErrBindResponse(c, err)
		return
	}

	user, err := h.userService.Create(c.Request.Context(), payload.CreateUserRequest{
		Name:     p.Name,
		Email:    p.Email,
		Password: p.Password,
	})
	if err != nil {
		if err == errorcustom.ErrEmailAlreadyRegistered {
			util.ErrBadRequestResponse(c, err.Error())
			return
		}

		util.ErrInternalResponse(c, err)
		return
	}

	token, err := util.GenerateJWT(util.JWTUser{
		UserID: user.ID,
		Email:  user.Email,
	}, h.cfg.JWT.SecretKey)

	if err != nil {
		util.ErrInternalResponse(c, err)
		return
	}

	util.GeneralSuccessResponse(c, "register success", payload.RegisterResponse{
		Token: token,
	})
}
