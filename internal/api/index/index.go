package index

import (
	"github.com/gin-gonic/gin"
	"go-gin-demo/internal/api"
	"go-gin-demo/internal/service/index"
	"go-gin-demo/pkg/logger"
	"net/http"
)

func Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("test info.....")
		logger.Debug("test debug.....")
		logger.Error("test error.....")

		api.Json(c, 200, "success", nil)
	}
}

func Demo(service index.DemoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := service.List(c)
		if err != nil {
			api.Error(c, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		api.Json(c, 200, "success", data)
	}
}
