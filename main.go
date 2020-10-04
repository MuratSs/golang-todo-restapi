package main

import (
	"github.com/MuratSs/golang-todo-restap/app"
	"github.com/MuratSs/golang-todo-restap/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
