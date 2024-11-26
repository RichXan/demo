package global

import (
	"fmt"
	"x-project/config"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var (
	AppConfig config.Config
)

func InitConfig(configPath string) {
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		viper.AddConfigPath("./")
		viper.AddConfigPath("./config")
		// 添加容器下配置文件路径
		viper.AddConfigPath("/etc/x-project/")
		viper.SetConfigName("config")
		viper.SetConfigType("yml")
	}
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		panic("Read config failed")
	}
	err = viper.Unmarshal(&AppConfig, func(config *mapstructure.DecoderConfig) {
		config.TagName = "yaml"
		// do anything your like
	})
	if err != nil {
		panic("Unmarshal config failed")
	}
	fmt.Printf("project: %s, port: %d, env: %s\n", AppConfig.System.Name, AppConfig.System.Port, AppConfig.System.Env)
}
