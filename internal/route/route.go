package route

import (
	"github.com/gin-gonic/gin"
	"go-gin-demo/internal/api/index"
	"net/http"
)

func SetRouter(g *gin.Engine) *gin.Engine {
	g.Use(gin.Recovery())
	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 NOT FOUND")
	})

	// 设置路由
	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	g.GET("/index", index.Index)

	return g
}
