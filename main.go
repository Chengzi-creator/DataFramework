package main

import (
	"InterLibrarySystem/controller"
	"InterLibrarySystem/middleware"
	"InterLibrarySystem/models"
	"InterLibrarySystem/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

func main() {
	//连接数据库
	err := utils.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer utils.DB.Close()

	//创建路由引擎
	r := gin.Default()

	//登录界面
	r.POST("/login", controller.Login)

	//普通用户
	UserGroup := r.Group("/user", middleware.AuthMiddleware())
	{
		UserGroup.GET("/index")
		UserGroup.GET("/search", controller.SearchBooks)
		{
			UserGroup.GET("/search/:book_id", controller.SearchBookByID)
		}
		UserGroup.GET("/ticket", func(c *gin.Context) {
			//从token中获得userid
			userid, exists := c.Get("userid")
			if !exists {
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "未授权",
				})
				return
			}
			userid = userid.(int)
			//根据userid查询ticket
			var ticket []models.Ticket
			err := utils.DB.Where("user_id=?", userid).Find(&ticket).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 0,
					"msg":  "查询失败",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"code":   1,
				"ticket": ticket,
			})
		})
	}

	//管理员
	AdministerGroup := r.Group("/administer", middleware.AuthMiddleware())
	{
		AdministerGroup.GET("/index")
	}

	//启动服务
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
