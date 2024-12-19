package controller

import (
	"InterLibrarySystem/repository"
	"InterLibrarySystem/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 初始化 TicketService
var ticketService = service.TicketService{
	Repo: &repository.TicketRepository{},
}

// GetTicketsByUserID  获取用户订单
func GetTicketsByUserID(c *gin.Context) {
	// 从 token 中获取用户ID
	userid, exists := c.Get("userid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 0,
			"msg":  "未授权",
		})
		return
	}

	// 类型断言
	uid, ok := userid.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 0,
			"msg":  "用户ID类型错误",
		})
		return
	}

	// 调用 Service 层获取订单
	tickets, err := ticketService.GetTicketsByUserID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}

	// 判断订单是否为空
	if len(tickets) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "暂无订单记录",
		})
		return
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"tickets": tickets,
	})
}
