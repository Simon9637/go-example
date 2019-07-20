package config

import (
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
	"log"
)

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

type Config struct {
	Name string
}

func (config *Config) initConfig() error {
	if config.Name != "" {
		// 解析指定文件
		viper.SetConfigFile(config.Name)
	} else {
		// 为指定文件时解析默认文件
		viper.AddConfigPath("conf")
		viper.SetConfigName("conf")
	}
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// 监控配置文件变化并热加载程序
func (config *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}
