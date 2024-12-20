package controller

import (
	"InterLibrarySystem/models"
	"InterLibrarySystem/repository"
	"InterLibrarySystem/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	//判断符合书籍是否空
	if len(books) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "暂无符合书籍",
		})
	} else {
		// 返回结果
		c.JSON(http.StatusOK, gin.H{
			"code":  1,
			"books": books,
		})
	}
}

// SearchBookByID 根据书籍 ID 查询书籍信息
func SearchBookByID(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("book_id"))

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

// CreateBook 创建书籍
func CreateBook(c *gin.Context) {
	var book models.Book
	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	err = bookService.CreateBook(&book)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "创建成功",
	})
}

// UpdateBook 更改书籍信息
func UpdateBook(c *gin.Context) {
	var book models.Book
	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	err = bookService.UpdateBook(&book)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "修改成功",
	})

}

// DeleteBook 删除书籍
func DeleteBook(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("book_id"))
	err := bookService.DeleteBookByBookID(bookID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "删除成功",
	})
}

// ShowBookShortage 显示缺书记录
func ShowBookShortage(c *gin.Context) {
	bookShortage, err := bookService.ShowBookShortage()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	if len(bookShortage) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "暂无缺书",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  bookShortage,
		})
	}
}

// CreateBookShortage 创建缺书记录
func CreateBookShortage(c *gin.Context) {
	var bookShortage models.BookShortage
	err := c.BindJSON(&bookShortage)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	//获取外键bookID
	book, err := bookService.GetBookByName(bookShortage.Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	bookShortage.BookID = book.ID

	err = bookService.CreateBookShortage(&bookShortage)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "创建成功",
	})
}
