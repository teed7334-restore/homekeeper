package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teed7334-restore/homekeeper/base"
	"github.com/teed7334-restore/homekeeper/beans"
	"github.com/teed7334-restore/homekeeper/env"
)

//PayCreditCard 使用信用卡付款
func PayCreditCard(c *gin.Context) {
	cfg := env.GetEnv()
	params := getCreditCardParams(c)
	params.CustomerId = cfg.GoMyPay.CustomerID
	jsonStr, _ := json.Marshal(params)
	jsonByte := []byte(jsonStr)
	url := cfg.GoMyPay.URL
	jsonByte = base.PostURL(url, jsonByte)
	c.JSON(http.StatusOK, jsonByte)
}

//getCreditCardParams 取得HTTP POST帶過來之參數
func getCreditCardParams(c *gin.Context) *beans.CreditCard {
	params := &beans.CreditCard{}
	err := c.BindJSON(params)
	if err != nil {
		log.Println(err)
	}
	return params
}
