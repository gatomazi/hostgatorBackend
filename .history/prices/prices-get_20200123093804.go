package prices

import (
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

}
