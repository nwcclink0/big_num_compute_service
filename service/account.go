package service

func AccountEmailActivated(email string) error {
	account, err := GetAccount(email)
	if err != nil {
		return err
	}
	account.Activated = true
	err = UpdateAccount(account)
	if err != nil {
		return err
	}
	return nil
}

func IsAccountEmailVerified(email string) (bool, error) {
	account, err := GetAccount(email)
	if err != nil {
		return false, err
	}
	activate := account.Activated
	return activate, nil
}

func GenerateAccountToken(email string, password string) (string, error) {
	account, err := GetAccount(email)
	if err != nil {
		return "", err
	}
	success, err := ValidateCredentialsByEmail(account, password)
	if err != nil {
		return "", err
	}
	if success == false {
		return "", ValidateCredentialsByEmailFailed
	}
	token, err := GenerateToken(account)
	return token, nil
}

func VerifiedAccountToken(email string, token string) (bool, error) {
	account, err := GetAccount(email)
	if err != nil {
		return false, err
	}
	validate, err := ValidateToken(token, account)
	if err != nil {
		return false, err
	}
	if validate == false {
		return false, err
	}
	return true, nil
}
