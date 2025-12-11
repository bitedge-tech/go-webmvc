package middleware

import (
	"net/http"
	"strings"

	"go-webmvc/internal/util"

	"github.com/gin-gonic/gin"
)

// JWTAuth 这里可以根据实际项目替换为你自己的 JWT 校验逻辑
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "缺少Token"})
			return
		}
		// 这里假设token格式为 "Bearer xxx"，实际可根据你的业务解析和校验
		parts := strings.SplitN(token, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" || parts[1] == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token格式错误"})
			return
		}
		claims, err := util.ParseToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token无效: " + err.Error()})
			return
		}
		// 可将用户信息写入上下文
		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}
