package prices

import (
	"log"
	"net/http"
	"testeH/database"
	"testeH/models"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

var db = database.GetDB()

//GetAll -
func GetAll(c *gin.Context) {
	var shared map[string]interface{} = make(map[string]interface{}, 0)
	var products map[string]interface{} = make(map[string]interface{}, 0)

	rowsPrice, err := db.Query("SELECT * FROM prices;")
	if err != nil {
		log.Fatal(err)
	}
	defer rowsPrice.Close()

	var prices []interface{}

	for rowsPrice.Next() {
		var singlePrice models.Product
		var objMount map[string]models.Product = make(map[string]models.Product, 0)
		//Scan nas variaveis do banco para a variavel vinho
		rowsPrice.Scan(&singlePrice.ID, &singlePrice.Name, &singlePrice.CleanName)
		nameIndex := singlePrice.CleanName
		singlePrice.CleanName = ""
		objMount[nameIndex] = singlePrice
		prices = append(prices, objMount)
	}

	products["products"] = prices
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
