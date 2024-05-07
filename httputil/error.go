package httputil

import "github.com/go-playground/validator/v10"

type ErrorCode string

const (
	BadRequest          ErrorCode = "0001"
	NotFound            ErrorCode = "0002"
	CreateUserError     ErrorCode = "0003"
	StatusForbidden     ErrorCode = "0004"
	StatusUnauthorized  ErrorCode = "0005"
	InternalServerError ErrorCode = "0006"
	InvalidStatus       ErrorCode = "9999"
)

var InternalServerErrorMessage string = "ระบบ REG เชื่อมต่อไม่ได้ในขณะนี้"
var InternalServerErrorMessageEn string = "REG system is currently unable to connect"

type HTTPBadRequestErrors struct {
	Code    ErrorCode             `json:"code,omitempty"`
	Message string                `json:"message,omitempty"`
	Errors  []HTTPBadRequestError `json:"errors"`
}

type HTTPBadRequestError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type HTTPError struct {
	Code      ErrorCode `json:"code"`
	Message   string    `json:"message"`
	MessageEn string    `json:"messageEn,omitempty"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "required_if":
		return "This field is required"
	case "excluded_unless":
		return "This field is excluded"
	case "gtefield":
		return "Should be greater than or equal " + fe.Param()
	case "oneof":
		return "Should be one of " + fe.Param()
	case "ipv4":
		return "Should be ipv4"
	}
	return "Unknown error"
}
