package controllers

import (
	"log"
	"net/http"
	"strconv"
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
			diffHour, diffMinute = calcTimeBlock(beginHour, endHour, beginMinute, endMinute, true)
		} else { //沒跨午休
			diffHour, diffMinute = calcTimeBlock(beginHour, endHour, beginMinute, endMinute, false)
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
			_diffHour, _diffMinute = calcTimeBlock(beginHour, "18", beginMinute, "0", true)
		} else { //沒跨午休
			_diffHour, _diffMinute = calcTimeBlock(beginHour, "18", beginMinute, "0", false)
		}
		diffDay = calcDate(beginYear, beginMonth, beginDay, endYear, endMonth, endDay)
		if endTime < 1200 { //有跨午休
			diffHour, diffMinute = calcTimeBlock("8", endHour, "30", endMinute, true)
		} else { //沒跨午休
			diffHour, diffMinute = calcTimeBlock("8", endHour, "30", endMinute, false)
		}
		diffHour = diffHour + _diffHour
		diffMinute = diffMinute + _diffMinute
	}
	c.JSON(http.StatusOK, gin.H{"diffDay": diffDay, "diffHour": diffHour, "diffMinute": diffMinute})
}

func getHoliday() map[string]int {
	cfg := env.GetEnv()
	list := models.GetHoliday()
	data := make(map[string]int)
	for _, item := range list {
		key := item.Date.Format(cfg.TimeFormat)
		data[key] = 1
	}
	return data
}

func appendZero(value string) string {
	if 2 > utf8.RuneCountInString(value) {
		value = "0" + value
	}
	return value
}

func isHoliday(year string, month string, day string) bool {
	cfg := env.GetEnv()
	list := getHoliday()
	checkTime := year + "-" + month + "-" + day + " 00:00:00"
	begin, _ := time.ParseInLocation(cfg.TimeFormat, checkTime, time.Local)
	_, ok := list[begin.Format(cfg.TimeFormat)]
	if ok {
		return true
	}
	return false
}

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
			if isHoliday(beginYear, _beginMonth, _beginDay) {
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

func calcTimeBlock(beginHour string, endHour string, beginMinute string, endMinute string, haveLunchTime bool) (diffHour int, diffMinute int) {
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

func isSameDay(pc *beans.Punchclock) bool {
	begin := pc.GetBegin().GetYear() + pc.GetBegin().GetMonth() + pc.GetBegin().GetDay()
	end := pc.GetEnd().GetYear() + pc.GetEnd().GetMonth() + pc.GetEnd().GetDay()
	return begin == end
}

func getPunchclockParams(c *gin.Context) *beans.Punchclock {
	params := &beans.Punchclock{}
	err := c.BindJSON(params)
	if err != nil {
		log.Println(err)
	}
	return params
}
