package account

import "github.com/boardware-cloud/common/code"

type VerificationCodePurpose string

const (
	CREATE_ACCOUNT VerificationCodePurpose = "CREATE_ACCOUNT"
	SET_PASSWORD   VerificationCodePurpose = "SET_PASSWORD"
	CREATE_2FA     VerificationCodePurpose = "CREATE_2FA"
	SIGNIN         VerificationCodePurpose = "SIGNIN"
	TICKET         VerificationCodePurpose = "TICKET"
)

func ToVerificationCodePurpose(s string) (VerificationCodePurpose, error) {
	r := VerificationCodePurpose(s)
	switch r {
	case CREATE_ACCOUNT, SET_PASSWORD, CREATE_2FA, SIGNIN, TICKET:
		return r, nil
	}
	return r, code.ErrEnum
}
