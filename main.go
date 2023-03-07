package main

import (
	"car/conf"
	"car/router"
)

func main() {
	conf.Init()
	r := router.NewRouter()
	r.Run(conf.HttpPort)
}
