package helper

import (
	"github.com/spf13/viper"
	"strings"
)

// Option 模式
type Option struct {
	addr     string
	register OptionRegisterFunc
	config   *viper.Viper
}

type OptionFunc func(opt *Option)
type OptionRegisterFunc func() // 注册函数类型

// 默认获取，使用addr
func OptionRpcxAddr(addr string) OptionFunc {
	return func(opt *Option) {
		opt.addr = addr
	}
}

func OptionRegister(fn OptionRegisterFunc) OptionFunc {
	return func(opt *Option) {
		opt.register = fn
	}
}

func OptionConfig(vp *viper.Viper) OptionFunc {
	return func(opt *Option) {
		opt.config = vp
	}
}

func OptionGet(opt *Option, field string) any {
	field = strings.ToLower(field)
	switch field {
	case "addr":
		return opt.addr
	case "register":
		return opt.register
	case "config":
		return opt.config
	}
	return nil
}
