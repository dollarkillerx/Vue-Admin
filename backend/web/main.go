package main

import (
	"Vue-Admin/backend/config"
	"Vue-Admin/backend/web/router"
	"log"
)

func main() {
	app := router.RegisterRouter()

	log.Println("http://127.0.0.1:8087")
	app.Run(config.MyConfig.App.Host)
}
