package controller

import (
	"InterLibrarySystem/repository"
	"InterLibrarySystem/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var bookService = service.BookService{
	Repo: &repository.BookRepository{},
}

// SearchBooks 根据书名查询符合书籍
func SearchBooks(c *gin.Context) {
	// 获取查询数据
	name := c.Query("book_name")

	// 调用 Service 层查询书籍
	books, err := bookService.SearchBooks(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"books": books,
	})
}

// SearchBookByID 根据书籍 ID 查询书籍信息
func SearchBookByID(c *gin.Context) {
	bookID := c.Param("book_id")

	// 调用 Service 层
	book, err := bookService.GetBookByID(bookID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"book": book,
	})
}
