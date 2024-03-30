package handlers

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

func GenerateRespone(c *gin.Context, valid bool, filename string) ([]byte, error) {
	var message string
	if valid {
		message = fmt.Sprintf("JSON from file %s is VALID", filename)
	} else {
		message = fmt.Sprintf("JSON from file %s is NOT VALID", filename)
	}

	response := Response{Valid: valid, Message: message}
	jsonData, err := json.Marshal(response)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error - Generating Response")
		return nil, err
	}
	return jsonData, nil
}
