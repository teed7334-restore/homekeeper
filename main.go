package main

import (
	"./env"
	"./route"
)

var cfg = env.GetEnv()

func main() {
	webService()
}

//webService Restful API服務
func webService() {
	api := route.API()
	api.Run(":8806")
}