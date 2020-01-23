package router

import (
	"agplanservoperadoraapi/functions/migracao"
	"hostgator/middleware"

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
	g.GET("migraRedis", func(c *gin.Context) { middleware.SetHeaders(c, "OPTIONS, GET"); migracao.MigraCollectionRedis(c) })
}
