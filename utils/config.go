package utils

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var _config_map map[string]*viper.Viper = make(map[string]*viper.Viper)

func initFromReadConfigFile(filename string) *viper.Viper {
	_config := viper.New()
	_config.SetConfigName(filename)           // name of config file (without extension)
	_config.SetConfigType("yaml")             // REQUIRED if the config file does not have the extension in the name
	_config.AddConfigPath(APP_PATH + "/conf") // path to look for the config file in
	//viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	_config.AddConfigPath(".")    // optionally look for config in the working directory
	err := _config.ReadInConfig() // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		fmt.Errorf("Fatal error config file: %v \n", err)
	}

	_config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	_config.WatchConfig()
	_config_map[filename] = _config
	return _config
}

func Config(name string) *viper.Viper {
	val, ok := _config_map[name]
	if !ok {
		val = initFromReadConfigFile(name)
	}
	return val
}
