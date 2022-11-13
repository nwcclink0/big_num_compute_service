package service

import (
	"fmt"
	"github.com/joho/godotenv"
	"testing"
)

func initAuthTest() error {
	Argon2Conf = &Argon2Params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
	err := godotenv.Load(".env_test")
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func TestGenerateHashPasswordAndCheckHashPassword(t *testing.T) {
	err := initAuthTest()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	hashPassword, err := GenerateHashPassword(password)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(password + "'s hash password is: " + hashPassword)
	success, err := compareHashPasswordAndPassword(hashPassword, password)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if success == false {
		t.Errorf("can't compare " + password + " with " + hashPassword)
		return
	}
}

func TestVerifyAccountEmail(t *testing.T) {
	passcode, err := GenerateTotp(email)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	verified, err := VerifyAccountEmail(email, passcode)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if verified == false {
		t.Errorf("email: " + email + " verified failed")
		return
	}
}

func TestGenerateToken(t *testing.T) {
	initTest()
	initAuthTest()
	hashedPassword, err := GenerateHashPassword(password)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	err = AddAccount(email, hashedPassword)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	account, err := GetAccount(email)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	token, err := GenerateToken(account)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	validate, err := ValidateToken(token, account)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if validate == false {
		t.Errorf(err.Error())
		return
	}

	err = DeleteAccount(email)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
