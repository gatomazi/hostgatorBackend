package prices

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"hostgatorBackend/database"
	"hostgatorBackend/models"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

var db = database.GetDB()

//GetAll -
func GetAll(c *gin.Context) {
	//Variaveis criadas para criar o retorno conforme exibido
	var shared map[string]interface{} = make(map[string]interface{}, 0)
	var products map[string]interface{} = make(map[string]interface{}, 0)

	rowsPrice, err := db.Query("SELECT * FROM plans;")
	if err != nil {
		log.Fatal(err)
	}
	defer rowsPrice.Close()

	products["products"] = selectPlan(rowsPrice)
	shared["shared"] = products

	//Retorno Json
	c.JSON(http.StatusOK, shared)
	return
}

//GetOne -
func GetOne(c *gin.Context) {
	//Variaveis criadas para criar o retorno conforme exibido

	var singlePlan models.Plan
	row := db.QueryRow("SELECT * FROM plans WHERE id = ?;", c.Param("id"))
	_ = row.Scan(&singlePlan.ID, &singlePlan.Name, &singlePlan.CleanName)
	singlePlan.Cycle = selectCycle(singlePlan.ID)

	c.JSON(http.StatusOK, singlePlan)
}
func selectPlan(rowsPrice *sql.Rows) (prices interface{}) {

	var objMount map[string]models.Plan = make(map[string]models.Plan, 0)
	for rowsPrice.Next() {
		var singlePlan models.Plan
		//Scan nas variaveis do banco para a variavel vinho
		rowsPrice.Scan(&singlePlan.ID, &singlePlan.Name, &singlePlan.CleanName)
		nameIndex := singlePlan.CleanName
		singlePlan.CleanName = ""

		singlePlan.Cycle = selectCycle(singlePlan.ID)

		//Criando o index com o nome sem espaço do tipo
		objMount[nameIndex] = singlePlan
		fmt.Println(objMount)
		prices = objMount
		// prices = append(prices, objMount)
	}
	return
}

func selectCycle(idPlan int) (cycles []interface{}) {
	//Query para selecionar os ciclos daquele plano
	rowsPrice, err := db.Query("SELECT type, priceRenew, priceOrder, months FROM cycles WHERE idPlan = ?;", idPlan)
	fmt.Println(idPlan)
	if err != nil {
		log.Fatal(err)
	}

	for rowsPrice.Next() {
		var singleCycle models.CycleInfo
		var objMount map[string]models.CycleInfo = make(map[string]models.CycleInfo, 0)
		//Scan nas variaveis do banco para a variavel singleCycle
		rowsPrice.Scan(&singleCycle.Type, &singleCycle.PriceRenew, &singleCycle.PriceOrder, &singleCycle.Months)
		fmt.Println(singleCycle)
		nameIndex := singleCycle.Type
		singleCycle.Type = ""
		//Criando o index com o nome do tipo
		objMount[nameIndex] = singleCycle
		cycles = append(cycles, objMount)
	}
	return
}
