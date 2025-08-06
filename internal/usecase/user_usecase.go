package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/robbiabd/auth-jwt-golang-boilerplate/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	UserRepository *repository.UserRepository
}
