package main

import (
	"clash_config/config"
	"clash_config/log"
	"clash_config/method"
	"clash_config/router"
)

func main() {
	log.LogInit()
	config.ConfigInit()
	method.Flags()
	router.Router()
}
