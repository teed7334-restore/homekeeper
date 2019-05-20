package main

import (
	"github.com/teed7334-restore/homekeeper/env"
	"github.com/teed7334-restore/homekeeper/route"
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
