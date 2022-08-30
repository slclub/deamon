package rpcx

import (
	"github.com/slclub/deamon/utils"
	rpcxServer "github.com/smallnest/rpcx/server"
)

var rpcxServerObj = rpcxServer.NewServer()

func Start(addr string) {
	go rpcxServerObj.Serve("tcp", addr)
}

/**
 * 注册自定义的 对外类，类下挂载着方法，从rpcx的client端旧可以调用类
 */
func ListenRegister(controller string, obj interface{}) {
	if utils.IsNil(obj) {
		panic(any("[RPCX][REGISTER][PANIC]" + controller))
	}
	rpcxServerObj.RegisterName(controller, obj, "")
}
