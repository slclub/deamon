package server

import (
	"fmt"
	"github.com/slclub/deamon/logger"
	"github.com/slclub/deamon/utils"
	"path/filepath"
	"testing"
)

func TestServer(t *testing.T) {
	initStart()
	serv := New()
	serv.Register(func() {
		logger.Printf("SERVER.REGISTER.FUNC.EXEC")
	})
	serv.Init()

	serv.Start()
}

func initStart() {
	p := utils.GetCurrentFilePath()
	utils.APP_PATH = filepath.Dir(p)
	utils.APP_PATH = filepath.Dir(utils.APP_PATH)
	fmt.Println("APP_PATH:", utils.APP_PATH)
}
