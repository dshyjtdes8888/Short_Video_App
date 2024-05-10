package main

import (
	"github.com/gin-gonic/gin"
	"ningmeng/data"
)

func main() {

	//初始化数据库
	go data.InitDb()

	// 创建一个默认的gin.Engine实例，用于构建HTTP路由。
	r := gin.Default()

	// 初始化路由，调用initRouter函数，用于设置API路由及其处理函数。
	InitRouter(r)

	r.Run()
}
