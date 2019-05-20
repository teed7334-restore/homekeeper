package controllers

import (
	"log"
	"net/http"
	"strings"

	"homekeeper/beans"

	"homekeeper/env"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

//GetRedis 取得Redis資料
func GetRedis(c *gin.Context) {
	params := getParams(c)
	client := initRedis()
	value, err := redis.String(client.Do("get", params.GetKey()))
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()
	c.JSON(http.StatusOK, gin.H{"value": value})
}

//SetRedis 設定Redis資料
func SetRedis(c *gin.Context) {
	params := getParams(c)
	client := initRedis()
	value, err := client.Do("set", params.GetKey(), params.GetValue())
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()
	c.JSON(http.StatusOK, gin.H{"success": value})
}

//IncrRedis 對Redis資料進行遞增
func IncrRedis(c *gin.Context) {
	params := getParams(c)
	client := initRedis()
	value, err := client.Do("incr", params.GetKey())
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()
	c.JSON(http.StatusOK, gin.H{"value": value})
}

//DecrRedis 對Redis資料進行遞減
func DecrRedis(c *gin.Context) {
	params := getParams(c)
	client := initRedis()
	value, err := client.Do("decr", params.GetKey())
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()
	c.JSON(http.StatusOK, gin.H{"value": value})
}

//HSetRedis 對Redis資料建立Hashmap
func HSetRedis(c *gin.Context) {
	params := getParams(c)
	client := initRedis()
	value, err := client.Do("hset", params.GetKey(), params.GetHkey(), params.GetValue())
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()
	c.JSON(http.StatusOK, gin.H{"value": value})
}

//HGetRedis 取得建立Hashmap的Redis資料
func HGetRedis(c *gin.Context) {
	params := getParams(c)
	client := initRedis()
	value, err := redis.String(client.Do("hget", params.GetKey(), params.GetHkey()))
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()
	c.JSON(http.StatusOK, gin.H{"value": value})
}

//SAddRedis 對Redis進行資料添加，並排除重複項目
func SAddRedis(c *gin.Context) {
	params := getParams(c)
	client := initRedis()
	value, err := client.Do("sadd", params.GetKey(), params.GetValue())
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()
	c.JSON(http.StatusOK, gin.H{"value": value})
}

//SCardRedis 取得Key中的資料集合總數
func SCardRedis(c *gin.Context) {
	params := getParams(c)
	client := initRedis()
	value, err := client.Do("scard", params.GetKey())
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()
	c.JSON(http.StatusOK, gin.H{"value": value})
}

//LPushRedis 對Redis中List資料的前面做資料添加
func LPushRedis(c *gin.Context) {
	params := getParams(c)
	client := initRedis()
	value, err := client.Do("lpush", params.GetKey(), params.GetValue())
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()
	c.JSON(http.StatusOK, gin.H{"value": value})
}

//RPushRedis 對Redis中List資料的後面做資料添加
func RPushRedis(c *gin.Context) {
	params := getParams(c)
	client := initRedis()
	value, err := client.Do("rpush", params.GetKey(), params.GetValue())
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()
	c.JSON(http.StatusOK, gin.H{"value": value})
}

//LSetRedis 從Redis中List資料的第...列做資料修改
func LSetRedis(c *gin.Context) {
	params := getParams(c)
	client := initRedis()
	value, err := client.Do("lset", params.GetKey(), params.GetHkey(), params.GetValue())
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()
	c.JSON(http.StatusOK, gin.H{"value": value})
}

//LRangeRedis 取得Redis中List特定範圍裡面的資料
func LRangeRedis(c *gin.Context) {
	params := getParams(c)
	getRange := strings.Split(params.GetValue(), ":")
	begin := getRange[0]
	end := getRange[1]
	client := initRedis()
	value, err := redis.Strings(client.Do("lrange", params.GetKey(), begin, end))
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()
	c.JSON(http.StatusOK, gin.H{"value": value})
}

func initRedis() redis.Conn {
	cfg := env.GetEnv()
	client, err := redis.Dial(cfg.Redis.Protocol, cfg.Redis.Host)
	if err != nil {
		log.Panic(err)
	}
	return client
}

func getParams(c *gin.Context) *beans.Redis {
	params := &beans.Redis{}
	err := c.BindJSON(params)
	if err != nil {
		log.Println(err)
	}
	return params
}
