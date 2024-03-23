package main

//https://stackoverflow.com/questions/76448408/cors-policy-response-to-preflight-request-doesnt-pass-access-control-check-no
//https://stackoverflow.com/questions/76448408/cors-policy-response-to-preflight-request-doesnt-pass-access-control-check-no
//https://en.wikipedia.org/wiki/Cross-origin_resource_sharing

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Statement struct {
	Sid      string   `json:"Sid"`
	Effect   string   `json:"Effect"`
	Action   []string `json:"Action"`
	Resource string   `json:"Resource"`
}

type PolicyDocument struct {
	Version   string      `json:"Version"`
	Statement []Statement `json:"Statement"`
}

type Policy struct {
	PolicyName     string         `json:"PolicyName"`
	PolicyDocument PolicyDocument `json:"PolicyDocument"`
}

type Response struct {
	Filename string `json:"filename"`
	Valid    bool   `json:"valid"`
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/validate", validate)
	router.Run("localhost:8080")
}

func validate(c *gin.Context) {

	file, err := c.FormFile("file")
	fileExtension := strings.ToLower(file.Filename[len(file.Filename)-4:])

	if err != nil || fileExtension != "json" {
		fmt.Println("Error reading request body:", err)
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// Open the uploaded file
	fileContent, err := file.Open()
	if err != nil {
		fmt.Println("Error opening file:", err)
		c.String(http.StatusInternalServerError, "Internal server error")
		return
	}
	defer fileContent.Close()

	// Read the content of the file
	fileBytes, err := io.ReadAll(fileContent)
	if err != nil {
		fmt.Println("Error reading file:", err)
		c.String(http.StatusInternalServerError, "Internal server error")
		return
	}

	// Decode JSON content into Policy struct
	var policy Policy
	err = json.Unmarshal(fileBytes, &policy)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		c.String(http.StatusBadRequest, "Invalid JSON format")
		return
	}

	fmt.Println(policy.PolicyName)
	fmt.Println(len(policy.PolicyDocument.Statement))

	var statements []Statement = policy.PolicyDocument.Statement
	var valid bool = true
	for _, statement := range statements {
		if regexp.MustCompile(`^[^*]*\*[^*]*$`).MatchString(statement.Resource) {
			valid = false
			break
		}
	}

	fmt.Println("Valid: ", valid)

	var response Response

	if valid {
		response = Response{Filename: file.Filename, Valid: true}
	} else {
		response = Response{Filename: file.Filename, Valid: false}
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.Data(http.StatusOK, "application/json", jsonData)
}
