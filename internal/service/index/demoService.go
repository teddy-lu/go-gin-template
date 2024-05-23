package index

import (
	"context"
	"go-gin-demo/internal/dao"
	"go-gin-demo/internal/models"
	"go-gin-demo/pkg/logger"
)

type DemoService struct {
	DemoStore dao.DemoStore
}

func NewDemoService(store dao.DemoStore) *DemoService {
	return &DemoService{DemoStore: store}
}

func (ds *DemoService) List(ctx context.Context) ([]*models.Demo, error) {
	model, err := ds.DemoStore.List(ctx)
	if err != nil {
		logger.Error("demo store list error", err)
		return nil, err
	}
	return model, nil
}
