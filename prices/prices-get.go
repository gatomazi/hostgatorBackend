package prices

import (
	"database/sql"
	"fmt"
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
	products["products"] = selectPlan(rowsPrice)
	shared["shared"] = products

	c.JSON(http.StatusOK, shared)
	return
}

func selectPlan(rowsPrice *sql.Rows) (prices []interface{}) {

	for rowsPrice.Next() {
		var singlePlan models.Plan
		var objMount map[string]models.Plan = make(map[string]models.Plan, 0)
		//Scan nas variaveis do banco para a variavel vinho
		rowsPrice.Scan(&singlePlan.ID, &singlePlan.Name, &singlePlan.CleanName)
		nameIndex := singlePlan.CleanName
		singlePlan.CleanName = ""

		singlePlan.Cycle = selectCycle(singlePlan.ID)

		objMount[nameIndex] = singlePlan
		prices = append(prices, objMount)
	}
	return
}

func selectCycle(idPlan int) (cycles []interface{}) {
	rowsPrice, err := db.Query("SELECT type, priceRenew, priceOrder, months FROM cycles WHERE idPlan = ?;", idPlan)
	fmt.Println(idPlan)
	if err != nil {
		log.Fatal(err)
	}
	for rowsPrice.Next() {
		var singleCycle models.CycleInfo
		var objMount map[string]models.CycleInfo = make(map[string]models.CycleInfo, 0)
		//Scan nas variaveis do banco para a variavel vinho
		rowsPrice.Scan(&singleCycle.Type, &singleCycle.PriceRenew, &singleCycle.PriceOrder, &singleCycle.Months)
		fmt.Println(singleCycle)
		nameIndex := singleCycle.Type
		singleCycle.Type = ""

		objMount[nameIndex] = singleCycle
		cycles = append(cycles, objMount)
	}
	return
}

//GetOne -
func GetOne(c *gin.Context) {
	// c.Param("id")
	var result []models.Plan
	var res interface{} = make(map[string]interface{}, 0)

	res.(map[string]interface{})["shared"].(map[string]interface{})["products"] = result

	c.JSON(http.StatusOK, res)
}
