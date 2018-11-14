package db

import (
	"errors"
	"fmt"
	"mabis/opensource/opengoods/config"

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
	return nil
}
