package errors

import (
	"fmt"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	// "strings"
	// "time"
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

func ValidateNotEmpty(fieldName string) validation.RuleFunc {
	return func(value interface{}) error {

		// 数値の場合
		if _, ok := value.(int); ok {
			return nil
		}

		// string型の場合
		_, ok := value.(string)
		if !ok {
			return validation.NewError("InvalidType", "Invalid type")
		}

		if strings.TrimSpace(value.(string)) == "" {
			return validation.NewError("Required", RequiredErrMsg(fieldName))
		}
		return nil
	}
}

func ValidateYYYY_MM_DD(fieldName string) validation.RuleFunc {
	return func(value interface{}) error {

		dateStr, ok := value.(string)
		if !ok {
			return validation.NewError("InvalidType", "Invalid type")
		}

		if strings.TrimSpace(dateStr) == "" {
			return nil
		}

		// 日付の解析を試みる
		_, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return validation.NewError("InvalidDate", InvalidTypeErrMsg(fieldName, "YYYY-MM-DD"))
		}

		return nil
	}
}
