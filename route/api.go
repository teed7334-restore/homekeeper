package route

import (
	"os"

	"github.com/teed7334-restore/homekeeper/controllers"

	"github.com/gin-gonic/gin"
)

//API Restful路由
func API() *gin.Engine {
	env := os.Getenv("env")
	gin.SetMode(env)
	route := gin.Default()
	route.POST("/Mail/SendMail", controllers.SendMail)
	route.POST("/Redis/Get", controllers.GetRedis)
	route.POST("/Redis/Set", controllers.SetRedis)
	route.POST("/Redis/Incr", controllers.IncrRedis)
	route.POST("/Redis/Decr", controllers.DecrRedis)
	route.POST("/Redis/HSet", controllers.HSetRedis)
	route.POST("/Redis/HGet", controllers.HGetRedis)
	route.POST("/Redis/SAdd", controllers.SAddRedis)
	route.POST("/Redis/SCard", controllers.SCardRedis)
	route.POST("/Redis/LPush", controllers.LPushRedis)
	route.POST("/Redis/RPush", controllers.RPushRedis)
	route.POST("/Redis/LSet", controllers.LSetRedis)
	route.POST("/Redis/LRange", controllers.LRangeRedis)
	route.POST("/PunchClock/CalcTime", controllers.CalcTime)
	route.POST("/PunchClock/ResetAllUseMinute", controllers.ResetAllUseMinute)
	return route
}
