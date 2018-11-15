package service

import (
	"encoding/json"
	"fmt"
	"mabis/opensource/opengoods/config"
	"mabis/opensource/opengoods/db"
	"mabis/opensource/opengoods/entity"

	"github.com/gin-gonic/gin"
)

type response struct {
	Errcode     int         `json:"errcode"`
	Description string      `json:"description"`
	Payload     interface{} `json:"payload"`
}

type queryClause struct {
	where    interface{}
	pageNo   int
	pageSize int
}

func (rsp *response) ToString() string {
	data, err := json.Marshal(rsp)

	if nil != err {
		return ""
	}

	return string(data)
}

type engine struct {
	db *db.DbContext
}

var eng *engine

const (
	groupcategory string = "category"
	grouporigin   string = "origin"
	groupgoods    string = "goods"
)

//handle add category
func categroryHandleAdd(c *gin.Context) {
	var category entity.Category
	rsp := &response{
		Errcode:     200,
		Description: "success",
	}
	err := c.BindJSON(&category)

	if nil != err {
		rsp.Errcode = 400
		rsp.Description = "invalid category data format"
		c.JSON(200, rsp.ToString())
		return
	}

	if nil == eng {
		// c.JSON(200, gin.H{"errcode": 500, "description": "fatal errors in service"})
		rsp.Errcode = 500
		rsp.Description = "fatal errors in service"
		c.JSON(200, rsp.ToString())
		return
	}

	err = eng.db.Save(category)

	if nil != err {
		rsp.Errcode = 501
		rsp.Description = "internal database error" + err.Error()
		// c.JSON(200, gin.H{"errcode": 501, "description": "internal database error " + err.Error()})
		return
	}

	c.JSON(200, rsp.ToString())
}

//handle query category
func categroryHandleQuery(c *gin.Context) {
	var query queryClause
	err := c.BindJSON(&query)
	rsp := &response{
		Errcode:     200,
		Description: "success",
	}

	if nil != err {
		rsp.Errcode = 400
		rsp.Description = "invalid category data format"
		c.JSON(200, rsp.ToString())
		// c.JSON(200, gin.H{"errcode": 400, "description": "invalid category data format"})
	}

	rsp.Payload, err = eng.db.Query(query.where, query.pageNo, query.pageSize)

	if nil != err {
		rsp.Errcode = 401
		rsp.Description = "internal error, query database failed"
		c.JSON(200, rsp.ToString())
	}

	c.JSON(200, rsp.ToString())
}

//handle delete category
func categroryHandleDelete(c *gin.Context) {
	var category entity.Category
	err := c.BindJSON(&category)
	rsp := &response{
		Errcode:     200,
		Description: "success",
	}

	if nil != err {
		rsp.Errcode = 400
		rsp.Description = "invalid category data format"
		c.JSON(200, rsp.ToString())
		c.JSON(200, gin.H{"errcode": 400, "description": "invalid category data format"})
	}
}

//handle add category
func originHandleAdd(c *gin.Context) {
	var category entity.Category
	rsp := &response{
		Errcode:     200,
		Description: "success",
	}
	err := c.BindJSON(&category)

	if nil != err {
		rsp.Errcode = 400
		rsp.Description = "invalid category data format"
		c.JSON(200, rsp.ToString())
		return
	}

	if nil == eng {
		// c.JSON(200, gin.H{"errcode": 500, "description": "fatal errors in service"})
		rsp.Errcode = 500
		rsp.Description = "fatal errors in service"
		c.JSON(200, rsp.ToString())
		return
	}

	err = eng.db.Save(category)

	if nil != err {
		rsp.Errcode = 501
		rsp.Description = "internal database error" + err.Error()
		c.JSON(200, rsp.ToString())
		// c.JSON(200, gin.H{"errcode": 501, "description": "internal database error " + err.Error()})
		return
	}

	c.JSON(200, rsp.ToString())
}

