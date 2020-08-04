package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/teed7334-restore/homekeeper/route"
)

func main() {
	webService()
}

//webService Restful API服務
func webService() {
	api := route.API()
	api.Run(":8806")
}
