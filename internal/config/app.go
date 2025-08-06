package config

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
	Producer *kafka.Producer
	Redis    *redis.Client
}

func NewBootstrapConfig(config *BootstrapConfig) {
	// setup repositories

	// setup producers

	// setup usecases

	// setup controllers

	// setup middlewares

	// setup routes
}
