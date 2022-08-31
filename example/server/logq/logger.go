package logq

import (
	"github.com/slclub/deamon/logger"
	"github.com/slclub/deamon/utils"
	"github.com/slclub/glog"
	"server/helper"
)

/**
 * example:
 * 		log.Info("[HANDLE_INFO][ECHO]", "show info done!")
 * 		log.InfoF("[HANDLE][INFO]%s", "That is fun joke")
 *		log.Debug(...)
 * 		log.DebugF(format, ...string)
 */
func InitLog() {
	configLog := utils.Config("server").GetStringMapString("Log")
	log_path := configLog["log_abs_path"]

	if configLog["log_abs_path"] == "" {
		log_path = utils.APP_PATH
	}
	glog.Set("path", log_path, configLog["log_rel_path"])
	glog.Set("name", configLog["log_name"])
	glog.Set("head", configLog["log_header"])
	glog.Set("stderr", configLog["stderr"])
	if helper.ConvAnyToInt(configLog["log_level"]) > 0 {
		glog.Set("debug", configLog["log_level"])
	}
	glog.Set("debug", configLog["debug"] == "true")
	//glog.Set("capacity", 3 * 1024 * 1024)// 设置缓冲区为3M // 现在goqueue 使用的2M 缓存，

}

// 如果想打印日志时带上文件名和行数，这个值需要对应上调度层级，才可以正确的打印我们代码文件
// warn 级别的错误自动带入文件名和行数。可以快速的定位日志代码位置
var log_depath = 1

// info
func Info(args ...interface{}) {
	glog.InfoDepth(log_depath, args...)
}
func InfoF(format string, args ...interface{}) {
	glog.Infof(format, args...)
}

// debug
func Debug(args ...interface{}) {
	glog.DebugDepth(log_depath, args...)
}
func DebugF(format string, args ...interface{}) {
	glog.Debugf(format, args...)
}

// warnning
func Warn(args ...interface{}) {
	glog.WarnningDepth(log_depath, args...)
}
func WarnF(format string, args ...interface{}) {
	glog.Warnningf(format, args...)
}

//error
func Error(args ...interface{}) {
	glog.ErrorDepth(log_depath, args...)
}
func ErrorF(format string, args ...interface{}) {
	glog.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	glog.FatalDepth(log_depath, args...)
}

type Log struct {
}

func New() logger.Logger {
	return &Log{}
}

func (this *Log) Printf(format string, args ...any) {
	glog.Infof(format, args...)
}

func (this *Log) Print(args ...any) {
	glog.Info(args...)
}

var _log logger.Logger = &Log{}
