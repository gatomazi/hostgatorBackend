package prices

import (
	"agplanservoperadoraapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetAll -
func GetAll(c *gin.Context) {
	collection := "usuario"

	c.JSON(http.StatusOK, models.MakeNewSuccessResponse(res))
	return
}
