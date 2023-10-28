package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//	app.Start()
	//	model.NewMysql()
	//
	//	g.GET("/login", func(context *gin.Context) {
	//		fmt.Println("Hello World ")
	//		context.HTML(http.StatusOK, "login.tmpl", nil)
	//	})
	//	g.POST("/login", func(context *gin.Context) {
	//		var user User
	//		ret := make(map[string]string)
	//		_ = context.ShouldBind(&user)
	//		err := conn.Table("user").Where("name = ? and password = ?", user.Name, user.Password).Find(&user)
	//		if err != nil {
	//			context.JSON(http.StatusBadGateway, map[string]string{
	//				"msg": "用户名或密码错误",
	//			})
	//		}
	//		context.JSON(http.StatusOK, ret)
	//	})
	//	if err := g.Run(":8080"); err != nil {
	//		fmt.Println(err)
	//	}
	//
	//application.Start()
}