//handle query category
func originHandleQuery(c *gin.Context) {
	var category entity.Category
	err := c.BindJSON(&category)
	rsp := &response{
		Errcode:     200,
		Description: "success",
	}

	if nil != err {
		rsp.Errcode = 400
		rsp.Description = "invalid category data format"
		c.JSON(200, rsp.ToString())
		// c.JSON(200, gin.H{"errcode": 400, "description": "invalid category data format"})
	}
}

//handle delete category
func originHandleDelete(c *gin.Context) {
	var category entity.Category
	err := c.BindJSON(&category)
	rsp := &response{
		Errcode:     200,
		Description: "success",
	}

	if nil != err {
		rsp.Errcode = 400
		rsp.Description = "invalid category data format"
		c.JSON(200, rsp.ToString())
		// c.JSON(200, gin.H{"errcode": 400, "description": "invalid category data format"})
	}
}

//handle add category
func goodsHandleAdd(c *gin.Context) {
	var category entity.Category
	rsp := &response{
		Errcode:     200,
		Description: "success",
	}
	err := c.BindJSON(&category)

	if nil != err {
		rsp.Errcode = 400
		rsp.Description = "invalid category data format"
		c.JSON(200, rsp.ToString())
		return
	}

	if nil == eng {
		// c.JSON(200, gin.H{"errcode": 500, "description": "fatal errors in service"})
		rsp.Errcode = 500
		rsp.Description = "fatal errors in service"
		c.JSON(200, rsp.ToString())
		return
	}

	err = eng.db.Save(category)

	if nil != err {
		rsp.Errcode = 501
		rsp.Description = "internal database error" + err.Error()
		// c.JSON(200, gin.H{"errcode": 501, "description": "internal database error " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"errcode": 200, "description": "success"})
}

//handle query category
func goodsHandleQuery(c *gin.Context) {
	var category entity.Category
	err := c.BindJSON(&category)
	rsp := &response{
		Errcode:     200,
		Description: "success",
	}

	if nil != err {
		// c.JSON(200, gin.H{"errcode": 400, "description": "invalid category data format"})
		rsp.Errcode = 400
		rsp.Description = "invalid category data format"
		c.JSON(200, rsp.ToString())
	}

	if nil == eng {
		// c.JSON(200, gin.H{"errcode": 500, "description": "fatal errors in service"})
		rsp.Errcode = 500
		rsp.Description = "fatal errors in service"
		c.JSON(200, rsp.ToString())
		return
	}
}

//handle delete category
func goodsHandleDelete(c *gin.Context) {
	var category entity.Category
	err := c.BindJSON(&category)
	rsp := &response{
		Errcode:     200,
		Description: "success",
	}

	if nil != err {
		rsp.Errcode = 400
		rsp.Description = "invalid category data format"
		c.JSON(200, rsp.ToString())
		// c.JSON(200, gin.H{"errcode": 400, "description": "invalid category data format"})
	}

	if nil == eng {
		// c.JSON(200, gin.H{"errcode": 500, "description": "fatal errors in service"})
		rsp.Errcode = 500
		rsp.Description = "fatal errors in service"
		c.JSON(200, rsp.ToString())
		return
	}

}

func main() {
	serverConfig := config.ServerCfg()
	fmt.Println("serverconfig is:", serverConfig.ToString())
	dbConfig := config.DBCfg()
	fmt.Println("database config is:", dbConfig.ToString())

	route := gin.Default()
	rCategory := route.Group(groupcategory)
	{
		rCategory.POST("add", categroryHandleAdd)
		rCategory.POST("query", categroryHandleQuery)
		rCategory.POST("delete", categroryHandleDelete)
	}

	rOrigin := route.Group(grouporigin)
	{
		rOrigin.POST("add", originHandleAdd)
		rOrigin.POST("query", originHandleQuery)
		rOrigin.POST("delete", originHandleDelete)
	}

	rGoods := route.Group(groupgoods)
	{
		rGoods.POST("add", goodsHandleAdd)
		rGoods.POST("query", goodsHandleQuery)
		rGoods.POST("delete", goodsHandleDelete)
	}

	route.Run(serverConfig.Url + ":" + string(serverConfig.Port))
}
