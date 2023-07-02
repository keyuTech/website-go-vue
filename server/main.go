package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"keyu.tech/website/core"
	"keyu.tech/website/global"
)

func main() {
	// 读取配置文件
	core.InitServer()

	// 初始化日志
	global.Log = core.InitLogger()
	global.Log.Warnln("hhhhhhh")
	global.Log.Errorln("hhhhhhh")

	// 连接数据库
	global.DB = core.InitGorm()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
