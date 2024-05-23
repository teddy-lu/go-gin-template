package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go-gin-demo/config"
	"go-gin-demo/db"
	"go-gin-demo/internal/dao"
	"go-gin-demo/pkg/logger"
	"go-gin-demo/routers"
	"gorm.io/gorm"
)

var conf = pflag.StringP("conf", "c", "", "config filepath")

func main() {
	pflag.Parse()

	// 读取配置
	if err := config.Run(*conf); err != nil {
		panic(err)
	}

	// 链接mysql
	dbm, m := db.GetMysqlPool().InitPool()
	if !m {
		logger.Error("init database pool failure...")
		panic("mysql init failed")
	}
	// 链接redis
	db.InitRedis()

	// 初始化日志
	logger.InitLogger()

	// 实例化server参数，并启动gin
	g := createApp(dbm)
	// Listen and Server in 0.0.0.0:8080
	if err := g.Run(fmt.Sprintf(":%s", viper.GetString("addr"))); err != nil {
		return
	}
}

func createApp(dbm *gorm.DB) *gin.Engine {
	demoDao := dao.NewDemoDao(dbm)
	serv := routers.NewServer(demoDao)
	mode := viper.GetString("mode")
	if mode == "debug" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.New()
	serv.SetRouter(g)
	return g
}
