package global

import (
	"api-platform/config"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger *zap.Logger
	Config *config.Config
	Redis  *redis.Client
	DB     *gorm.DB
)
