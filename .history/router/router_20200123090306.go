package router

import (
	"net/http"

	"agplanservoperadoraapi/functions/migracao"
	"agplanservoperadoraapi/middleware"

	"github.com/gin-gonic/gin"
)

// StartRouter -
func StartRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.OptionsRequest())
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

func verifyOptions(c *gin.Context) {
	setHeaders(c, "OPTIONS, POST, GET, PUT, DELETE")
	c.Status(http.StatusNoContent)
}

func setHeaders(c *gin.Context, m string) { // TODO passar para middleware
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, Content-Length")
	c.Header("Content-type", "application/json")
	c.Header("Access-Control-Allow-Methods", m)
}
