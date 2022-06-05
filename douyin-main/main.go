package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hjk-cloud/douyin/models"
	"github.com/hjk-cloud/douyin/routers"
)

// @title api
// @version 1.0
// @description douyin
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath 这里写base path
func main() {
	models.Init()

	r := gin.Default()

	routers.InitRouter(r)

	r.Run()
}
