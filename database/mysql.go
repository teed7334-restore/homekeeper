package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

//Db 資料庫連結器
var Db *gorm.DB

//Err 錯誤處理器
var Err error

func init() {
	user := os.Getenv("database.user")
	passwd := os.Getenv("database.password")
	host := os.Getenv("database.host")
	database := os.Getenv("database.database")
	charset := os.Getenv("database.charset")
	parseTime := os.Getenv("database.parseTime")
	loc := os.Getenv("database.loc")
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=%s&loc=%s", user, passwd, host, database, charset, parseTime, loc)
	Db, Err = gorm.Open("mysql", dsn)
	if Err != nil {
		log.Println(Err)
	}
}
