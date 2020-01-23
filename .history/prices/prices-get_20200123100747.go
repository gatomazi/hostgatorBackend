package prices

import (
	"fmt"
	"hostgator/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetAll -
func GetAll(c *gin.Context) {
	// c.Request.URL.Query()
	var result []models.Product
	var res map[string]interface{} = make(map[string]interface{}, 0)

	fmt.Println(result)
	res["shared"].(map[string]interface{})["products"] = result
	c.JSON(http.StatusOK, res)
	return
}

//GetOne -
func GetOne(c *gin.Context) {
	// c.Param("id")
	var result []models.Product
	var res interface{} = make(map[string]interface{}, 0)

	res.(map[string]interface{})["shared"].(map[string]interface{})["products"] = result

	c.JSON(http.StatusOK, res)
}
