package main

import (
	"EIM/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	// 启动服务端
	var server = new(cmd.Server)
	server.RunConfig()
	logrus.Info(server.Config)
	server.RunDataBase()
}
