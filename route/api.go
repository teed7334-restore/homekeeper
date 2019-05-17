package route

import (
	"../controllers"
	"../env"
	"github.com/gin-gonic/gin"
)

var cfg = env.GetEnv()

//API Restful路由
func API() *gin.Engine {
	gin.SetMode(cfg.Env)
	route := gin.Default()
	route.POST("/Mail/SendMail", controllers.SendMail)
	return route
}
