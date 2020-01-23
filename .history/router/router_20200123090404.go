package router

import (
	"agplanservoperadoraapi/functions/migracao"
	"net/http"

	"github.com/gin-gonic/gin"
)

// StartRouter -
func StartRouter() *gin.Engine {
	router := gin.Default()
	router.Use(optionsRequest())
	var (
		publicGroup = router.Group("/OpPlanserv/v1")
	)

	migracaoRedis(publicGroup)

	return router
}

func migracaoRedis(g *gin.RouterGroup) {
	g.GET("migraRedis", func(c *gin.Context) { setHeaders(c, "OPTIONS, GET"); migracao.MigraCollectionRedis(c) })
	g.GET("testeRedis", func(c *gin.Context) { setHeaders(c, "OPTIONS, GET"); migracao.TesteRedis(c) })
}

func setHeaders(c *gin.Context, m string) { // TODO passar para middleware
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, Content-Length")
	c.Header("Content-type", "application/json")
	c.Header("Access-Control-Allow-Methods", m)
}

func optionsRequest() gin.HandlerFunc {

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
