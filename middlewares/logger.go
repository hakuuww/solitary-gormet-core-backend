package middlewares

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func MyLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string{
		return fmt.Sprintf("%s %s %s\n", params.ClientIP, params.Method, params.Path)
	})
}