package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"json-validator/internal/messages"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}

func GenerateRespone(c *gin.Context, valid bool, filename string) ([]byte, error) {
	var message string
	if valid {
		message = fmt.Sprintf(messages.VALID, filename)

	} else {
		message = fmt.Sprintf(messages.NOT_VALID, filename)
	}

	response := Response{Valid: valid, Message: message}
	jsonData, err := json.Marshal(response)
	if err != nil {
		c.String(http.StatusInternalServerError, messages.INTERNAL_ERR)
		return nil, err
	}
	return jsonData, nil
}
