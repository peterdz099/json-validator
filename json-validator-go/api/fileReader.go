package api

import (
	"bytes"
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
		fmt.Println("Error reading request body:", err)
		c.String(http.StatusBadRequest, err.Error())
		return Policy{}, "", err
	}

	if file != nil {
		if fileExtension := strings.ToLower(file.Filename[len(file.Filename)-4:]); fileExtension != "json" {
			fmt.Println("file error: Wrong file extension")
			c.String(http.StatusUnsupportedMediaType, "Wrong file extension")
			return Policy{}, "", err
		}
	}

	fileContent, err := file.Open()
	if err != nil {
		fmt.Println("Error opening file:", err)
		c.String(http.StatusInternalServerError, err.Error())
		return Policy{}, "", err
	}
	defer fileContent.Close()

	fileBytes, err := io.ReadAll(fileContent)
	if err != nil {
		fmt.Println("Error reading file:", err)
		c.String(http.StatusInternalServerError, err.Error())
		return Policy{}, "", err
	}

	var policy Policy
	decoder := json.NewDecoder(bytes.NewReader(fileBytes))
	decoder.DisallowUnknownFields()
	decodeErr := decoder.Decode(&policy)

	if decodeErr != nil {
		fmt.Println("Error decoding JSON:", decodeErr)
		c.String(http.StatusBadRequest, "invalid format: invalid json format")
		return Policy{}, "", decodeErr
	}

	isFormatValid, err := verifyJsonFormat(policy)
	if !isFormatValid && err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return Policy{}, "", err
	}
	return policy, file.Filename, nil
}
