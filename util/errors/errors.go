package errors

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

const (
	requiredErrMsgFormat    = "Param '%s' is required."
	invalidTypeErrMsgFormat = "Param '%s' must be a '%s'"
)

//	{
//		message: "",
//		fields: [
//			"Param 'id' is required"
//		]
//	}
type ErrorStruct struct {
	Message string   `json:"message"`
	Fields  []string `json:"fields"`
}

func RequiredErrMsg(field string) string {
	return fmt.Sprintf(requiredErrMsgFormat, field)
}

func InvalidTypeErrMsg(field, expectedType string) string {
	return fmt.Sprintf(invalidTypeErrMsgFormat, field, expectedType)
}

type options struct {
	// tag
	tag string
}

type Option func(o *options)

func Validate(value, fieldName string) []string {
	validate := validator.New()
	validate.RegisterValidation("custom_required", customRequired)

	validateErr := validate.Var(value, "custom_required,number")

	errMessages := make([]string, 0)
	if validateErr == nil {
		return errMessages
	}
	for _, err := range validateErr.(validator.ValidationErrors) {
		typ := err.Tag()
		switch typ {
		case "custom_required":
			errMessages = append(errMessages, RequiredErrMsg(fieldName))
		case "number":
			errMessages = append(errMessages, InvalidTypeErrMsg(fieldName, "number"))
		default:
		}
	}

	return errMessages
}

func customRequired(fl validator.FieldLevel) bool {
	v := fl.Field().String()
	return strings.TrimSpace(v) != ""
}

// dateValidation validates YYYY-mm-DD
func dateValidation(fl validator.FieldLevel) bool {
	_, err := time.Parse("2006-01-02", fl.Field().String())
	return err != nil
}
