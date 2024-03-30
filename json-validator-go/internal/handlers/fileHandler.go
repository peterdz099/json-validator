package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"

	"json-validator/internal/messages"

	"github.com/gin-gonic/gin"
)

func ReadFile(c *gin.Context) (Policy, string, error) {
	file, err := c.FormFile("file")

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return Policy{}, "", err
	}

	if file != nil {
		if fileExtension := strings.ToLower(file.Filename[len(file.Filename)-4:]); fileExtension != "json" {
			c.String(http.StatusUnsupportedMediaType, messages.UNSUPPORTED_FILE_ERR)
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
			c.String(http.StatusBadRequest, messages.INVALID_TYPE_ERR)
		} else {
			c.String(http.StatusBadRequest, messages.INVALID_FORMAT_ERR)
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
