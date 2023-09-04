package authenication

import "errors"

type AuthenticationFactor string

const (
	PASSWORD AuthenticationFactor = "PASSWORD"
	TOTP     AuthenticationFactor = "TOTP"
	WEBAUTHN AuthenticationFactor = "WEBAUTHN"
	EMAIL    AuthenticationFactor = "EMAIL"
)

func ToAuthenticationFactor(s string) (AuthenticationFactor, error) {
	factor := AuthenticationFactor(s)
	if factor != PASSWORD && factor != TOTP && factor != WEBAUTHN && factor != EMAIL {
		return "", errors.New("enum error")
	}
	return factor, nil
}
