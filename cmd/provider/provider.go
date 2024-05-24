package provider

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-demo/config"
	"go-gin-demo/db"
	"go-gin-demo/internal/dao"
	"go-gin-demo/internal/models"
	"go-gin-demo/pkg/logger"
	"go-gin-demo/routers"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type App struct {
	Conf   *config.Config
	Server *http.Server
}

func newApp(config *config.Config, server *http.Server) *App {
	return &App{Conf: config, Server: server}
}

func CreateServEngine(cfg *config.Config, g *gin.Engine) *App {
	mainApp := newApp(cfg, &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Addr),
		Handler:      g,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})
	return mainApp
}

func CreateDb(cfg *config.Config) *gorm.DB {
	// 链接mysql
	dbm, m := db.GetMysqlPool().InitPool(cfg)
	if !m {
		logger.Error("init database pool failure...")
		panic("mysql init failed")
	}
	mErr := MigratorDb(dbm)
	if mErr != nil {
		logger.Error("migrator database failure...")
		panic("mysql database migrator failed")
	}
	// 链接redis
	db.InitRedis(cfg)

	return dbm
}

func MigratorDb(dbm *gorm.DB) error {
	fmt.Println("migrator database...")
	return dbm.AutoMigrate(&models.Demo{})
}

func CreateGinServer(dbm *gorm.DB, mode string) *gin.Engine {
	demoDao := dao.NewDemoDao(dbm)
	serv := routers.NewServer(demoDao)
	if mode == "debug" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.New()
	serv.SetRouter(g)
	return g
}
