package db

import (
	"errors"
	"fmt"
	"mabis/opensource/opengoods/config"
	"mabis/opensource/opengoods/entity"
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*DbContext context of the database
 */
type DbContext struct {
	Db *gorm.DB
}

/*New create a db  context
DBConfig config of database
return a instance of database context when there is no error happen, otherwise return nil
*/
func New(config *config.DBConfig) *DbContext {

	connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.User, config.Password, config.Host,
		config.Port, config.Database, config.Charset)
	db, err := gorm.Open("mysql", connectString)

	if nil != err {

		return nil
	}

	return &DbContext{
		db,
	}
}

/*Release free the resources
description: release a database context
*/
func (dbctx *DbContext) Release() {
	if dbctx != nil {
		dbctx.Db.Close()
	}
}

/*Save save a new item to database
param: interface{} object to save
*/
func (dbctx *DbContext) Save(v interface{}) error {

	if nil == v {
		return errors.New("parameter null pointer")
	}

	dbctx.Db = dbctx.Db.Create(v)
	dbctx.Db = dbctx.Db.Commit()
	err := dbctx.Db.Error
	return err
}

/*Delete delete a item to database
param: interface{} object to delete
*/
func (dbctx *DbContext) Delete(v interface{}) error {
	if nil == v {
		return errors.New("parameter null pointer")
	}

	dbctx.Db = dbctx.Db.Delete(v)
	dbctx.Db = dbctx.Db.Commit()
	err := dbctx.Db.Error
	return err
}

/*Query Query specify items from database
param: interface{} where clause
param:pageno  no of current page
param:pagesize size of each page
return slice of objects and error when error is nil query success, others fail
*/
func (dbctx *DbContext) Query(where interface{}, pageno int, pagesize int) (interface{}, error) {
	if nil == where {
		return nil, errors.New("parameter null pointer")
	}

	var clist interface{}
	var err error
	fmt.Println("where type name: ", reflect.TypeOf(where), ">>", reflect.TypeOf(where).Name())

	switch reflect.TypeOf(where).Name() {
	case "Category":
		clist = make([]entity.Category, 0)
		err = dbctx.Db.Where(where).Offset(pageno * pagesize).Limit(pagesize).Find(clist).Error
	case "Goods":
		clist = make([]entity.Goods, 0)
		err = dbctx.Db.Where(where).Offset(pageno * pagesize).Limit(pagesize).Find(clist).Error

	case "Origin":
		clist = make([]entity.Origin, 0)
		err = dbctx.Db.Where(where).Offset(pageno * pagesize).Limit(pagesize).Find(clist).Error
	}

	return clist, err
}
