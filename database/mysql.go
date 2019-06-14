package database

import (
	"fmt"

	"github.com/teed7334-restore/homekeeper/env"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Db 資料庫連結器
var Db *gorm.DB

//Err 錯誤處理器
var Err error

func init() {
	db := env.GetEnv().Database
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=%s&loc=%s", db.User, db.Password, db.Host, db.Database, db.Charset, db.ParseTime, db.Loc)
	Db, Err = gorm.Open("mysql", dsn)
	if Err != nil {
		panic(Err)
	}
}
