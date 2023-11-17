package main

import "github.com/gin-gonic/gin" //导包

func Hello(c *gin.Context) {
	//gin.Context是gin框架的上下文，封装了request(请求)和response(响应)
	//里边是对页面的渲染效果，即给浏览器响应什么内容，可以是string的内容，也可以是json的内容，也可以是html的内容
	c.String(200, "这是我的第一个gin框架程序")
}
func main() {
	//Default返回一个默认的路由引擎Engine,他是框架非常重要的数据结构，是框架的入口，
	//所有的注册路由都是在这个结构体上进行的，所以在使用gin框架时，一定要先初始化一个Engine
	//引擎-框架核心发动机-默认服务器，整个web服务都由它驱动
	//Default底层调用了New(),相当于定义的升级版，多加了中间件-engine.Use(Logger(),Recovery())
	r := gin.Default()
	//r := gin.New()
	//路由：通过访问"/"的GET请求，来执行后面的匿名函数
	//前面“/”是路由规则，后边是路由函数，是一个协程
	//路由请求方式：GET,POST,PUT,DELETE,HEAD,OPTIONS,PATCH,ANY
	r.GET("/", func(c *gin.Context) {
		c.String(200, "这是我的第一个gin框架程序")
	})
	//即函数可以直接写成匿名函数，还可以外部定义函数，然后传入
	//或 写成 r.GET("/", Hello)

	//启动引擎，服务器启动传入host+port，默认8080端口
	r.Run(":8080")
}
