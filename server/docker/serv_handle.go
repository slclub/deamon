package docker

import (
	"context"
	"github.com/slclub/deamon/helper"
	"github.com/slclub/deamon/helper/yamlz"
	"github.com/slclub/deamon/logger"
	"github.com/slclub/deamon/msg"
	"github.com/slclub/deamon/server/model"
	"github.com/slclub/deamon/server/syscmd"
)

type DockerServer struct {
}

// 获取某个服务数据返回
func (this *DockerServer) List(ctx context.Context, req *msg.Request, res *msg.Response) error {
	res.MsgCode = helper.SUCCESS
	proc := model.Mgr().GetByName(req.Name)
	if proc == nil {
		return nil
	}
	//if !model.Mgr().Check(proc.Name) {
	//	return nil
	//}
	res.Data = map[string]any{
		"Name":  req.Name,
		"State": proc.State,
		"Addr":  "",
	}
	return nil
}

// 临时 状态改变
func (this *DockerServer) Pause(ctx context.Context, req *msg.Request, res *msg.Response) error {
	res.MsgCode = helper.SUCCESS
	proc := model.Mgr().GetByName(req.Name)
	if proc == nil {
		return nil
	}
	if req.Data == nil {
		logger.Printf("RPCX.SERV.PAUSE data empty!")
		return nil
	}
	state_any, ok := req.Data["State"]

	if !ok {
		return nil
	}

	val, ok := state_any.(string)
	if !ok {
		return nil
	}
	if val == helper.DEAMON_ENABLE || val == helper.DEAMON_DISABLE {
		proc.State = val
	}
	return nil
}

// start
func (this *DockerServer) Start(ctx context.Context, req *msg.Request, res *msg.Response) error {
	res.MsgCode = helper.SUCCESS
	proc := model.Mgr().GetByName(req.Name)
	if proc == nil {
		return nil
	}
	docker_item := yamlz.GetDockerItem(req.Name)
	is_bash := docker_item.Mode == helper.OS_SCRIPT_BASH
	data, err := syscmd.BashCmd(docker_item.Start, is_bash)
	if err != nil {
		logger.Printf("RPCX.SERVER.Start FAIL ID:%v, Name:%v  start-output:%v err:%v", docker_item.ID, docker_item.Name, string(data), err)
		return err
	}
	logger.Printf("RPCX.SERVER.Start OK  ID:%v, Name:%v start-output:%v", docker_item.ID, docker_item.Name, string(data))
	return nil
}

// stop
func (this *DockerServer) Stop(ctx context.Context, req *msg.Request, res *msg.Response) error {
	res.MsgCode = helper.SUCCESS
	proc := model.Mgr().GetByName(req.Name)
	if proc == nil {
		return nil
	}
	docker_item := yamlz.GetDockerItem(req.Name)
	is_bash := docker_item.Mode == helper.OS_SCRIPT_BASH
	data, err := syscmd.BashCmd(docker_item.Stop, is_bash)
	if err != nil {
		logger.Printf("RPCX.SERVER.Stop FAIL ID:%v, Name:%v  stop-output:%v", docker_item.ID, docker_item.Name, string(data))
		return err
	}
	logger.Printf("RPCX.SERVER.Stop OK  ID:%v, Name:%v stop-output:%v", docker_item.ID, docker_item.Name, string(data))
	return nil
}

// restart
func (this *DockerServer) Restart(ctx context.Context, req *msg.Request, res *msg.Response) error {
	err := this.Stop(ctx, req, res)
	if err != nil {
		return err
	}
	return this.Start(ctx, req, res)
}
