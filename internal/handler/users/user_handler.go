package users

import (
	"go-webmvc/internal/handler"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	handler.Success(c, nil)
}
