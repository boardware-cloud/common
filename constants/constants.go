package constants

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
	HTTP MonitorType   = "HTTP"
	PING MonitorStatus = "PING"
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
