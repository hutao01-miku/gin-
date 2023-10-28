package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

type Search struct {
	Name string `json:"name"`
	Cid  int    `json:"cid"`
}

func (o OrderController) GetUserInfo(c *gin.Context) {
	ReturnSuccess(c, 0, "success", "order info", 1)
}
func (o OrderController) GetList(c *gin.Context) {
	//cid := c.PostForm("cid")
	//name := c.DefaultPostForm("name", "wangwu")
	//ReturnSuccess(c, 0, cid, name, 1)
	//param := make(map[string]interface{})
	//err := c.BindJSON(&param)
	search := &Search{}
	err := c.BindJSON(&search)
	if err != nil {
		ReturnSuccess(c, 0, search.Name, search.Cid, 1)
		return
	}
	ReturnError(c, 404, gin.H{"err": err})
}
