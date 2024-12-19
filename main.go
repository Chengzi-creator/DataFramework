package main

import (
	"InterLibrarySystem/controller"
	"InterLibrarySystem/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	//UserGroup := r.Group("/user", middleware.AuthMiddleware())  需验证token
	UserGroup := r.Group("/user")
	{
		UserGroup.GET("/index")
		UserGroup.GET("/search", controller.SearchBooks)
		{
			UserGroup.GET("/search/:book_id", controller.SearchBookByID)
			UserGroup.GET("/search/:book_id", controller.C)
		}
		UserGroup.GET("/ticket", controller.GetTicketsByUserID)
	}

	//管理员
	//AdministerGroup := r.Group("/administer", middleware.AuthMiddleware())需验证token
	AdministerGroup := r.Group("/administer")
	{
		AdministerGroup.GET("/index")
	}

	//启动服务
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
