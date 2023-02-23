package main

import (
	"go_mall/conf"
	"go_mall/routes"
)

// main
// @Description: 开启所有服务
func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
