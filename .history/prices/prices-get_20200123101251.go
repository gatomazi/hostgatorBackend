package prices

import (
	"hostgator/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetAll -
func GetAll(c *gin.Context) {
	// c.Request.URL.Query()
	var result []*models.Product
	var res map[string][string]interface{} = make(map[string]map[string]interface{}, 0)

	teste := &models.Product{ID: 2, Name: "Teste"}

	result = append(result, teste)

	res["shared"]["products"] = result
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
