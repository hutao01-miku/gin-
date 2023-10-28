package controllers

import "github.com/gin-gonic/gin"

type UserController struct {
}

func (u UserController) GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	ReturnSuccess(c, 0, name, id, 1)
}
func (u UserController) GetList(c *gin.Context) {
	ReturnError(c, 404, "没有相关信息list")
}
