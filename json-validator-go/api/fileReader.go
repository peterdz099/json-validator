package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

func readFile(c *gin.Context) (Policy, string, error) {
	file, err := c.FormFile("file")

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return Policy{}, "", err
	}

	if file != nil {
		if fileExtension := strings.ToLower(file.Filename[len(file.Filename)-4:]); fileExtension != "json" {
			c.String(http.StatusUnsupportedMediaType, "Wrong file extension")
			return Policy{}, "", err
		}
	}

	fileContent, err := file.Open()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return Policy{}, "", err
	}
	defer fileContent.Close()

	fileBytes, err := io.ReadAll(fileContent)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return Policy{}, "", err
	}

	var policy Policy
	decoder := json.NewDecoder(bytes.NewReader(fileBytes))
	decoder.DisallowUnknownFields()
	decodeErr := decoder.Decode(&policy)

	if decodeErr != nil {
		if reflect.TypeOf(decodeErr) == reflect.TypeOf(&json.UnmarshalTypeError{}) {
			c.String(http.StatusBadRequest, "invalid type: found invalid field type")
		} else {
			c.String(http.StatusBadRequest, "invalid format: invalid JSON format")
		}
		return Policy{}, "", decodeErr
	}

	isFormatValid, err := verifyJsonFormat(policy)
	if !isFormatValid && err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return Policy{}, "", err
	}
	return policy, file.Filename, nil
}
