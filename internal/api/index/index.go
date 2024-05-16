package index

import (
	"github.com/gin-gonic/gin"
	"go-gin-demo/internal/api"
	"go-gin-demo/pkg/logger"
)

func Index(c *gin.Context) {
	logger.Info("test info.....")
	logger.Debug("test debug.....")
	logger.Error("test error.....")
	api.Json(c, 200, "success", nil)
}
