package utils

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// InitMySQL 连接数据源库
func InitMySQL() (err error) {
	dsn := "kanolity:1234@tcp(localhost:3306)/interlibrarysystem?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return DB.DB().Ping()
}
