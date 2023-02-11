package HaloResponse

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type WithData struct {
	Data    interface{} `json:"data"`
	Code    int64       `json:"code"`
	Message string      `json:"message"`
}

type WithInfo struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func Data(ctx *gin.Context, data interface{}, message string, code int64) {
	ctx.JSON(http.StatusOK, WithData{data, code, message})
}

func Info(ctx *gin.Context, message string, code int64) {
	ctx.JSON(
		http.StatusOK,
		WithInfo{
			code,
			message,
		})
}

func Fail(ctx *gin.Context, message string, code int64, httpCode ...int) {
	switch len(httpCode) {
	case 1:
		ctx.JSON(httpCode[0], WithInfo{code, message})
	default:
		ctx.JSON(http.StatusOK, WithInfo{code, message})
	}
}
