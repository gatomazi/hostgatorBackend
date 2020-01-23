package router

import (
	"fmt"
	"net/http"

	"agplanservoperadoraapi/functions/migracao"
	"agplanservoperadoraapi/middleware"
	"strings"

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

func callBack(c *gin.Context, param string) bool {

	switch param {
	case "jwt":
		var auth = c.GetHeader("authorization")
		var token string

		if strings.Contains(auth, "Bearer ") || auth == "" {
			return false
		}
		token = strings.Replace(auth, "Bearer ", "", 1)

		if token == "" {
			return false
		}
		fmt.Println(auth)
		return false
	case "google":
		return false
	default:
		return false
	}
}
