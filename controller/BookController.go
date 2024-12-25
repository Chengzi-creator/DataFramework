package controller

import (
	"InterLibrarySystem/models"
	"InterLibrarySystem/repository"
	"InterLibrarySystem/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var bookService = service.BookService{
	Repo: &repository.BookRepository{},
}

// SearchBooks 查询符合书籍
func SearchBooks(c *gin.Context) {
	var books []models.Book
	var err error
	// //获取查询数据
	// name := c.Query("book_name")
	// seriesNo := c.Query("series_no")
	// publish := c.Query("publish")
	// keyword := c.Query("keyword")
	// writer := c.Query("writer")

	// // 使用 map 去重
	// 使用 map 去重
	uniqueBooks := make(map[int]models.Book)

	// 获取查询参数
	name := c.DefaultQuery("book_name", "")
	seriesNo := c.DefaultQuery("series_no", "")
	publish := c.DefaultQuery("publish", "")
	keyword := c.DefaultQuery("keyword", "")
	writer := c.DefaultQuery("writer", "")

	// 添加书籍到 uniqueBooks 中的辅助函数
	addBooks := func(newBooks []models.Book) {
		for _, book := range newBooks {
			uniqueBooks[book.ID] = book
		}
	}

	// 查询书名
	if name != "" {
		var books2 []models.Book
		books2, err = bookService.SearchBooksByName(name)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  err.Error(),
			})
			return
		}
		addBooks(books2)
	}

	// 查询系列号
	if seriesNo != "" {
		var books2 []models.Book
		books2, err = bookService.SearchBooksBySeriesNo(seriesNo)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  err.Error(),
			})
			return
		}
		addBooks(books2)
	}

	// 查询出版商
	if publish != "" {
		var books2 []models.Book
		books2, err = bookService.SearchBooksByPublish(publish)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  err.Error(),
			})
			return
		}
		addBooks(books2)
	}

	// 查询关键字
	if keyword != "" {
		var books2 []models.Book
		books2, err = bookService.SearchBooksByKeyword(keyword)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  err.Error(),
			})
			return
		}
		addBooks(books2)
	}

	// 查询作者
	if writer != "" {
		var books2 []models.Book
		books2, err = bookService.SearchBooksByWriter(writer)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  err.Error(),
			})
			return
		}
		addBooks(books2)
	}

	// 将去重后的书籍添加到 books 列表
	for _, book := range uniqueBooks {
		books = append(books, book)
	}

	// 返回结果
	if len(books) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "暂无符合书籍",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":  1,
			"books": books,
		})
	}

	//判断搜索方式并查询书籍
	// if name != "" {
	// 	books, err = bookService.SearchBooksByName(name)
	// } else if seriesNo != "" {
	// 	books, err = bookService.SearchBooksBySeriesNo(seriesNo)
	// } else if publish != "" {
	// 	books, err = bookService.SearchBooksByPublish(publish)
	// } else if keyword != "" {
	// 	books, err = bookService.SearchBooksByKeyword(keyword)
	// } else if writer != "" {
	// 	books, err = bookService.SearchBooksByWriter(writer)
	// } else {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"code": 0,
	// 		"msg":  "未传递任何有效的查询参数",
	// 	})
	// 	return
	// }
	// //处理错误
	// if err != nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": 0,
	// 		"msg":  err.Error(),
	// 	})
	// 	return
	// }
	// //判断符合书籍是否空
	// if len(books) == 0 {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": 1,
	// 		"msg":  "暂无符合书籍",
	// 	})
	// } else {
	// 	// 返回结果
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code":  1,
	// 		"books": books,
	// 	})
	// }
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
