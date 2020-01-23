package router

import (
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
