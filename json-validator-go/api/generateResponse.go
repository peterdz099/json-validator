package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}

func generateRespone(c *gin.Context, valid bool, filename string) ([]byte, error) {
	var message string
	if valid {
		message = fmt.Sprintf("JSON from file %s is VALID", filename)
	} else {
		message = fmt.Sprintf("JSON from file %s is NOT VALID", filename)
	}

	response := Response{Valid: valid, Message: message}
	jsonData, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		c.String(http.StatusInternalServerError, "Internal Server Error - Generating Response")
		return nil, err
	}
	return jsonData, nil
}
