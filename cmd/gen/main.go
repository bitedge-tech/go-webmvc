package main

import (
	config "go-webmvc/config"
	"go-webmvc/internal/repository/model"
	"go-webmvc/pkg/db"
	"go-webmvc/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/gen"
)

// Querier Dynamic SQL
type Querier interface {
	// FilterWithNameAndRole SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {

	// 加载配置
	config.LoadConfig()

	// 初始化日志
	if err := logger.Init(config.AppConfig.Log); err != nil {
		logger.Error("初始化日志失败:", zap.Error(err))
	}
	defer logger.Sync()

	//初始化数据库连接
	db.InitMySQL()
	// 确保程序退出时关闭数据库连接
	defer db.CloseMySQL()

	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/repository/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(db.DB) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(
		model.SysUser{},
		model.SysRole{},
		model.SysMenu{},
	)

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	g.ApplyInterface(func(Querier) {},
		model.SysUser{},
		model.SysMenu{},
	)

	// Generate the code
	g.Execute()
}
