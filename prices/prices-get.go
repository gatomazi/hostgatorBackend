package prices

import (
	"database/sql"
	"log"
	"net/http"
	"testeH/models"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

//GetAll -
func GetAll(c *gin.Context) {

	db, err := sql.Open("mysql", "hostgator:hostgator123@tcp(172.18.0.2:3306)/teste")
	// db, err = sql.Open("mysql", "XHFHkmpD8B:Mb6k8MT6RS@tcp(remotemysql.com)/XHFHkmpD8B")

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// var shared map[string]interface{} = make(map[string]interface{}, 0)
	// var products map[string]interface{} = make(map[string]interface{}, 0)
	// var planos map[string]interface{} = make(map[string]interface{}, 0)

	// PlanoP := models.Product{ID: 5, Name: "Plano P"}
	// PlanoM := models.Product{ID: 12, Name: "Plano M"}
	// PlanoTurbo := models.Product{ID: 16, Name: "Plano Turbo"}

	// planos["planoP"] = PlanoP
	// planos["planoM"] = PlanoM
	// planos["planoTurbo"] = PlanoTurbo

	// products["products"] = planos
	// shared["shared"] = products
	// c.JSON(http.StatusOK, shared)
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
