package users

import (
	"go-webmvc/internal/dto"
	"go-webmvc/internal/handler"
	"go-webmvc/internal/service"

	"github.com/gin-gonic/gin"
)

// UserInfo 获取用户信息的处理函数 (Get 请求)
// @Summary 获取用户信息
// @Description 根据用户ID获取用户的详细信息
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param user_id query int true "用户ID"
// @Success 200 {object} dto.UserInfoResponse "成功返回用户信息"
// @Failure 400 {object} dto.BaseResponse "参数绑定失败或请求错误"
// @Failure 500 {object} dto.BaseResponse "服务器内部错误"
// @Router /user/userInfo [get]
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

// UserInfo2 使用 POST 方法获取用户信息
func UserInfo2(c *gin.Context) {

	// 1. 参数绑定:从客户端请求中获取用户ID
	req := dto.UserInfoRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
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
