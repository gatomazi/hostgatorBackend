package prices

import (
	"agplanservoperadoraapi/models"
	"agplanservoperadoraapi/utils/filter"
	"agplanservoperadoraapi/utils/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetAll -
func GetAll(c *gin.Context) {
	collection := "usuario"

	redisDAO := redis.NewRedisDAO(collection)
	redisParams, paginator := filter.RedisParamFilter(c.Request.URL.Query(), collection, nil)
	res := redisDAO.GetEntityRedis(redisParams, paginator)

	c.JSON(http.StatusOK, models.MakeNewSuccessResponse(res))
	return
}
