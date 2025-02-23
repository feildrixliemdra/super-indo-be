package router

import (
	"super-indo-be/internal/config"
	"super-indo-be/internal/handler"
	"super-indo-be/internal/middleware"

	"github.com/gin-gonic/gin"
)

type router struct {
	rtr     *gin.Engine
	handler *handler.Handler
	cfg     *config.Config
}

func NewRouter(rtr *gin.Engine, cfg *config.Config, handler *handler.Handler) Router {
	return &router{
		rtr,
		handler,
		cfg,
	}
}

type Router interface {
	Init()
}

func (r *router) Init() {

	r.rtr.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	v1Group := r.rtr.Group("/v1")
	v1Group.POST("/login", r.handler.AuthHandler.Login)
	v1Group.POST("/register", r.handler.AuthHandler.Register)

	userRouter := v1Group.Group("/users", middleware.JWTAuth(r.cfg.JWT.SecretKey))
	userRouter.GET("/detail", r.handler.UserHandler.GetDetail)

	categoryRouter := v1Group.Group("/categories", middleware.JWTAuth(r.cfg.JWT.SecretKey))
	categoryRouter.POST("", r.handler.CategoryHandler.Create)
	categoryRouter.GET("", r.handler.CategoryHandler.GetAll)

}
