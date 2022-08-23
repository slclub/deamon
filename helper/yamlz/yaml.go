package yamlz

import (
	"github.com/slclub/deamon/logger"
	"github.com/slclub/deamon/utils"
	"io/ioutil"
)

func ReadYamlFile(file_name string, app_paths ...string) []byte {
	app_path := utils.APP_PATH
	if len(app_paths) > 0 {
		app_path = app_paths[0]
	}
	if file_name == "" {
		logger.Printf("YAMLZ.READ.FILE fail path:%v err:%v", app_path+file_name, "empty yaml file")
		return nil
	}
	if file_name[0] != '/' {
		file_name = "/" + file_name
	}

	data, err := ioutil.ReadFile(app_path + file_name)
	if err != nil {
		logger.Printf("YAMLZ.READ.FILE fail path:%v err:%v", app_path+file_name, err)
		return nil
	}
	return data
}

func Reload() {
	_load_server()
	// _load_client()
}

func Init() {
	_load_server()
}
