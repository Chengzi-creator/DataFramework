package controller

import (
	"InterLibrarySystem/models"
	"InterLibrarySystem/repository"
	"InterLibrarySystem/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// 初始化 TicketService
var ticketService = service.TicketService{
	Repo: &repository.TicketRepository{},
}

// GetTicketsByUserID  获取用户订单
func GetTicketsByUserID(c *gin.Context) {
	// 从 token 中获取用户ID
	userid := c.Query("userid")
	uid, _ := strconv.Atoi(userid)

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

func CreateTicket(c *gin.Context) {
	// 获取书籍 ID
	bookId, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"msg":  "书籍ID无效",
		})
		return
	}

	// 获取订单数量
	quantity, err := strconv.Atoi(c.PostForm("quantity"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"msg":  "数量无效",
		})
		return
	}

	// 获取地址
	address := c.PostForm("address")

	// 从 query中获取用户 ID
	userid := c.Query("userid")
	uid, _ := strconv.Atoi(userid)

	// 从 BookService 获取书籍信息
	book, err := bookService.GetBookByID(bookId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}

	// 创建订单
	ticket := models.Ticket{
		Price:       book.Price,
		Time:        time.Now(),
		Quantity:    quantity,
		UserID:      uid,
		Address:     address,
		Description: book.Name,
	}

	// 调用 Service 创建订单
	err = ticketService.CreateTicket(ticket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "订单创建成功",
	})
}
