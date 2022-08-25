package syscmd

import (
	"bytes"
	"github.com/slclub/deamon/logger"
	"os/exec"
	"strings"
)

// 执行 bash 脚本
func BashCmd(cmd string, shell bool) (string, error) {
	if shell {
		return RunCmd("/bin/bash", "-c", cmd)
	}
	cmds := strings.Split(cmd, " ")
	if len(cmds) == 1 {
		return RunCmd(cmd)
	}
	return RunCmd(cmds[0], cmds[1:]...)
}

func RunCmd(name string, args ...string) (string, error) {
	//cmd := exec.Command("/bin/bash", "-c", cmd_str)
	cmd := exec.Command(name, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		args = append([]string{name}, args...)
		logger.Log().Printf("SERVER.COMMAND.RUN.CMD-FAIL  CMD:%v  ERROR: %v ; %v", strings.Join(args, " "), err, stderr.String())
		//return &out
	}
	return out.String(), err
}
