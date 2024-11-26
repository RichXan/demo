package global

import (
	"fmt"
	"xproject/config"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var (
	Config config.Config
)

func InitConfig(configPath string) {
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		panic("Read config failed")
	}
	err = viper.Unmarshal(&Config, func(config *mapstructure.DecoderConfig) {
		config.TagName = "yaml"
		// do anything your like
	})
	if err != nil {
		panic("Unmarshal config failed")
	}
	fmt.Println(Config.System.Env)
}
