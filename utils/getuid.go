package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// GetUserIDFromToken 从 Token 中获取用户 ID
func GetUserIDFromToken(c *gin.Context) (int, error) {
	// 从上下文中获取 userid
	userid, exists := c.Get("userid")
	if !exists {
		return 0, errors.New("未授权")
	}

	// 类型断言
	uid, ok := userid.(int)
	if !ok {
		return 0, errors.New("用户ID类型错误")
	}

	return uid, nil
}
