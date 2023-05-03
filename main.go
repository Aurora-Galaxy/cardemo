package main

import (
	"car/conf"
	"car/router"
	"car/pkg/util"
)

func main() {
	conf.Init()
	go util.CheckPile()
	r := router.NewRouter()
	r.Run(conf.HttpPort)
}
