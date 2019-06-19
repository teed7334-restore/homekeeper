package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/teed7334-restore/homekeeper/beans"
	"github.com/teed7334-restore/homekeeper/env"
	"github.com/teed7334-restore/homekeeper/models"
)

//CalcTime 計算時數
func CalcTime(c *gin.Context) {
	params := getPunchclockParams(c)
	beginYear := params.GetBegin().GetYear()
	beginMonth := params.GetBegin().GetMonth()
	beginDay := params.GetBegin().GetDay()
	beginHour := params.GetBegin().GetHour()
	beginMinute := params.GetBegin().GetMinute()
	endYear := params.GetEnd().GetYear()
	endMonth := params.GetEnd().GetMonth()
	endDay := params.GetEnd().GetDay()
	endHour := params.GetEnd().GetHour()
	endMinute := params.GetEnd().GetMinute()

	beginTime, _ := strconv.Atoi(beginHour + beginMinute)
	endTime, _ := strconv.Atoi(endHour + endMinute)

	diffDay := 0
	diffHour := 0
	diffMinute := 0

	beginHour, beginMinute, endHour, endMinute = changeToLunchTime(beginTime, endTime, beginHour, beginMinute, endHour, endMinute)

	if isSameDay(params) { //只請當天
		if beginTime < 1200 && endTime > 1330 { //有跨午休
			diffHour, diffMinute = calcTime(beginHour, endHour, beginMinute, endMinute, true)
		} else { //沒跨午休
			diffHour, diffMinute = calcTime(beginHour, endHour, beginMinute, endMinute, false)
		}
		_beginMonth := appendZero(beginMonth)
		_beginDay := appendZero(beginDay)
		if isHoliday(beginYear, _beginMonth, _beginDay) {
			diffDay = 0
			diffHour = 0
			diffMinute = 0
		}
	} else { //跨日
		_diffHour := 0
		_diffMinute := 0
		if beginTime < 1200 { //有跨午休
			_diffHour, _diffMinute = calcTime(beginHour, "18", beginMinute, "0", true)
		} else { //沒跨午休
			_diffHour, _diffMinute = calcTime(beginHour, "18", beginMinute, "0", false)
		}
		diffDay = calcDate(beginYear, beginMonth, beginDay, endYear, endMonth, endDay)
		if endTime < 1200 { //有跨午休
			diffHour, diffMinute = calcTime("8", endHour, "30", endMinute, true)
		} else { //沒跨午休
			diffHour, diffMinute = calcTime("8", endHour, "30", endMinute, false)
		}
		diffHour = diffHour + _diffHour
		diffMinute = diffMinute + _diffMinute
	}
	diffDay, diffHour, diffMinute = setCarry(diffDay, diffHour, diffMinute)
	c.JSON(http.StatusOK, gin.H{"diffDay": diffDay, "diffHour": diffHour, "diffMinute": diffMinute})
}

//setCarry 設定進位
func setCarry(diffDay int, diffHour int, diffMinute int) (_diffDay int, _diffHour int, _diffMinute int) {
	_diffDay = diffDay
	_diffHour = diffHour
	_diffMinute = diffMinute
	_diffHour = diffMinute/60 + _diffHour
	_diffMinute = diffMinute % 60
	_diffDay = _diffHour/8 + _diffDay
	_diffHour = _diffHour % 8
	return _diffDay, _diffHour, _diffMinute
}

//getHoliday 取得休假日期
func getHoliday() map[string]int {
	cfg := env.GetEnv()
	list := models.GetHoliday()
	data := make(map[string]int)
	for _, item := range list {
		datetime := item.Date.Format(cfg.TimeFormat)
		date := strings.Split(datetime, " ")[0]
		key := date
		data[key] = 1
	}
	return data
}

//appendZero 給不足二位數的數字前面加零
func appendZero(value string) string {
	if 2 > utf8.RuneCountInString(value) {
		value = "0" + value
	}
	return value
}

//isHoliday 判斷是否為假期
func isHoliday(year string, month string, day string) bool {
	list := getHoliday()
	checkTime := year + "-" + month + "-" + day
	_, ok := list[checkTime]
	if ok {
		return true
	}
	return false
}

//calcDate 計算跨日天數
func calcDate(beginYear string, beginMonth string, beginDay string, endYear string, endMonth string, endDay string) int {
	cfg := env.GetEnv()
	diffDay := 0
	_beginMonth := appendZero(beginMonth)
	_beginDay := appendZero(beginDay)
	_endMonth := appendZero(endMonth)
	_endDay := appendZero(endDay)
	checkTime := beginYear + "-" + _beginMonth + "-" + _beginDay + " 00:00:00"
	begin, _ := time.ParseInLocation(cfg.TimeFormat, checkTime, time.Local)
	checkTime = endYear + "-" + _endMonth + "-" + _endDay + " 00:00:00"
	end, _ := time.ParseInLocation(cfg.TimeFormat, checkTime, time.Local)
	if begin.Before(end) {
		for {
			now := begin.Format(cfg.TimeFormat)
			nowDate := strings.Split(now, " ")[0]
			nowArr := strings.Split(nowDate, "-")
			if isHoliday(nowArr[0], nowArr[1], nowArr[2]) {
				add, _ := time.ParseDuration("24h")
				begin = begin.Add(add)
				continue
			}
			if !begin.Before(end) {
				break
			}
			add, _ := time.ParseDuration("24h")
			begin = begin.Add(add)
			diffDay++
		}
	}

	return diffDay
}

//calcTimeBlock 計算時間區間，並扣除午休
func calcTime(beginHour string, endHour string, beginMinute string, endMinute string, haveLunchTime bool) (diffHour int, diffMinute int) {
	_beginHour, _ := strconv.Atoi(beginHour)
	_endHour, _ := strconv.Atoi(endHour)
	_beginMinute, _ := strconv.Atoi(beginMinute)
	_endMinute, _ := strconv.Atoi(endMinute)
	diffHour = _endHour - _beginHour
	diffMinute = diffHour*60 + _endMinute - _beginMinute
	if haveLunchTime {
		diffMinute = diffMinute - 90
	}
	diffHour = diffMinute / 60
	diffMinute = diffMinute % 60
	return diffHour, diffMinute
}

//changeToLunchTime 當請假起迄日為午休區間時，置換為午休結束
func changeToLunchTime(beginTime int, endTime int, beginHour string, beginMinute string, endHour string, endMinute string) (string, string, string, string) {
	if beginTime >= 1200 && beginTime <= 1330 { //當啟始時間於午休區間時
		beginHour = "13"
		beginMinute = "30"
	}
	if endTime >= 1200 && endTime <= 1330 { //當結束時間於午休區間時
		endHour = "13"
		endMinute = "30"
	}
	return beginHour, beginMinute, endHour, endMinute
}

//isSameDay 判斷二個時間區間是否為同一天
func isSameDay(pc *beans.Punchclock) bool {
	begin := pc.GetBegin().GetYear() + pc.GetBegin().GetMonth() + pc.GetBegin().GetDay()
	end := pc.GetEnd().GetYear() + pc.GetEnd().GetMonth() + pc.GetEnd().GetDay()
	return begin == end
}

//getPunchclockParams 取得HTTP POST帶過來之參數
func getPunchclockParams(c *gin.Context) *beans.Punchclock {
	params := &beans.Punchclock{}
	err := c.BindJSON(params)
	if err != nil {
		log.Println(err)
	}
	return params
}
