package code

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrUnauthorized = errors.New("unauthorized")
var ErrPermissionDenied = errors.New("permission denied")
var ErrNotFound = errors.New("not found")
var ErrTooManyRequests = errors.New("too many requests")
var ErrEmailExists = errors.New("email exists")
var ErrUndefined = errors.New("undefined error")
var ErrVerificationCode = errors.New("verification code error")
var ErrEnum = errors.New("enum parse error")

type ErrorCode int

const (
	UNATHORIZED_ERROR             ErrorCode = 401
	PERMISSION_ERROR              ErrorCode = 403
	NOT_FOUND_ERROR               ErrorCode = 404
	TOO_MANY_REQUESTS             ErrorCode = 429
	EMAIL_EXISTS_ERROR            ErrorCode = 6002
	AUTHENTICATION_ERROR          ErrorCode = 6003
	SERVICE_KEY_DUPLICATION_ERROR ErrorCode = 6004
	VERIFICATION_CODE_ERROR       ErrorCode = 6005
	VERIFICATION_CODE_FREQUENT    ErrorCode = 6006
	UNDEFINE_ERROR                ErrorCode = 9999
)

func GetCode(err error) (statusCode int, errorCode ErrorCode) {
	switch err {
	case ErrUnauthorized:
		statusCode = http.StatusUnauthorized
		errorCode = UNATHORIZED_ERROR
	case ErrPermissionDenied:
		statusCode = http.StatusForbidden
		errorCode = PERMISSION_ERROR
	case ErrNotFound:
		statusCode = http.StatusNotFound
		errorCode = NOT_FOUND_ERROR
	case ErrTooManyRequests:
		statusCode = http.StatusTooManyRequests
		errorCode = TOO_MANY_REQUESTS
	case ErrEmailExists:
		statusCode = http.StatusForbidden
		errorCode = EMAIL_EXISTS_ERROR
	case ErrUndefined:
		statusCode = http.StatusInternalServerError
		errorCode = UNDEFINE_ERROR
	}
	return statusCode, errorCode
}

func GinHandler(ctx *gin.Context, err error) {
	statusCode, errorCode := GetCode(err)
	ctx.JSON(statusCode, gin.H{
		"code":    errorCode,
		"message": err.Error(),
	})
}
