package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestValidJsons(t *testing.T) {
	router := gin.New()
	router.POST("/validate", validate)

	folderPath := "../test-files/valid-jsons"
	files, err := os.ReadDir(folderPath)
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {

		filePath := filepath.Join(folderPath, file.Name())
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			t.Fatalf("Failed to read file: %v", err)
		}

		var requestBody bytes.Buffer
		multipartWriter := multipart.NewWriter(&requestBody)

		fileWriter, err := multipartWriter.CreateFormFile("file", filepath.Base(filePath))
		if err != nil {
			t.Fatalf("Failed to create form file: %v", err)
		}
		if _, err := fileWriter.Write(fileContent); err != nil {
			t.Fatalf("Failed to write file content: %v", err)
		}

		if err := multipartWriter.Close(); err != nil {
			t.Fatalf("Failed to close multipart writer: %v", err)
		}

		request, err := http.NewRequest("POST", "/validate", &requestBody)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

		filename := filepath.Base(filePath)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code, "Expected HTTP status 200")

		message := fmt.Sprintf("JSON from file %s is VALID", filename)
		expectedData := Response{
			Valid:   true,
			Message: message,
		}

		expectedJSON, err := json.Marshal(expectedData)
		if err != nil {
			t.Fatalf("Failed to marshal expected data to JSON: %v", err)
		}

		assert.Equal(t, string(expectedJSON), response.Body.String(), "Expected JSON data")
	}
}

// test to be finished
func TestNotValidJsons(t *testing.T) {
	router := gin.New()
	router.POST("/validate", validate)

	folderPath := "../test-files/not-valid-jsons"
	files, err := os.ReadDir(folderPath)
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {

		filePath := filepath.Join(folderPath, file.Name())
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			t.Fatalf("Failed to read file: %v", err)
		}

		var requestBody bytes.Buffer
		multipartWriter := multipart.NewWriter(&requestBody)

		fileWriter, err := multipartWriter.CreateFormFile("file", filepath.Base(filePath))
		if err != nil {
			t.Fatalf("Failed to create form file: %v", err)
		}
		if _, err := fileWriter.Write(fileContent); err != nil {
			t.Fatalf("Failed to write file content: %v", err)
		}

		if err := multipartWriter.Close(); err != nil {
			t.Fatalf("Failed to close multipart writer: %v", err)
		}

		request, err := http.NewRequest("POST", "/validate", &requestBody)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

		filename := filepath.Base(filePath)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code, "Expected HTTP status 200")

		message := fmt.Sprintf("JSON from file %s is NOT VALID", filename)
		expectedData := Response{
			Valid:   true,
			Message: message,
		}

		expectedJSON, err := json.Marshal(expectedData)
		if err != nil {
			t.Fatalf("Failed to marshal expected data to JSON: %v", err)
		}

		assert.Equal(t, string(expectedJSON), response.Body.String(), "Expected JSON data")
	}
}

func TestInvalidJsonFormat(t *testing.T) {
	router := gin.New()
	router.POST("/validate", validate)

	folderPath := "../test-files/invalid-format-jsons"
	files, err := os.ReadDir(folderPath)
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {

		filePath := filepath.Join(folderPath, file.Name())
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			t.Fatalf("Failed to read file: %v", err)
		}

		var requestBody bytes.Buffer
		multipartWriter := multipart.NewWriter(&requestBody)

		fileWriter, err := multipartWriter.CreateFormFile("file", filepath.Base(filePath))
		if err != nil {
			t.Fatalf("Failed to create form file: %v", err)
		}
		if _, err := fileWriter.Write(fileContent); err != nil {
			t.Fatalf("Failed to write file content: %v", err)
		}

		if err := multipartWriter.Close(); err != nil {
			t.Fatalf("Failed to close multipart writer: %v", err)
		}

		request, err := http.NewRequest("POST", "/validate", &requestBody)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

		filename := filepath.Base(filePath)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code, "Expected HTTP status 200")

		message := fmt.Sprintf("JSON from file %s is VALID", filename)
		expectedData := Response{
			Valid:   true,
			Message: message,
		}

		expectedJSON, err := json.Marshal(expectedData)
		if err != nil {
			t.Fatalf("Failed to marshal expected data to JSON: %v", err)
		}

		assert.Equal(t, string(expectedJSON), response.Body.String(), "Expected JSON data")
	}
}

func TestUnsupportedMediaType(t *testing.T) {
	router := gin.New()
	router.POST("/validate", validate)

	filePath := "../test-files/test.txt"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)

	fileWriter, err := multipartWriter.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}
	if _, err := fileWriter.Write(fileContent); err != nil {
		t.Fatalf("Failed to write file content: %v", err)
	}

	if err := multipartWriter.Close(); err != nil {
		t.Fatalf("Failed to close multipart writer: %v", err)
	}

	request, err := http.NewRequest("POST", "/validate", &requestBody)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusUnsupportedMediaType, response.Code, "Expected HTTP status 415")

}

func TestBadRequest(t *testing.T) {
	router := gin.New()
	router.POST("/validate", validate)

	testData := []byte(`{"key": "value"}`)
	request, _ := http.NewRequest("POST", "/validate", bytes.NewBuffer(testData))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code, "Expected HTTP status 400")

	request, _ = http.NewRequest("GET", "/validate", bytes.NewBuffer(testData))
	request.Header.Set("Content-Type", "application/json")

	response = httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusNotFound, response.Code, "Expected HTTP status 404")

}

func TestNotFound(t *testing.T) {
	router := gin.New()
	router.POST("/validate", validate)

	filePath := "../test-files/valid-jsons/test1.json"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)

	fileWriter, err := multipartWriter.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}
	if _, err := fileWriter.Write(fileContent); err != nil {
		t.Fatalf("Failed to write file content: %v", err)
	}

	if err := multipartWriter.Close(); err != nil {
		t.Fatalf("Failed to close multipart writer: %v", err)
	}

	request, err := http.NewRequest("GET", "/validate", &requestBody)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusNotFound, response.Code, "Expected HTTP status 404")
}
