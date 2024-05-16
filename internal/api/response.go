package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Json(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func Created(c *gin.Context, data interface{}) {
	code := http.StatusCreated
	c.JSON(code, Response{
		Code:    code,
		Message: "Created Success",
		Data:    data,
	})
}

func Deleted(c *gin.Context, data interface{}) {
	code := http.StatusNoContent
	c.JSON(code, Response{
		Code:    code,
		Message: "Deleted Success",
		Data:    data,
	})
}

func Error(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
