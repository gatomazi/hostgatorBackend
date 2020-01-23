package router

import (
	"hostgator/middleware"
	"hostgator/prices"

	"github.com/gin-gonic/gin"
)

// StartRouter -
func StartRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.OptionsRequest())
	var (
		publicGroup = router.Group("/")
	)

	migracaoRedis(publicGroup)

	return router
}

func migracaoRedis(g *gin.RouterGroup) {
	g.GET("prices", func(c *gin.Context) { middleware.SetHeaders(c, "OPTIONS, GET"); prices.GetAll(c) })
	g.GET("prices/:id", func(c *gin.Context) { middleware.SetHeaders(c, "OPTIONS, GET"); prices.GetOne(c) })
}
