package option

import (
	"github.com/slclub/deamon/helper"
	"github.com/slclub/deamon/logger"
	"server/logq"
)

// 自定义日志
func InitMyLog() helper.OptionFunc {
	return func(option *helper.Option) {
		logger.Log(logq.New())
	}
}
