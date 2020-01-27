package prices

import (
	"database/sql"
	"hostgatorBackend/database"
	"hostgatorBackend/models"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

var db = database.GetDB()

//GetAll -
func GetAll(c *gin.Context) {
	var err error
	//Variaveis criadas para criar o retorno conforme exibido
	var shared map[string]interface{} = make(map[string]interface{}, 0)
	var products map[string]interface{} = make(map[string]interface{}, 0)

	rowsPrice, err := db.Query("SELECT * FROM plans;")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.MakeNewErrorResponse("Internal Error", http.StatusBadRequest, "Ocorreu um erro ao consultar o banco", err.Error()))
		return
		// log.Fatal(err)
	}
	defer rowsPrice.Close()

	products["products"], err = selectPlan(rowsPrice)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.MakeNewErrorResponse("Internal Error", http.StatusBadRequest, "Ocorreu um erro ao consultar o banco", err.Error()))
		return
	}
	shared["shared"] = products

	//Retorno Json
	c.JSON(http.StatusOK, shared)
	return
}

//GetOne -
func GetOne(c *gin.Context) {
	//Variaveis criadas para criar o retorno conforme exibido
	var err error
	var singlePlan models.Plan
	row := db.QueryRow("SELECT * FROM plans WHERE id = ?;", c.Param("id"))
	err = row.Scan(&singlePlan.ID, &singlePlan.Name, &singlePlan.CleanName)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.MakeNewErrorResponse("Internal Error", http.StatusBadRequest, "Ocorreu um erro ao consultar a tabela plans pelo id "+c.Param("id"), err.Error()))
		return
	}
	singlePlan.CleanName = ""

	singlePlan.Cycle, err = selectCycle(singlePlan.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.MakeNewErrorResponse("Internal Error", http.StatusBadRequest, "Ocorreu um erro ao buscar os cycles", err.Error()))
		return
	}

	c.JSON(http.StatusOK, singlePlan)
}

func selectPlan(rowsPrice *sql.Rows) (prices interface{}, err error) {

	var objMount map[string]models.Plan = make(map[string]models.Plan, 0)
	for rowsPrice.Next() {
		var singlePlan models.Plan
		//Scan nas variaveis do banco para a variavel singlePlan
		err = rowsPrice.Scan(&singlePlan.ID, &singlePlan.Name, &singlePlan.CleanName)
		if err != nil {
			return
		}

		nameIndex := singlePlan.CleanName
		singlePlan.CleanName = ""

		singlePlan.Cycle, err = selectCycle(singlePlan.ID)
		if err != nil {
			return
		}

		//Criando o index com o nome sem espaço do tipo
		objMount[nameIndex] = singlePlan
		prices = objMount
		// prices = append(prices, objMount)
	}
	return
}

func selectCycle(idPlan int) (cycles interface{}, err error) {
	//Query para selecionar os ciclos daquele plano
	rowsPrice, err := db.Query("SELECT type, priceRenew, priceOrder, months FROM cycles WHERE idPlan = ?;", idPlan)
	if err != nil {
		return
		// log.Fatal(err)
	}

	var objMount map[string]models.CycleInfo = make(map[string]models.CycleInfo, 0)
	for rowsPrice.Next() {
		var singleCycle models.CycleInfo
		//Scan nas variaveis do banco para a variavel singleCycle
		rowsPrice.Scan(&singleCycle.Type, &singleCycle.PriceRenew, &singleCycle.PriceOrder, &singleCycle.Months)
		nameIndex := singleCycle.Type
		singleCycle.Type = ""
		//Criando o index com o nome do tipo
		objMount[nameIndex] = singleCycle
		cycles = objMount
	}
	return
}
