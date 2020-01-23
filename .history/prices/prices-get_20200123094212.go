package prices

import (
	"hostgator/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetAll -
func GetAll(c *gin.Context) {
	// c.Request.URL.Query()

	c.JSON(http.StatusOK, "")
	return
}

//GetOne -
func GetOne(c *gin.Context) {
	// c.Param("id")
	var result models.Product
	var res map[string][string]interface{} = make(map[string][string]interface{}, 0)

	// res["shared"] =

	c.JSON(http.StatusOK, "")
}
