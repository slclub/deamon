package servsys

import (
	"context"
	"github.com/slclub/deamon/helper"
	"github.com/slclub/deamon/logger"
	"github.com/slclub/deamon/msg"
	"github.com/slclub/deamon/plugin/rpcx"
)

// example : DockerServer.List
func GetServListByName(cc *rpcx.RpcxClient, serv_name string, controllers ...string) []*msg.ServItem {
	controller := "DockerServer"
	if controllers != nil && len(controllers) > 0 {
		controller = controllers[0]
	}
	clients := cc.GetClientByController(controller)
	vadata := make([]*msg.ServItem, 0)
	for i, cli := range clients {
		res := msg.Response{}
		err := cli.Call(context.Background(), "List", &msg.Request{Name: serv_name}, &res)
		if err != nil || res.MsgCode != helper.SUCCESS {
			logger.Printf("SERVICE.GetServListByName controller:%v; code:%v err:%v", serv_name, res.MsgCode, err)
			continue
		}
		dm, ok := res.Data.(map[string]any)
		if !ok {
			continue
		}
		vadata = append(vadata, &msg.ServItem{
			Index: i,
			Name:  dm["Name"].(string),
			State: dm["State"].(string),
			Addr:  rpcx.ServiceAddrs[cli.GetIndex()],
		})
	}
	return vadata
}

// start
func StartWithID(cc *rpcx.RpcxClient, serv_name string, id int) {
	controller := "DockerServer"
	clients := cc.GetClientByController(controller)
	if id > len(clients) {
		return
	}
	for _, cli := range clients {
		if cli.GetIndex() == id {
			StartWithName(cc, serv_name, cli.GetAddr())
		}
	}
}

func StartWithName(cc *rpcx.RpcxClient, serv_name, addr string) {
	controller := "DockerServer"
	client := cc.GetByAddr(controller, addr)
	if client == nil {
		return
	}
	res := msg.Response{}
	err := client.Call(context.Background(), "Start", &msg.Request{Name: serv_name}, &res)
	logger.Printf("SERVICE.START name:%v code:%v err:%v", serv_name, res.MsgCode, err)
}

func StopWithName(cc *rpcx.RpcxClient, serv_name string, addr string) {
	controller := "DockerServer"
	client := cc.GetByAddr(controller, addr)
	if client == nil {
		return
	}
	res := msg.Response{}
	err := client.Call(context.Background(), "Stop", &msg.Request{Name: serv_name}, &res)
	logger.Printf("SERVICE.STOP name:%v code:%v err:%v", serv_name, res.MsgCode, err)
}

func RestartWithName(cc *rpcx.RpcxClient, serv_name, addr string) {
	controller := "DockerServer"
	client := cc.GetByAddr(controller, addr)
	if client == nil {
		return
	}
	res := msg.Response{}
	err := client.Call(context.Background(), "Restart", &msg.Request{Name: serv_name}, &res)
	logger.Printf("SERVICE.RESTART name:%v code:%v err:%v", serv_name, res.MsgCode, err)
}
