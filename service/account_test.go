package service

import "testing"

const email = "yuantingwei@pm.me"
const password = "yt"

func TestAccountEmailActivated(t *testing.T) {
	initTest()
	err := AddAccount(email, password)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	result, err := IsAccountEmailVerified(email)
	if result {
		t.Errorf("account " + email + " already verified")
		return
	}

	err = AccountEmailActivated(email)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	result, err = IsAccountEmailVerified(email)
	if result == false {
		t.Errorf("account " + email + " didn't verified")
		return
	}

	err = DeleteAccount(email)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}

func TestGenerateAccountToken(t *testing.T) {
	err := initAuthTest()
	initTest()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	hashedPassword, err := GenerateHashPassword(password)
	err = AddAccount(email, hashedPassword)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	token, err := GenerateAccountToken(email, password)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if len(token) == 0 {
		t.Errorf("Generate ")
		return
	}
	err = DeleteAccount(email)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

}
