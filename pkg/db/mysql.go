package db

import (
	"fmt"
	"go-webmvc/config"
	"go-webmvc/internal/repository/model"
	"go-webmvc/pkg/logger"
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitMySQL() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConfig.Database.User,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.Host,
		config.AppConfig.Database.Port,
		config.AppConfig.Database.Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{

		Logger:                                   gormLogger.Default.LogMode(gormLogger.Info), // 设置日志级别
		DisableForeignKeyConstraintWhenMigrating: true,                                        // 关键配置：禁用迁移时的外键约束

	})
	if err != nil {
		logger.Log.Error("Failed to connect to MySQL", zap.Error(err))
	}

	// 配置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		logger.Error("Failed to get sql.DB", zap.Error(err))
		os.Exit(1)
	}
	sqlDB.SetMaxOpenConns(100)                 // 最大连接数
	sqlDB.SetMaxIdleConns(10)                  // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // 连接最大生命周期

	logger.Log.Info("MySQL connected successfully")
}

func CloseMySQL() {
	sqlDB, err := DB.DB()
	if err != nil {
		logger.Log.Error("Failed to get sql.DB", zap.Error(err))
		return
	}
	if err := sqlDB.Close(); err != nil {
		logger.Log.Error("Failed to close database connection", zap.Error(err))
	} else {
		logger.Log.Info("Database connection closed.")
	}
}

func MigrateDB() error {
	// 自动迁移所有模型
	return DB.AutoMigrate(
		&model.SysUser{},
		&model.SysRole{},
		&model.SysMenu{},
	)
}
