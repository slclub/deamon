package yamlz

type DockerItem struct {
	ID            int    `yaml:"ID"`
	Name          string `yaml:"Name"`
	Open          bool   `yaml:"Open"`
	CheckCmd      string `yaml:"CheckCmd"`
	CheckValue    string `yaml:"CheckValue"`
	RestartBefore string `yaml:"RestartBefore"`
	Restart       string `yaml:"Restart"`
}
