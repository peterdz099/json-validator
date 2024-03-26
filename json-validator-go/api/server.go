package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ApiServer() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/validate", validate)
	return router
}

func validate(c *gin.Context) {

	policy, filename, err := readFile(c) //if type policy its ok if er
	if err != nil {
		return
	} else {
		valid := isJsonValid(policy)
		returnData, err := generateRespone(c, valid, filename)

		if err != nil {
			return
		}

		c.Data(http.StatusOK, "application/json", returnData)
	}

}
