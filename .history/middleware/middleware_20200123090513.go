package middleware

import "github.com/gin-gonic/gin"

func SetHeaders(c *gin.Context, m string) {
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, Content-Length")
	c.Header("Content-type", "application/json")
	c.Header("Access-Control-Allow-Methods", m)
}
