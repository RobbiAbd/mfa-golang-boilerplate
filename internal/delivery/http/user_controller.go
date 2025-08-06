package http

import (
	"github.com/gofiber/fiber"
	"github.com/robbiabd/auth-jwt-golang-boilerplate/internal/usecase"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	UseCase *usecase.UserUseCase
}

func NewUserController(usecase *usecase.UserUseCase, logger *logrus.Logger) *UserController {
	return &UserController{
		Log:     logger,
		UseCase: usecase,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	return nil
}
