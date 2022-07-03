package initialize

import (
	"api-platform/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Configs() {
	var config = "config.yaml"
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic("Fatal error config file")
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Errorf("config file changed:%s", e.Name)
		if err = v.Unmarshal(&global.Config); err != nil {
			fmt.Errorf("unmarshal config error:%s", err)
		}
	})

	if err = v.Unmarshal(&global.Config); err != nil {
		fmt.Errorf("unmarshal config error %s", err)
	}

}
