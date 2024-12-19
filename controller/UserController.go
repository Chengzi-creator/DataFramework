package controller

import (
	"InterLibrarySystem/repository"
	"InterLibrarySystem/service"
	"InterLibrarySystem/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var userService = service.UserService{
	UserRepo: &repository.UserRepository{},
}

// Login 登录
func Login(c *gin.Context) {
	username := c.PostForm("username") //获取username
	password := c.PostForm("password") //获取password

	//调用service层进行登录操作
	user, err := userService.Login(username, password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}

	//生成token
	token, err := utils.GenerateToken(user.ID, user.Username)
	c.JSON(http.StatusOK, gin.H{
		"code":          1,
		"msg":           "登陆成功",
		"token":         token,
		"is_administer": user.IsAdminister,
	})
	return
}
