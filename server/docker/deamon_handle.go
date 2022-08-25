package docker

import (
	"github.com/slclub/deamon/helper"
	"github.com/slclub/deamon/helper/yamlz"
	"github.com/slclub/deamon/logger"
	"github.com/slclub/deamon/server/model"
	"github.com/slclub/deamon/server/syscmd"
	"strings"
)

type Request struct {
	ServName string
	ServID   int
}

type Response struct {
	MsgCode int
	Desc    string
}

// 定时100 S 更新一次yaml watch
// 1000 帧 执行一次
func TickReloadYamlHandle() func() {
	frame, step := 1000, 1000
	model.Mgr().Add(&model.DeamonProcceser{Name: helper.DEAMON_CONFIG_SERV, State: helper.DEAMON_ENABLE})
	return func() {
		step--
		if step > 0 {
			return
		}

		step = frame
		if !model.Mgr().Check(helper.DEAMON_CONFIG_SERV) {
			return
		}
		yamlz.Reload()

	}
}

// 依据 docker 监听 多个服务
func WatchServerHandle() []func() {
	dockers := yamlz.Server.Commands.Docker

	fns := make([]func(), 0)
	for i, _ := range dockers {
		fn := explainServerCommand(&dockers[i])

		if fn == nil {
			continue
		}
		fns = append(fns, fn)
	}
	return fns
}

func explainServerCommand(docker_item *yamlz.DockerItem) func() {
	open := docker_item.Open

	if !open {
		return nil
	}

	model.Mgr().Add(&model.DeamonProcceser{ID: docker_item.ID, Name: docker_item.Name, State: helper.DEAMON_ENABLE})
	frame := 20 // 20  帧 2秒
	step := frame

	return func() {
		step--
		if step > 0 {
			return
		}
		step = frame

		if !model.Mgr().Check(docker_item.Name) {
			return
		}

		is_bash := docker_item.Mode == helper.OS_SCRIPT_BASH
		// 服务检查
		check_info, _ := syscmd.BashCmd(docker_item.CheckCmd, is_bash)
		server_ok := strings.Contains(string(check_info), docker_item.CheckValue)
		if server_ok {
			return
		}

		// 重启服务之前，输出 待守护进程的错误信息
		data, _ := syscmd.BashCmd(docker_item.RestartBefore, is_bash)
		logger.Printf("DEAMON.SERVER.RESTART BEFORE ID:%v, Name:%v CMD:%v stdout:%v", docker_item.ID, docker_item.Name, docker_item.RestartBefore, string(data))

		// 重启 监听的服务
		data, err := syscmd.BashCmd(docker_item.Restart, is_bash)
		if err == nil {
			logger.Printf("DEAMON.SERVER.RESTART OK  ID:%v, Name:%v restart-output:%v", docker_item.ID, docker_item.Name, string(data))
			return
		}
		logger.Printf("DEAMON.SERVER.RESTART FAIL ID:%v, Name:%v  restart-output:%v", docker_item.ID, docker_item.Name, string(data))
	}
}
