package main

import (
	"github.com/yuhando/simpleapi/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":9000")
}
