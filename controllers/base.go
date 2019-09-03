package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teed7334-restore/homekeeper/beans"
	"github.com/teed7334-restore/homekeeper/env"
)

//ChainResultObject Hyperledger REST回傳物件
type ChainResultObject interface {
	GetError() *beans.APIError
}

var cfg = env.GetEnv()

//GetURL 透過HTTP GET取得網頁資料
func GetURL(url string) []byte {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panicln(err)
	}
	client := &http.Client{}
	result, err := client.Do(request)
	if err != nil {
		log.Panicln(err)
	}
	body, _ := ioutil.ReadAll(result.Body)
	defer result.Body.Close()
	return body
}

//PostURL 透過HTTP POST扔資料給特定網頁
func PostURL(url string, params []byte) []byte {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(params))
	request.Header.Set("X-Custom-Header", "counter")
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	result, err := client.Do(request)
	if err != nil {
		log.Panicln(err)
	}
	body, _ := ioutil.ReadAll(result.Body)
	defer result.Body.Close()
	return body
}

//getParams 取得HTTP POST帶過來之參數
func getParams(c *gin.Context, params interface{}) {
	err := c.BindJSON(params)
	if err != nil {
		log.Panicln(err)
	}
}

//getChainParams 取得Hyperledger鏈上傳來的資料
func getChainParams(url string, params []byte, action string, resultObject ChainResultObject) {
	result := []byte{}
	switch action {
	case "GET":
		result = GetURL(url)
	case "POST":
		result = PostURL(url, params)
	}
	err := json.Unmarshal(result, resultObject)
	if err != nil {
		log.Panicln(err)
		return
	}
	if resultObject.GetError() != nil {
		log.Panicln(resultObject.GetError().GetMessage())
		return
	}
}
