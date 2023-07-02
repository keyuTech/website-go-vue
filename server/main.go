package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"keyu.tech/website/core"
	"keyu.tech/website/global"
)

func main() {
	// 读取配置文件
	core.InitServer()
	fmt.Println(global.Config)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
