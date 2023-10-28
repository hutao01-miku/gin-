package router

import (
	"awesomeProject1/demo/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.GET("/info/:id/:name", controllers.UserController{}.GetUserInfo)
		user.POST("/list", controllers.UserController{}.GetList)
		user.PUT("/add", func(c *gin.Context) {
			c.String(200, "user add")
		})
		user.DELETE("/delete", func(c *gin.Context) {
			c.String(200, "user delete")
		})
	}
	order := r.Group("/order")
	{
		order.GET("/info", controllers.OrderController{}.GetUserInfo)
		order.POST("/list", controllers.OrderController{}.GetList)
	}
	return r
}
