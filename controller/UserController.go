package controller

import (
	"InterLibrarySystem/models"
	"InterLibrarySystem/repository"
	"InterLibrarySystem/service"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strconv"
)

var userService = service.UserService{
	UserRepo: &repository.UserRepository{},
}

// Login 登录
func Login(c *gin.Context) {
	username := c.PostForm("username") //获取username
	password := c.PostForm("password") //获取password

	//进行登录操作
	user, err := userService.Login(username, password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":          1,
		"msg":           "登陆成功",
		"is_administer": user.IsAdminister,
	})
}

// ShowUserinfo 展示用户信息
func ShowUserinfo(c *gin.Context) {
	//获取userid
	userid := c.Query("userid")
	uid, _ := strconv.Atoi(userid)
	//查询信息
	user, err := userService.FindByUserID(uid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"user": user,
	})

}

// ChangeByUserID 根据用户ID更改用户信息
func ChangeByUserID(c *gin.Context) {
	//获取userid
	userid := c.Query("userid")
	uid, _ := strconv.Atoi(userid)
	//获取更改后的信息
	var newUser models.User
	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	//查询信息
	user, err := userService.FindByUserID(uid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	//更改信息
	user.Username = newUser.Username
	user.Password = newUser.Password
	user.Address = newUser.Address
	err = userService.UpdateByUserID(user)
}

// Register 注册
func Register(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	err = userService.RegisterNewUser(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "注册成功",
	})
}
