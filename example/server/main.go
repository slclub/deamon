package main

import (
	"github.com/slclub/deamon/helper"
	"github.com/slclub/deamon/server"
	"server/logq"
	"server/option"
)

func main() {
	logq.InitLog()

	serv := server.New()
	serv.Register(helper.OptionRegisterFunc(func() {
		server.RegisterRoute(serv)
	}))
	serv.Init(option.InitMyLog())
	//logger.Log().Print(utils.Config("server").GetStringMapString("Log"))

	serv.Start()
	select {}
}
