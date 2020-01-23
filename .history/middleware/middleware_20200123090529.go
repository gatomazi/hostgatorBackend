package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetHeaders(c *gin.Context, m string) {
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, Content-Length")
	c.Header("Content-type", "application/json")
	c.Header("Access-Control-Allow-Methods", m)
}

func OptionsRequest() gin.HandlerFunc {

	return func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
			c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
			c.AbortWithStatus(http.StatusOK)
		}
	}
}
