package middleware

import (
	config2 "backend/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Mysql() gin.HandlerFunc {
	config := config2.Mysql()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Database, config.Charset)
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功")
	return func(c *gin.Context) {

	}
}
