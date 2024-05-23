package dao

import (
	"context"
	"go-gin-demo/internal/models"
	"gorm.io/gorm"
)

// demoDao定义demo的DAO层
type demoDao struct {
	// 定义数据库连接等相关操作
	db *gorm.DB
}

func NewDemoDao(db *gorm.DB) DemoStore {
	return &demoDao{db: db}
}

// DemoStore 这里定义了Demo的CRUD方法
type DemoStore interface {
	Create(ctx context.Context, data *models.Demo) error
	List(ctx context.Context) ([]*models.Demo, error)
}

func (dao *demoDao) Create(ctx context.Context, data *models.Demo) error {
	return dao.db.WithContext(ctx).Create(data).Error
}

func (dao *demoDao) List(ctx context.Context) ([]*models.Demo, error) {
	var dm []*models.Demo
	err := dao.db.WithContext(ctx).Find(&dm).Error
	return dm, err
}
