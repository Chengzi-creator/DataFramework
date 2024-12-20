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

	//注册界面
	r.POST("/register", controller.Register)
	//登录界面
	r.POST("/login", controller.Login)

	//普通用户
	UserGroup := r.Group("/user")
	{
		UserGroup.GET("/index")                          //菜单
		UserGroup.GET("/search", controller.SearchBooks) //搜索
		{
			UserGroup.GET("/search/:book_id", controller.SearchBookByID) //书籍详情
			UserGroup.POST("/search/:book_id", controller.CreateTicket)  //订单
		}
		UserGroup.GET("/ticket", controller.GetTicketsByUserID) //订单详情
		UserGroup.GET("/userinfo", controller.ShowUserinfo)     //用户信息
		UserGroup.PUT("/userinfo", controller.ChangeByUserID)   //更改信息
	}

	//管理员
	AdministerGroup := r.Group("/administer")
	{
		AdministerGroup.GET("/index")                               //菜单
		AdministerGroup.POST("/create_book", controller.CreateBook) //添加书籍
		AdministerGroup.GET("/search", controller.SearchBooks)      //搜索
		{
			AdministerGroup.GET("/search/:book_id", controller.SearchBookByID) //书籍详情
			AdministerGroup.PUT("/search/:book_id", controller.UpdateBook)     //更改书籍
			AdministerGroup.DELETE("/search/:book_id", controller.DeleteBook)  //删除书籍
		}
		AdministerGroup.GET("/book_shortage", controller.ShowBookShortage)    //缺书信息
		AdministerGroup.POST("/book_shortage", controller.CreateBookShortage) //新增缺书
		AdministerGroup.GET("/supplier_info", controller.ShowSupplierInfo)    //供应商信息
		AdministerGroup.POST("/purchase", controller.CreatePurchase)          //购书
	}

	//启动服务
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
