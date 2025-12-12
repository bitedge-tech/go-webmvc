package users

import (
	"go-webmvc/internal/dto"
	"go-webmvc/internal/handler"
	"go-webmvc/internal/service"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {

	// 1. 参数绑定:从客户端请求中获取用户ID
	req := dto.UserInfoRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		handler.Failed(c, "参数绑定失败")
		return
	}

	// 2. 调用服务层获取用户信息
	userService := service.Services.User
	userInfo, err := userService.UserInfo(req.UserID)
	if err != nil {
		handler.Failed(c, "获取用户信息失败")
		return
	}

	// 3. 返回成功响应给客户端
	if userInfo == nil {
		handler.Failed(c, "用户不存在")
		return
	}
	handler.Success(c, userInfo)

}
