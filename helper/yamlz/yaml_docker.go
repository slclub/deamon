package yamlz

type DockerItem struct {
	ID            int    `yaml:"ID"`
	Name          string `yaml:"Name"`
	Open          bool   `yaml:"Open"`
	Mode          string `yaml:"Mode"`
	CheckCmd      string `yaml:"CheckCmd"`
	CheckValue    string `yaml:"CheckValue"`
	RestartBefore string `yaml:"RestartBefore"`
	Restart       string `yaml:"Restart"`
	Start         string `yaml:"Start"`
	Stop          string `yaml:"Stop"`
}
