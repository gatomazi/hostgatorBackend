package prices

import (
	"agplanservoperadoraapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetAll -
func GetAll(c *gin.Context) {
	collection := "usuario"
	// c.Request.URL.Query()

	c.JSON(http.StatusOK, models.MakeNewSuccessResponse(""))
	return
}
