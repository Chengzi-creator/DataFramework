package middleware

import (
	"InterLibrarySystem/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware 验证token中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的 Token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 0, "msg": "未登录，请提供 Token"})
			c.Abort()
			return
		}

		// 验证 Token 并解析 Claims
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 0, "msg": "无效的 Token"})
			c.Abort()
			return
		}

		// 将解析出来的用户信息存入上下文
		c.Set("userid", claims.UserID)
		c.Set("username", claims.Username)
		c.Next() // 继续执行后续请求
	}
}
