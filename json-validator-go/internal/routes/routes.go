package routes

import (
	"json-validator/internal/handlers"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/validate", Validate)
	return router
}

func Validate(c *gin.Context) {
	policy, filename, err := handlers.ReadFile(c)
	if err != nil {
		return
	} else {
		valid := handlers.IsJsonValid(policy)
		returnData, err := handlers.GenerateRespone(c, valid, filename)

		if err != nil {
			return
		}
		c.Data(http.StatusOK, "application/json", returnData)
	}
}
