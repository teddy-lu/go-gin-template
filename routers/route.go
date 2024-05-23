package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-demo/internal/api/index"
	"go-gin-demo/internal/dao"
	idx "go-gin-demo/internal/service/index"
	"go-gin-demo/pkg/logger"
	"net/http"
	"time"
)

type MyServer struct {
	demoService *idx.DemoService
}

func NewServer(demoStore dao.DemoStore) *MyServer {
	var demoService = idx.NewDemoService(demoStore)
	return &MyServer{demoService: demoService}
}

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		// 延迟
		latency := time.Since(start)

		logger.Debug(fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"",
			c.ClientIP(),
			time.Now().Format(time.RFC1123),
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.Proto,
			c.Writer.Status(),
			latency,
			c.Request.UserAgent(),
			c.Errors.ByType(gin.ErrorTypePrivate).String(),
		))
	}
}

func (svr *MyServer) SetRouter(g *gin.Engine) *gin.Engine {
	g.Use(MyLogger())
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

	g.GET("/index", index.Index())
	g.GET("/demo", index.Demo(*svr.demoService))

	return g
}
