package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func readFile(c *gin.Context) (Policy, string, error) {
	file, err := c.FormFile("file")

	if err != nil {
		if file != nil {
			if fileExtension := strings.ToLower(file.Filename[len(file.Filename)-4:]); fileExtension != "json" {
				fmt.Println("Wrong file extension:", err)
				c.String(http.StatusUnsupportedMediaType, "Wrong file extension")
				return Policy{}, "", err
			}
		}
		fmt.Println("Error reading request body:", err)
		c.String(http.StatusBadRequest, "Bad request")
		return Policy{}, "", err
	}

	fileContent, err := file.Open()
	if err != nil {
		fmt.Println("Error opening file:", err)
		c.String(http.StatusInternalServerError, "Internal server error - Openning file")
		return Policy{}, "", err
	}
	defer fileContent.Close()

	fileBytes, err := io.ReadAll(fileContent)
	if err != nil {
		fmt.Println("Error reading file:", err)
		c.String(http.StatusInternalServerError, "Internal server error - Reading file")
		return Policy{}, "", err
	}

	var policy Policy
	err = json.Unmarshal(fileBytes, &policy)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		c.String(http.StatusBadRequest, "Invalid JSON format")
		return Policy{}, "", err
	}

	return policy, file.Filename, nil
}
