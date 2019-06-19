package models

import (
	"time"

	db "github.com/teed7334-restore/homekeeper/database"
)

//Leaves 休假申請物件
type Leaves struct {
	ID        int
	Startdate time.Time
	Enddate   time.Time
	Starttime string
	Endtime   string
}

//SetUseMinuteToZero 將使用時間設為零
func SetUseMinuteToZero() {
	sql := "UPDATE leaves SET useMinute = ?"
	db.Db.Exec(sql, 0)
	sql = "UPDATE leaves_history SET useMinute = ?"
	db.Db.Exec(sql, 0)
}

//SetUseMinute 設定休假花費時間
func SetUseMinute(useMinute int, ID int) {
	sql := "UPDATE leaves SET useMinute = ? WHERE id = ?"
	db.Db.Exec(sql, useMinute, ID)
	sql = "UPDATE leaves_history SET useMinute = ? WHERE id = ?"
	db.Db.Exec(sql, useMinute, ID)
}

//GetLeaveHistory 取得假期日誌
func GetLeaveHistory() []*Leaves {
	list := []*Leaves{}
	sql := "SELECT id, startdate, enddate, starttime, endtime FROM leaves"
	db.Db.Raw(sql).Scan(&list)
	return list
}
