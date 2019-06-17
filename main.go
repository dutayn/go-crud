package main

import (
	"go-crud/app"
	"go-crud/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")

}
