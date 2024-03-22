package main

//https://stackoverflow.com/questions/76448408/cors-policy-response-to-preflight-request-doesnt-pass-access-control-check-no
//https://stackoverflow.com/questions/76448408/cors-policy-response-to-preflight-request-doesnt-pass-access-control-check-no
//https://en.wikipedia.org/wiki/Cross-origin_resource_sharing

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("validate", validate)
	router.Run("localhost:8080")
}

// postAlbums adds an album from JSON received in the request body.
func validate(c *gin.Context) {
	// Read the request body
	body, err := c.GetRawData()
	if err != nil {
		fmt.Println("Error reading request body:", err)
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	// Print the request body to console
	fmt.Println("Received JSON:", string(body))

	// You can handle the JSON data here as needed

	c.String(http.StatusOK, "Request received")
}
