package controller

import (
	"InterLibrarySystem/models"
	"InterLibrarySystem/repository"
	"InterLibrarySystem/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 初始化 TicketService
var ticketService = service.TicketService{
	Repo: &repository.TicketRepository{},
}
var userServiceTicket = service.UserService{
	UserRepo: &repository.UserRepository{},
}

// GetTicketsByUserID  获取用户订单
func GetTicketsByUserID(c *gin.Context) {
	// 从 query 中获取用户ID
	userid := c.Query("user_id")
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

// CreateTicket 创建订单
func CreateTicket(c *gin.Context) {
	var req struct {
		userID   int    `form:"user_id" binding:"required"`
		bookID   int    `form:"book_id" binding:"required"`
		address  string `form:"address" binding:"required"`
		quantity int    `form:"quantity" binding:"required"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}

	// 从 BookService 获取书籍信息
	book, err := bookService.GetBookByID(req.bookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}

	//判断余额是否足够
	user, _ := userServiceTicket.FindByUserID(req.userID)
	price := float64(req.quantity) * book.Price
	var newBalance float64
	switch user.CreditRating {
	case 0:
		newBalance = user.Balance - price
	case 1:
		newBalance = user.Balance - price*0.9
	case 2, 3:
		newBalance = user.Balance - price*0.85
	case 4:
		newBalance = user.Balance - price*0.8
	case 5:
		newBalance = user.Balance - price*0.75
	}
	if newBalance < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "余额不足",
		})
		return
	}
	// 创建订单
	ticket := models.Ticket{
		Price:       newBalance,
		Time:        time.Now(),
		Quantity:    req.quantity,
		UserID:      req.userID,
		Address:     req.address,
		Description: book.Name,
	}
	err = ticketService.CreateTicket(ticket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	//扣除书费,更新用户信息
	user.Balance = newBalance
	err = userServiceTicket.UpdateByUserID(user)
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
