package models

import (
	"log"
	"time"

	db "github.com/teed7334-restore/homekeeper/database"
)

//Holiday 假期資料表
type Holiday struct {
	ID              int `gorm:"AUTO_INCREMENT"`
	Date            time.Time
	Name            string
	IsHoliday       int
	HolidayCategory string
	Description     string
}

//GetHoliday 取得休假日資料
func GetHoliday() []*Holiday {
	list := []*Holiday{}
	err := db.Db.Where("is_holiday = ?", 1).Find(&list).Error
	if err != nil {
		log.Fatal(err)
	}
	return list
}
