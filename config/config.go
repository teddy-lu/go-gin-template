package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type DbConfig struct {
	Name        string
	Host        string
	Port        int64
	Username    string
	Password    string
	Charset     string
	MaxIdleCons int `mapstructure:"max_idle_cons"`
	MaxOpenCons int `mapstructure:"max_open_cons"`
}

type RedisConfig struct {
	Host string
	Port int64
	Auth string
	Db   int
}
type LogConfig struct {
	Level       string
	FileFormat  string `mapstructure:"file_format"`
	MaxSaveDays int64  `mapstructure:"max_save_days"`
	FileType    string `mapstructure:"file_type"`
}

type Config struct {
	Filename     string
	Name         string
	Mode         string
	Addr         int64
	MaxPingCount int64 `mapstructure:"max_ping_count"`
	Db           *DbConfig
	Redis        *RedisConfig
	Log          *LogConfig
}

func Run(file string) (*Config, error) {
	c := Config{Filename: file}

	if err := c.init(); err != nil {
		return nil, err
	}

	setDefault()

	var conf Config
	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(fmt.Errorf("config params unable to decode into struct, %v", err))
	}

	c.watchConfig()

	return &conf, nil
}

// 设置配置文件默认值
func setDefault() {
	viper.SetDefault("name", "go-project")
	viper.SetDefault("mode", "info")
	viper.SetDefault("addr", 8080)
	viper.SetDefault("max_ping_count", 3)

	viper.SetDefault("db.name", "")
	viper.SetDefault("db.host", "127.0.0.1")
	viper.SetDefault("db.port", 3306)
	viper.SetDefault("db.username", "")
	viper.SetDefault("db.max_idle_cons", 10)
	viper.SetDefault("db.max_open_cons", 2)

	viper.SetDefault("redis.host", "127.0.0.1")
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("redis.auth", "")
	viper.SetDefault("redis.db", 0)

	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.file_format", "%Y%m%d")
	viper.SetDefault("log.max_save_days", 30)
	viper.SetDefault("log.file_type", "one")
}

func (c *Config) init() error {
	if c.Filename != "" {
		viper.SetConfigFile(c.Filename)
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
