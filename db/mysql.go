package db

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"sync"
)

type MysqlPool struct {
}

var instance *MysqlPool
var once sync.Once

var db *gorm.DB
var err error

func GetMysqlPool() *MysqlPool {
	once.Do(func() {
		// 初始化连接池
		instance = &MysqlPool{}
	})
	return instance
}

func (pool MysqlPool) InitPool() (isSuc bool) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.name"),
		viper.GetString("db.charset"),
	)

	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(errors.New("init mysql pool failed"))
		return false
	}

	db.DB().SetMaxIdleConns(viper.GetInt("db.max_idle_cons"))
	db.DB().SetMaxOpenConns(viper.GetInt("db.max_open_cons"))
	return true
}
