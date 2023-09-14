package service

import "database/sql/driver"

type ServiceName string

func (ServiceName) GormDataType() string {
	return "VARCHAR(128)"
}

func (s *ServiceName) Scan(value any) error {
	*s = ServiceName(string(value.([]byte)))
	return nil
}

func (s ServiceName) Value() (driver.Value, error) {
	return string(s), nil
}

const (
	ARGUS ServiceName = "ARGUS"
)

// Service type
type ServiceType string

const (
	SAAS ServiceType = "SAAS"
	PAAS ServiceType = "PAAS"
	IAAS ServiceType = "IAAS"
)
