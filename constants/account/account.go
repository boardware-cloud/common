package account

import (
	"database/sql/driver"

	errorCode "github.com/boardware-cloud/common/code"
)

// Account type
type Role string

const (
	ROOT  Role = "ROOT"
	ADMIN Role = "ADMIN"
	USER  Role = "USER"
)

func ToRole(role string) (Role, error) {
	r := Role(role)
	switch r {
	case ROOT, ADMIN, USER:
		return r, nil
	}
	return r, errorCode.ErrEnum
}

func (Role) GormDataType() string {
	return "VARCHAR(128)"
}

func (r *Role) Scan(value any) error {
	*r = Role(string(value.([]byte)))
	return nil
}

func (s Role) Value() (driver.Value, error) {
	return string(s), nil
}

type TokenType string

const (
	JWT TokenFormat = "JWT"
)

type TokenFormat string

const (
	BEARER TokenType = "BEARER"
)
