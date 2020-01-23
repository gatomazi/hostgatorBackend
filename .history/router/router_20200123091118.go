package router

import (
	"hostagator/price"
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

func teste(){
	teste := "shared": {
		"products": {
		"planoP": {
		"name": "Plano P",
		"id": 5,
		"cycle": {
		"monthly": {
		"priceRenew": "24.19",
		"priceOrder": "24.19",
		"months": 1
		},
		"semiannually": {
		"priceRenew": "128.34",
		"priceOrder": "128.34",
		"months": 6
		},
		"biennially": {
		"priceRenew": "393.36",
		"priceOrder": "393.36",
		"months": 24
		},
		"triennially": {
		"priceRenew": "561.13",
		"priceOrder": "561.13",
		"months": 36
		},
		"quarterly": {
		"priceRenew": "67.17",
		"priceOrder": "67.17",
		"months": 3
		},
		"annually": {
		"priceRenew": "220.66",
		"priceOrder": "220.66",
		"months": 12
		}
		}
		},
	}
	}
}