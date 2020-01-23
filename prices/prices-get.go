package prices

import (
	"net/http"
	"testeH/models"

	"github.com/gin-gonic/gin"
)

//GetAll -
func GetAll(c *gin.Context) {
	// c.Request.URL.Query()
	// var result *models.Product
	var shared map[string]interface{} = make(map[string]interface{}, 0)
	var products map[string]interface{} = make(map[string]interface{}, 0)
	var planos map[string]interface{} = make(map[string]interface{}, 0)

	PlanoP := models.Product{ID: 5, Name: "Plano P"}
	PlanoM := models.Product{ID: 12, Name: "Plano M"}
	PlanoTurbo := models.Product{ID: 16, Name: "Plano Turbo"}

	planos["planoP"] = PlanoP
	planos["planoM"] = PlanoM
	planos["planoTurbo"] = PlanoTurbo

	products["products"] = planos
	shared["shared"] = products
	c.JSON(http.StatusOK, shared)
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
