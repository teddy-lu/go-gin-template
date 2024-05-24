package logger

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		// 延迟
		latency := time.Since(start)

		// 获取GET请求的参数并且转换为json
		params := c.Request.URL.Query()
		paramsJson, _ := json.Marshal(params)

		// 获取POST请求的参数并且转换为json
		body := c.Request.Body
		bodyBytes, _ := io.ReadAll(body)
		var jsonData map[string]interface{}
		_ = json.Unmarshal(bodyBytes, &jsonData)
		bodyJson, _ := json.Marshal(jsonData)

		Info(fmt.Sprintf("%s - [%s] \"%s %s %s %d %s [Body Data: %s] [Query Data: %s] \"%s\" %s\"",
			c.ClientIP(),
			time.Now().Format(time.RFC1123),
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.Proto,
			c.Writer.Status(),
			latency,
			string(bodyJson),
			paramsJson,
			c.Request.UserAgent(),
			c.Errors.ByType(gin.ErrorTypePrivate).String(),
		))
	}
}
