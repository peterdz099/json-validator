package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

func readFile(c *gin.Context) (Policy, string, error) {
	file, err := c.FormFile("file")

	if err != nil {
		fmt.Println("Error reading request body:", err)
		c.String(http.StatusBadRequest, "Bad request")
		return Policy{}, "", err
	}

	if file != nil {
		if fileExtension := strings.ToLower(file.Filename[len(file.Filename)-4:]); fileExtension != "json" {
			fmt.Println("Wrong file extension:", err)
			c.String(http.StatusUnsupportedMediaType, "Wrong file extension")
			return Policy{}, "", err
		}
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
	correct := isPolicyStructureCorrect(policy)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		c.String(http.StatusBadRequest, "Invalid JSON format")
		return Policy{}, "", err
	} else if !correct {
		fmt.Println("Invalid Policy structure")
		c.String(http.StatusBadRequest, "Invalid JSON format")
		return Policy{}, "", errors.New("reading JSON: not valid JSON format")
	}
	return policy, file.Filename, nil
}

func isPolicyStructureCorrect(policy Policy) bool {
	policyDocumentBytes, err := json.Marshal(policy.PolicyDocument)
	if err != nil {
		return false
	}

	var policyDocument PolicyDocument
	err = json.Unmarshal(policyDocumentBytes, &policyDocument)
	if err != nil {
		return false
	}

	policyDocumentStatmentBytes, err := json.Marshal(policy.PolicyDocument.Statement)
	if err != nil {
		return false
	}

	var policyDocumentStatment []Statement
	err = json.Unmarshal(policyDocumentStatmentBytes, &policyDocumentStatment)

	if err == nil {
		for _, statement := range policyDocumentStatment {
			typeOfResource := reflect.TypeOf(statement.Resource)
			if typeOfResource != reflect.TypeOf("") && typeOfResource != reflect.TypeOf([]string{}) {
				return false
			}
		}
		return true
	} else {
		return false
	}

}
