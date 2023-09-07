package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 6XXX error code is defined by boardware cloud.
type ErrorCode int

const (
	UNATHORIZED_ERROR             ErrorCode = 401
	NOT_FOUND_ERROR               ErrorCode = 404
	PERMISSION_ERROR              ErrorCode = 403
	TOO_MANY_REQUESTS             ErrorCode = 429
	EMAIL_EXISTS_ERROR            ErrorCode = 6002
	AUTHENTICATION_ERROR          ErrorCode = 6003
	SERVICE_KEY_DUPLICATION_ERROR ErrorCode = 6004
	VERIFICATION_CODE_ERROR       ErrorCode = 6005
	VERIFICATION_CODE_FREQUENT    ErrorCode = 6006
	UNDEFINE_ERROR                ErrorCode = 9999
)

type Error struct {
	StatusCode int       `json:"statusCode"`
	Code       ErrorCode `json:"code"`
	Message    string    `json:"message"`
}

func (err Error) GinHandler(c *gin.Context) {
	GinErrorHandler(c, err)
}

func GinErrorHandler(c *gin.Context, err Error) {
	c.JSON(
		err.StatusCode,
		gin.H{
			"code":    err.Code,
			"message": err.Message,
		},
	)
}

func VerificationCodeFrequent() *Error {
	return &Error{
		StatusCode: http.StatusForbidden,
		Code:       VERIFICATION_CODE_FREQUENT,
		Message:    "Verification code frequent!",
	}
}

func VerificationCodeError() *Error {
	return &Error{
		StatusCode: http.StatusForbidden,
		Code:       VERIFICATION_CODE_ERROR,
		Message:    "Verification Code error.",
	}
}

func EmailExists() *Error {
	return &Error{
		StatusCode: http.StatusForbidden,
		Code:       EMAIL_EXISTS_ERROR,
		Message:    "Email exists.",
	}
}

func AuthenticationError() *Error {
	return &Error{
		StatusCode: http.StatusUnauthorized,
		Code:       AUTHENTICATION_ERROR,
		Message:    "Account doesn't exists or password fault.",
	}
}

func PermissionError() *Error {
	return &Error{
		StatusCode: http.StatusForbidden,
		Code:       PERMISSION_ERROR,
		Message:    "Permission denied.",
	}
}

func UnauthorizedError() *Error {
	return &Error{
		StatusCode: http.StatusUnauthorized,
		Code:       UNATHORIZED_ERROR,
		Message:    "Unauthorized.",
	}
}

func NotFoundError() *Error {
	return &Error{
		StatusCode: http.StatusNotFound,
		Code:       NOT_FOUND_ERROR,
		Message:    "Not Found 404.",
	}
}

func UndefineError(message string) *Error {
	return &Error{
		StatusCode: http.StatusInternalServerError,
		Code:       UNDEFINE_ERROR,
		Message:    message,
	}
}

func TooManyRequestsError() *Error {
	return &Error{
		StatusCode: http.StatusTooManyRequests,
		Code:       TOO_MANY_REQUESTS,
		Message:    "Too many requests.",
	}
}

func (e Error) Error() string {
	return e.Message
}
