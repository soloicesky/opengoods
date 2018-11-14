package service

import (
	"fmt"
	"mabis/opensource/opengoods/config"
	"mabis/opensource/opengoods/entity"

	"github.com/gin-gonic/gin"
)

const (
	groupcategory string = "category"
	grouporigin   string = "origin"
	groupgoods    string = "goods"
)

func categrory_handle_add(c *gin.Context) {
	var category entity.Category
	err := c.BindJSON(&category)

	if nil != err {
		c.JSON(200, gin.H{"errcode": 400, "description": "invalid category data format"})
	}
}

func categrory_handle_query(c *gin.Context) {
	var category entity.Category
	err := c.BindJSON(&category)

	if nil != err {
		c.JSON(200, gin.H{"errcode": 400, "description": "invalid category data format"})
	}
}

func categrory_handle_delete(c *gin.Context) {
	var category entity.Category
	err := c.BindJSON(&category)

	if nil != err {
		c.JSON(200, gin.H{"errcode": 400, "description": "invalid category data format"})
	}
}

func main() {
	serverConfig := config.ServerCfg()
	fmt.Println("serverconfig is:", serverConfig.ToString())
	dbConfig := config.DBCfg()
	fmt.Println("database config is:", dbConfig.ToString())

	route := gin.Default()
	r_category := route.Group(groupcategory)
	{
		r_category.POST("add", categrory_handle_add)
		r_category.POST("query", categrory_handle_query)
		r_category.POST("delete", categrory_handle_delete)
	}

	route.Run(serverConfig.Url + ":" + string(serverConfig.Port))
}
