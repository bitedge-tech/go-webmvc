package handler

import (
	"go-webmvc/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Success 通用成功响应
func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, &dto.BaseResponse{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

// Failed 通用失败响应
func Failed(ctx *gin.Context, msg string) {
	sendMsg := "failed"
	if msg != "" {
		sendMsg = msg
	}
	ctx.JSON(http.StatusOK, &dto.BaseResponse{
		Code: 1,
		Msg:  sendMsg,
		Data: nil,
	})
}
