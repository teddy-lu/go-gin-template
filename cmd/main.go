package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go-gin-demo/config"
	"go-gin-demo/db"
	"go-gin-demo/internal/route"
	"go-gin-demo/pkg/logger"
)

var conf = pflag.StringP("conf", "c", "", "config filepath")

func main() {
	pflag.Parse()

	// 读取配置
	if err := config.Run(*conf); err != nil {
		panic(err)
	}

	// 链接mysql
	if m := db.GetMysqlPool().InitPool(); !m {
		logger.Error("init database pool failure...")
		panic("mysql init failed")
	}
	// 链接redis
	db.InitRedis()

	// 初始化日志
	logger.InitLogger()

	gin.SetMode(viper.GetString("mode"))
	g := gin.New()
	g = route.SetRouter(g)
	// Listen and Server in 0.0.0.0:8080
	if err := g.Run(viper.GetString("addr")); err != nil {
		return
	}
}
