package handler

import (
	"net/http"

	"github.com/backend-guchitter-app/util/errors"
	"github.com/backend-guchitter-app/util/utils"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ValidationErr interface {
}

func BadRequest(c *gin.Context, vErr validation.ErrorObject) {
	c.JSON(http.StatusBadRequest, errors.ErrorStruct{
		Message: "Bad request.",
		Fields: []string{
			vErr.Error(),
		},
	})
}

func BadRequests(c *gin.Context, vErr validation.Errors) {
	errorMessages := make([]string, 0)
	for _, e := range vErr {
		errorMessages = append(errorMessages, e.Error())
	}
	c.JSON(http.StatusBadRequest, errors.ErrorStruct{
		Message: "Bad request.",
		Fields:  utils.SortStrings(errorMessages),
	})
}
