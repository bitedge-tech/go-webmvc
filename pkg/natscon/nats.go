package natscon

import (
	"fmt"
	config "go-webmvc/config"
	"go-webmvc/pkg/logger"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

var Conn *nats.Conn

func InitNATS() {
	cfg := config.AppConfig.Nats
	url := fmt.Sprintf("nats://%s:%s", cfg.Host, cfg.Port)
	var opts []nats.Option
	if cfg.User != "" {
		opts = append(opts, nats.UserInfo(cfg.User, cfg.Password))
	}
	var err error
	Conn, err = nats.Connect(url, opts...)
	if err != nil {
		logger.Log.Error("Failed to connect to NATS", zap.Error(err))
	} else {
		logger.Log.Info("NATS connected successfully")
	}
}

func CloseNATS() {
	if Conn != nil {
		Conn.Close()
		logger.Log.Info("NATS connection closed.")
	}
}
