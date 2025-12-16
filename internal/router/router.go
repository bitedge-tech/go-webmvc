package router

import (
	"go-webmvc/internal/handler/index"
	"go-webmvc/internal/handler/login"
	"go-webmvc/internal/handler/users"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	docs "go-webmvc/docs"

	// swagger UI packages
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	// 设置Gin为调试模式
	app_mode := viper.GetString("app.env")
	if app_mode == "production" {
		gin.SetMode(gin.ReleaseMode)

	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()

	// Configure swagger info so the generated UI points to local server (placeholder values; `swag init` will overwrite real docs)
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.BasePath = "/"

	//处理跨域
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},                                                                             // 允许的来源域名，使用 * 表示允许任何域名
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},                      // 允许的 HTTP 方法
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Requested-With"}, // 允许的 HTTP 头
		ExposeHeaders: []string{"Content-Length", "Content-Type"},                                                // 暴露给浏览器的头信息
		//AllowCredentials: true,                                                                                      // 是否允许发送认证信息（cookies）
		MaxAge: 12 * time.Hour, // 预检请求的有效期
	}))
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Next()
	})

	// 系统根路径 http://localhost:8080/
	r.GET("/", index.Index)

	// swagger 文档路径 http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// /login 及其子路径不需要验证
	{
		loginGroup := r.Group("/login")
		loginGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "login.go root"})
		})

		loginGroup.POST("/index", login.Login)
		loginGroup.GET("/captcha", login.GetCaptcha)
		loginGroup.GET("/captcha_img", login.GetCaptchaImg)
	}

	{
		// /user 及其子路径需要JWT验证
		//authGroup := r.Group("/user", middleware.JWTAuth())
		authGroup := r.Group("/user")

		// 获取用户信息路由
		authGroup.GET("/userInfo", users.UserInfo)

		//authGroup.POST("/userInfo", users.UserInfo)

	}

	return r
}
