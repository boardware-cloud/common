package constants

import (
	"database/sql/driver"
	"fmt"

	"github.com/boardware-cloud/common/code"
)

// Service type
type ServiceType string

const (
	SAAS ServiceType = "SAAS"
	PAAS ServiceType = "PAAS"
	IAAS ServiceType = "IAAS"
)

// Account type
type Role string

const (
	ROOT  Role = "ROOT"
	ADMIN Role = "ADMIN"
	USER  Role = "USER"
)

func RoleValidate(role string) bool {
	switch Role(role) {
	case ROOT, ADMIN, USER:
		return true
	}
	return false
}

func (Role) GormDataType() string {
	return "VARCHAR(128)"
}

func (s *Role) Scan(value any) error {
	fmt.Println(value)
	if RoleValidate(value.(string)) {
		return nil
	}
	return code.ErrEnum
}

func (s Role) Value() (driver.Value, error) {
	return []byte(s), nil
}

type TokenType string

const (
	JWT TokenFormat = "JWT"
)

type TokenFormat string

const (
	BEARER TokenType = "BEARER"
)

type HttpMehotd string

const (
	HEAD    HttpMehotd = "HEAD"
	POST    HttpMehotd = "POST"
	GET     HttpMehotd = "GET"
	PUT     HttpMehotd = "PUT"
	PATCH   HttpMehotd = "PATCH"
	DELETE  HttpMehotd = "DELETE"
	OPTIONS HttpMehotd = "OPTIONS"
)

type MonitorType string

const (
	HTTP MonitorType = "HTTP"
)

type MonitoringResult string

const (
	OK      MonitoringResult = "OK"
	TIMEOUT MonitoringResult = "TIMEOUT"
	DOWN    MonitoringResult = "DOWN"
)

type MonitorStatus string

const (
	ACTIVED    MonitorStatus = "ACTIVED"
	DISACTIVED MonitorStatus = "DISACTIVED"
)

type NotificationType string

const (
	EMAIL NotificationType = "EMAIL"
)

type VerificationCodePurpose string

const (
	CREATE_ACCOUNT VerificationCodePurpose = "CREATE_ACCOUNT"
	SET_PASSWORD   VerificationCodePurpose = "SET_PASSWORD"
	CREATE_2FA     VerificationCodePurpose = "CREATE_2FA"
	SIGNIN         VerificationCodePurpose = "SIGNIN"
	TICKET         VerificationCodePurpose = "TICKET"
)

type CreateVerificationResult string

const (
	ACCOUNT_EXISTS  CreateVerificationResult = "ACCOUNT_EXISTS"
	SUCCESS_CREATED CreateVerificationResult = "SUCCESS_CREATED"
	FREQUENT        CreateVerificationResult = "FREQUENT"
)

type HttpBodyForm string

const (
	RAW                   HttpBodyForm = "RAW"
	X_WWW_FORM_URLENCODED HttpBodyForm = "X_WWW_FORM_URLENCODED"
)

func (h HttpBodyForm) Request() string {
	switch h {
	case RAW:
		return "raw"
	case X_WWW_FORM_URLENCODED:
		return "x-www-form-urlencoded"
	}
	return ""
}

type ContentType string

const (
	JSON ContentType = "JSON"
	XML  ContentType = "XML"
)

func (c ContentType) Request() string {
	switch c {
	case JSON:
		return "application/json"
	case XML:
		return "application/xml"
	}
	return ""
}
