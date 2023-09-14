package authenication

import errorCode "github.com/boardware-cloud/common/code"

type AuthenticationFactor string

const (
	PASSWORD AuthenticationFactor = "PASSWORD"
	TOTP     AuthenticationFactor = "TOTP"
	WEBAUTHN AuthenticationFactor = "WEBAUTHN"
	EMAIL    AuthenticationFactor = "EMAIL"
)

func (AuthenticationFactor) GormDataType() string {
	return "VARCHAR(128)"
}

func ToAuthenticationFactor(s string) (AuthenticationFactor, error) {
	factor := AuthenticationFactor(s)
	switch factor {
	case PASSWORD, TOTP, WEBAUTHN, EMAIL:
		return factor, nil
	}
	return factor, errorCode.ErrEnum
}
