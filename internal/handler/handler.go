package handler

import (
	"super-indo-be/internal/config"
	"super-indo-be/internal/service"
)

type Handler struct {
	UserHandler IUserHandler
}

func InitiateHandler(cfg *config.Config, services *service.Service) *Handler {
	return &Handler{
		UserHandler: NewUserHandler(cfg, services.UserService),
	}
}
