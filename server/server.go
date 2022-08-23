package server

import (
	"github.com/slclub/deamon/helper"
	"github.com/slclub/deamon/helper/yamlz"
	"github.com/slclub/deamon/logger"
	"github.com/slclub/deamon/server/docker"
	"github.com/slclub/deamon/utils"
	rpcxServer "github.com/smallnest/rpcx/server"
	"github.com/spf13/viper"
	"time"
)

type Server struct {
	option         *helper.Option
	rpcx           *rpcxServer.Server
	frame_duration time.Duration // ms
	tick_handles   []func()
}

/**
 * server.Init()
 * server.Start()
 */
func New() *Server {

	//logq.InitLog()
	return &Server{
		option:         &helper.Option{},
		frame_duration: helper.FRAME_DEFAULT,
		rpcx:           rpcxServer.NewServer(),
	}
}

func (this *Server) Init(opt_funcs ...helper.OptionFunc) {
	config := this.getConfig()
	if config == nil {
		helper.OptionConfig(utils.Config("server"))(this.option)
		config = this.getConfig()
	}
	opf := helper.OptionRpcxAddr(config.GetStringMapString("SrvRpcx")["addr"])
	opf(this.option)

	// 设置默认帧
	frame := config.GetInt("SrvRpcx.Frame")
	if frame > 0 {
		this.frame_duration = time.Duration(frame)
	}

	logger.Log(logger.New())

	if len(opt_funcs) > 0 {
		for _, fn := range opt_funcs {
			if fn == nil {
				continue
			}
			fn(this.option)
		}
	}

	this.rpcx = rpcxServer.NewServer()
	yamlz.Init()

	this.addServrsDefault()
	logger.Printf("DEAMON.SERVER.INIT.OK")
}

func (this *Server) Start() {
	addr, _ := helper.OptionGet(this.option, "addr").(string)

	// 执行注入的 注册函数
	if reg, ok := helper.OptionGet(this.option, "register").(helper.OptionRegisterFunc); ok && reg != nil {
		reg()
	}
	this.Tick()
	logger.Printf("LISTEN.RPC.ADDR:%v", addr)
	this.rpcx.Serve("tcp", addr)
}

// 注册一个空函数
func (this *Server) Register(reg helper.OptionRegisterFunc) {
	// 默认加入 tick handle
	this.AddTickHandle(docker.TickReloadYamlHandle())
	if reg != nil {
		helper.OptionRegister(reg)(this.option)
	}
}

func (this *Server) AddServ(controller string, rcvr interface{}) {
	meta_any := helper.OptionGet(this.option, "metadata")
	metadata := ""
	if metad, ok := meta_any.(string); ok && meta_any != nil {
		metadata = metad
	}
	this.rpcx.RegisterName(controller, rcvr, metadata)
}

func (this *Server) AddTickHandle(tick_fn func()) {
	this.tick_handles = append(this.tick_handles, tick_fn)
}

// 固定监听服务消息
func (this *Server) Tick() {
	go func() {
		ticker := time.NewTicker(this.frame_duration * time.Millisecond)
		defer ticker.Stop()
		for {

			<-ticker.C
			for _, tick_fn := range this.tick_handles {
				if tick_fn == nil {
					continue
				}
				tick_fn()
			}

		}
		logger.Printf("DEAMON.SERVER.TICK.QUIT")
	}()
}

func (this *Server) addServrsDefault() {
	srvs := docker.WatchServerHandle()
	for _, fn := range srvs {
		this.AddTickHandle(fn)
	}
}

func (this *Server) getConfig() *viper.Viper {
	val_any := helper.OptionGet(this.option, "config")
	val, ok := val_any.(*viper.Viper)
	if !ok {
		return nil
	}
	if val == nil {
		return nil
	}
	return val
}
