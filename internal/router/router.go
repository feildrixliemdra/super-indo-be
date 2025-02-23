package router

import (
	"super-indo-be/internal/config"
	"super-indo-be/internal/handler"
	"super-indo-be/internal/middleware"

	_ "super-indo-be/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @title           Super Indo API
// @version         1.0
// @description     This is a Super Indo API.

// @contact.name   Feildrix Liemdra
// @contact.url    https://github.com/feildrixliemdra
// @contact.email  feildrix@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @schemes http https
func (r *router) Init() {

	r.rtr.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Super Indo API is running",
		})
	})

	// CORS middleware
	r.rtr.Use(middleware.CORSMiddleware)

	// Swagger
	if r.cfg.Swagger.IsEnabled {
		r.rtr.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	v1Group := r.rtr.Group("/v1")

	//Authentication Route
	v1Group.POST("/login", r.handler.AuthHandler.Login)
	v1Group.POST("/register", r.handler.AuthHandler.Register)

	// User Route
	userRouter := v1Group.Group("/users", middleware.JWTAuth(r.cfg.JWT.SecretKey))
	userRouter.GET("/detail", r.handler.UserHandler.GetDetail)

	// Category Route
	categoryRouter := v1Group.Group("/categories", middleware.JWTAuth(r.cfg.JWT.SecretKey))
	categoryRouter.POST("", r.handler.CategoryHandler.Create)
	categoryRouter.GET("", r.handler.CategoryHandler.GetAll)

	// Product Route
	productRouter := v1Group.Group("/products", middleware.JWTAuth(r.cfg.JWT.SecretKey))
	productRouter.POST("", r.handler.ProductHandler.Create)
	productRouter.GET("", r.handler.ProductHandler.GetAll)
	productRouter.GET("/:id", r.handler.ProductHandler.GetByID)
}
