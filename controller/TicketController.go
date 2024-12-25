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
	//获取userid
	userid, _ := strconv.Atoi(c.Param("user_id"))
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

	// 从 BookService 获取书籍信息
	book, err := bookService.GetBookByID(bookId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}

	//判断余额是否足够
	user, _ := userServiceTicket.FindByUserID(userid)
	price := float64(quantity) * book.Price
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
	}
	// 创建订单
	ticket := models.Ticket{
		Price:       book.Price,
		Time:        time.Now(),
		Quantity:    quantity,
		UserID:      userid,
		Address:     address,
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
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "订单创建成功",
	})
}
