package main

import (
	"keyu.tech/website/core"
	"keyu.tech/website/global"
	"keyu.tech/website/routers"
)

func main() {
	// 读取配置文件
	core.InitServer()

	// 初始化日志
	global.Log = core.InitLogger()

	// 连接数据库
	global.DB = core.InitGorm()

	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	router.Run(addr)
}
