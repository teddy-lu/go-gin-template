package db

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func (pool MysqlPool) InitPool() (orm *gorm.DB, isSuc bool) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.name"),
		viper.GetString("db.charset"),
	)
	dbConfig := mysql.New(mysql.Config{
		DSN: dsn,
	})

	db, err = gorm.Open(dbConfig)
	if err != nil {
		panic(errors.New("init mysql pool failed"))
		return nil, false
	}

	dbase, err := db.DB()
	if err != nil {
		panic(errors.New("get mysql DB failed"))
		return nil, false
	}

	// 设置最大空闲连接数
	dbase.SetMaxIdleConns(viper.GetInt("db.max_idle_cons"))
	// 设置最大连接数
	dbase.SetMaxOpenConns(viper.GetInt("db.max_open_cons"))
	// 设置每个链接的过期时间
	//dbase.SetConnMaxLifetime(time.Duration(viper.GetInt("db.conn_max_lifetime")) * time.Second)

	return db, true
}
