package yamlz

import (
	"github.com/slclub/deamon/logger"
	"gopkg.in/yaml.v2"
)

// 对外开放的
var (
	Server *ServerYaml
)

type ServerYaml struct {
	SrvRpcx  ServerRpcxYaml `yaml:"SrvRpcx"`
	Commands CommandYaml    `yaml:"Commands"`
}

type ServerRpcxYaml struct {
	Addr  string `yaml:"addr"`
	Frame int    `yaml:"Frame"`
}

type CommandYaml struct {
	Docker []DockerItem `yaml:"docker"`
}

func newServerYaml() *ServerYaml {
	return &ServerYaml{
		SrvRpcx: ServerRpcxYaml{},
		Commands: CommandYaml{
			Docker: make([]DockerItem, 0),
		},
	}
}
func _load_server() {
	Server = newServerYaml()
	data := ReadYamlFile("conf/server.yaml")
	if data == nil {
		return
	}
	// 安全加载
	var srvyaml ServerYaml
	err := yaml.Unmarshal(data, &srvyaml)
	if err != nil {
		logger.Printf("YAMLZ.UNMARSHA err:%v, file:conf/server.yaml", err)
		return
	}
	Server = &srvyaml
}

func GetDockerItem(name string) *DockerItem {
	//docker_item := DockerItem{}
	for i := 0; i < len(Server.Commands.Docker); i++ {
		if Server.Commands.Docker[i].Name == name {
			return &Server.Commands.Docker[i]
		}
	}
	return nil
}
