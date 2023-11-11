package middlewares

import "github.com/gin-gonic/gin"

func MyAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"haku_layaya" : "Yushan100!",
		"dummy_user01" : "password",
	})

}