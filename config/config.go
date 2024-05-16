package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Run(cfg string) error {
	c := Config{Name: cfg}

	if err := c.init(); err != nil {
		return err
	}

	c.watchConfig()

	return nil
}

func (c *Config) init() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("./..")
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Failed to read config file: %s \n", err))
	}

	return nil
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
