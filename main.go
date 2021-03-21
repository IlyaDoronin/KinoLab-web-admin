package main

import (
	"github.com/yeahyeahcore/KinoLab-web-admin/conf"
	"github.com/yeahyeahcore/KinoLab-web-admin/server"
)

func main() {
	conf.Load("conf/config.json")
	server.Start()
}
