package main

import (
	"twb-otus-25/demoapp/src/app"
	"twb-otus-25/demoapp/src/config"
	"twb-otus-25/demoapp/src/log"
)

func main() {
    log := log.New()
    config := config.New(log)
	app.Run(log, config)
}