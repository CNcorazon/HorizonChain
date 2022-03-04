package main

import (
	"horizon/server/route"

	"github.com/sirupsen/logrus"
)

func main() {
	// Logrus是一个用的非常广泛的Go语言的日志包
	// 可以查看这个文章学习一下 https://juejin.cn/post/6844904061393698823
	logrus.SetLevel(logrus.TraceLevel)

	// 我使用的框架是Gin框架，是一个非常简单高效的Go Web Server 框架
	// 可以看下这个文章学习一下 https://cloud.tencent.com/developer/article/1739957
	//初始化router服务
	r := route.InitRoute()
	r.Run() //启动HTTP服务，默认在8080端口启动服务
}
