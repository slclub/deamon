package utils

import (
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var APP_PATH string

func init() {
	APP_PATH = GetRootPath(false)
}

// 获取可执行文件的绝对根路径
func GetRootPath(has_bin bool) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}

	if has_bin {
		dir = filepath.Dir(dir)
	}
	// For testing command.
	if dir[:4] == "/tmp" {
		dir, err = os.Getwd()
	}
	return dir
}

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Println(localAddr.String(), "===")
	ip = strings.Split(localAddr.String(), ":")[0]
	//fmt.Println("ip:", ip)
	return
}

func GetCurrentFilePath() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(any(errors.New("Can not get current file path info")))
	}
	return file
}
