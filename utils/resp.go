package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func Result(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}
