package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	config "go-webmvc/config"
	"go-webmvc/pkg/logger"
	"go.uber.org/zap"
)

var (
	Client *redis.Client
	Ctx    = context.Background()
)

func InitRedis() {
	cfg := config.AppConfig.Redis
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       0,
	})

	if err := Client.Ping(Ctx).Err(); err != nil {
		logger.Log.Error("Failed to connect to Redis", zap.Error(err))
	} else {
		logger.Log.Info("Redis connected successfully")
	}
}

func CloseRedis() {
	if err := Client.Close(); err != nil {
		logger.Log.Error("Failed to close Redis connection", zap.Error(err))
	} else {
		logger.Log.Info("Redis connection closed.")
	}
}
