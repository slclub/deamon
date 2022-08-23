package syscmd

import (
	"github.com/slclub/deamon/logger"
	"testing"
)

func TestBashCmd(t *testing.T) {
	initStart()
	out := BashCmd("ps -ef", true)
	logger.Log().Printf("std :%v", string(out))
}

func initStart() {
	lg := logger.New()
	logger.Log(lg)
}
