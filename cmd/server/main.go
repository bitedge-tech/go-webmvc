package main

import (
	"go-webmvc/config"
	"go-webmvc/internal/repository/query"
	"go-webmvc/internal/router"
	"go-webmvc/internal/service"
	"go-webmvc/pkg/db"
	"go-webmvc/pkg/logger"
	"go-webmvc/pkg/redis"
	"log"
)

func main() {

	// 加载配置
	config.LoadConfig()

	// 初始化日志
	if err := logger.Init(config.AppConfig.Log); err != nil {
		log.Fatal("初始化日志失败:", err)
	}
	defer logger.Sync()

	//初始化数据库连接
	db.InitMySQL()
	// 确保程序退出时关闭数据库连接
	defer db.CloseMySQL()

	// 进行数据库迁移:根据 internal/dal/model 目录下的模型自动创建和更新数据库表结构
	if err := db.MigrateDB(); err != nil {
		logger.Error("数据库迁移失败: " + err.Error())
		return
	}

	//初始化Redis连接,如不需要可注释掉
	redis.InitRedis()
	defer redis.CloseRedis()

	//初始化gorm-gen生成的query
	query.SetDefault(db.DB)

	// 初始化NATS连接
	//natCon.InitNATS()
	//defer natCon.CloseNATS()

	//初始化Service
	service.InitService()

	// 设置并启动路由
	r := router.SetupRouter()

	// 启动服务器
	port := config.AppConfig.App.Port
	logger.Info("服务器启动，监听端口 " + port)
	err := r.Run(":" + port)
	if err != nil {
		logger.Error("服务器启动失败: " + err.Error())
		return
	}
}
