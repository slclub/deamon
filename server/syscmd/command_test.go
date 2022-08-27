package syscmd

import (
	"github.com/slclub/deamon/logger"
	"os/exec"
	"testing"
)

func TestBashCmd(t *testing.T) {
	initStart()
	out, err := BashCmd("ps aux | grep tiandi", true)
	logger.Log().Printf("std :%v ", string(out))
	if err != nil {
		t.Error(err)
	}
}

func TestDockerComm(t *testing.T) {
	initStart()
	out, err := RunCmd("bash", "-c", "cat ~/.vim/vimrc | grep self")
	logger.Log().Printf("std :%v ", string(out))
	if err != nil {
		t.Error(err)
	}
}

func TestExecCmdCombind(t *testing.T) {
	out, err := exec.Command("bash", "-c", "cat ~/.vim/vimrc | grep angel").CombinedOutput()
	logger.Log().Printf("exec-std :%v ", string(out))
	if err != nil {
		t.Error(err)
	}
}

func TestExecCmdMany(t *testing.T) {
	out, err := exec.Command("bash", "-c", "cat /proc/cpuinfo").Output()
	logger.Log().Printf("bash -c: %s", string(out))
	if err != nil {
		t.Error(err)
	}
}

func initStart() {
	lg := logger.New()
	logger.Log(lg)
}
