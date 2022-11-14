package service

import (
	"big_num_compute_service/rpc"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"os"

	//"net/rpc/jsonrpc"
	"strconv"
)

type BigNumCompute struct {
}

func (BigNumCompute) Create(args []string, result *string) error {
	if len(args) != 4 {
		return ArgumentNumberCreateFailed
	}

	name := args[0]
	if len(name) == 0 {
		return NameEmptyFailed
	}
	LogAccess.Debug("create number object name: " + name)
	numberStr := args[1]
	if len(numberStr) == 0 {
		return NumberEmptyFailed
	}
	number, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		return NumberParsingFailed
	}

	email := args[2]
	if len(email) == 0 {
		return EmailEmptyFailed
	}
	if validateEmailFormat(email) == false {
		return EmailFormatFailed
	}

	token := args[3]
	if len(token) == 0 {
		return TokenEmptyFailed
	}

	validate, err := VerifiedAccountToken(email, token)
	if err != nil {
		return ValidateTokenFailed
	}
	if validate == false {
		return ValidateTokenFailed
	}

	LogAccess.Debug("create number object number: ", number)

	err = CreateObj(name, number)
	if err != nil {
		return fmt.Errorf("create number object failed: %s", err)
	}

	*result = ResultSuccess
	return nil
}

func (BigNumCompute) Delete(args []string, result *string) error {
	if len(args) != 3 {
		return ArgumentNumberDeleteFailed
	}
	name := args[0]
	LogAccess.Debug("delete object name")

	email := args[1]
	if len(email) == 0 {
		return EmailEmptyFailed
	}
	if validateEmailFormat(email) == false {
		return EmailFormatFailed
	}

	token := args[2]
	if len(token) == 0 {
		return TokenEmptyFailed
	}

	validate, err := VerifiedAccountToken(email, token)
	if err != nil {
		return ValidateTokenFailed
	}
	if validate == false {
		return ValidateTokenFailed
	}

	err = DeleteObj(name)
	if err != nil {
		return fmt.Errorf("delete number object failed: %s", err)
	}
	*result = ResultSuccess
	return nil
}

func (BigNumCompute) Update(args []string, result *string) error {
	if len(args) != 4 {
		return ArgumentNumberUpdateFailed
	}
	name := args[0]
	LogAccess.Debug("create number object name: " + name)
	numberStr := args[1]
	number, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		return NumberParsingFailed
	}

	email := args[2]
	if len(email) == 0 {
		return EmailEmptyFailed
	}
	if validateEmailFormat(email) == false {
		return EmailFormatFailed
	}

	token := args[3]
	if len(token) == 0 {
		return TokenEmptyFailed
	}

	validate, err := VerifiedAccountToken(email, token)
	if err != nil {
		return ValidateTokenFailed
	}
	if validate == false {
		return ValidateTokenFailed
	}

	err = UpdateObj(name, number)
	if err != nil {
		return fmt.Errorf("can't update %s with %s, error: %s", args[0], args[1], err)
	}
	*result = ResultSuccess
	return nil
}

func (BigNumCompute) Add(args []string, result *string) error {
	if len(args) != 4 {
		return ArgumentNumberAddFailed
	}

	email := args[2]
	if len(email) == 0 {
		return EmailEmptyFailed
	}
	if validateEmailFormat(email) == false {
		return EmailFormatFailed
	}

	token := args[3]
	if len(token) == 0 {
		return TokenEmptyFailed
	}

	validate, err := VerifiedAccountToken(email, token)
	if err != nil {
		return ValidateTokenFailed
	}
	if validate == false {
		return ValidateTokenFailed
	}

	computeResult, err := Compute(args[0], args[1], AddOp)
	if err != nil {
		return fmt.Errorf("can't multiply of %s and %s, error: %s", args[0], args[1], err)
	}
	*result = fmt.Sprint(computeResult)
	return nil
}

func (BigNumCompute) Subtract(args []string, result *string) error {
	if len(args) != 4 {
		return ArgumentNumberSubtractFailed
	}

	email := args[2]
	if len(email) == 0 {
		return EmailEmptyFailed
	}
	if validateEmailFormat(email) == false {
		return EmailFormatFailed
	}

	token := args[3]
	if len(token) == 0 {
		return TokenEmptyFailed
	}

	validate, err := VerifiedAccountToken(email, token)
	if err != nil {
		return ValidateTokenFailed
	}
	if validate == false {
		return ValidateTokenFailed
	}

	computeResult, err := Compute(args[0], args[1], SubtractOp)
	if err != nil {
		return fmt.Errorf("can't multiply of %s and %s, error: %s", args[0], args[1], err)
	}
	*result = fmt.Sprint(computeResult)
	return nil
}

func (BigNumCompute) Multiply(args []string, result *string) error {
	if len(args) != 4 {
		return ArgumentNumberMultiplyFailed
	}

	email := args[2]
	if len(email) == 0 {
		return EmailEmptyFailed
	}
	if validateEmailFormat(email) == false {
		return EmailFormatFailed
	}

	token := args[3]
	if len(token) == 0 {
		return TokenEmptyFailed
	}

	validate, err := VerifiedAccountToken(email, token)
	if err != nil {
		return ValidateTokenFailed
	}
	if validate == false {
		return ValidateTokenFailed
	}

	computeResult, err := Compute(args[0], args[1], MultiplyOp)
	if err != nil {
		return fmt.Errorf("can't multiply of %s and %s, error: %s", args[0], args[1], err)
	}
	*result = fmt.Sprint(computeResult)
	return nil
}

func (BigNumCompute) Divide(args []string, result *string) error {
	if len(args) != 4 {
		return ArgumentNumberDivideFailed
	}

	email := args[2]
	if len(email) == 0 {
		return EmailEmptyFailed
	}
	if validateEmailFormat(email) == false {
		return EmailFormatFailed
	}

	token := args[3]
	if len(token) == 0 {
		return TokenEmptyFailed
	}

	validate, err := VerifiedAccountToken(email, token)
	if err != nil {
		return ValidateTokenFailed
	}
	if validate == false {
		return ValidateTokenFailed
	}

	computeResult, err := Compute(args[0], args[1], DivideOp)
	if err != nil {
		return fmt.Errorf("can't multiply of %s and %s, error: %s", args[0], args[1], err)
	}
	*result = fmt.Sprint(computeResult)
	return nil
}

func sendOptMail(email string, passcode string) error {
	from := os.Getenv("MAIL_ACCOUNT")
	if len(from) == 0 {
		return fmt.Errorf("mail sender failed")
	}
	password := os.Getenv("MAIL_AUTH")
	if len(password) == 0 {
		return fmt.Errorf("mail sender failed")
	}

	toEmailAddress := email
	to := []string{toEmailAddress}

	host := os.Getenv("MAIL_SMTP_HOST")
	port := os.Getenv("MAIL_SMTP_PORT")
	address := host + ":" + port

	subject := "Account verify code\r\n"
	body := "Verify code is " + passcode + ". It will expire in 30 seconds"
	//contentType := "Content-Type: text/html; charset=UTF-8"
	message := []byte(fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s",
		toEmailAddress, "yt", from, subject, body))

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}
	return nil
}

func (BigNumCompute) CreateAccount(args []string, result *string) error {
	if len(args) != 2 {
		return fmt.Errorf("arguments shoud be [email:password]")
	}

	email := args[0]
	if len(email) == 0 {
		return fmt.Errorf("email is empty")
	}
	if validateEmailFormat(email) == false {
		return EmailFormatFailed
	}
	password := args[1]
	if len(password) == 0 {
		return fmt.Errorf("password is empty")
	}

	hashedPassword, err := GenerateHashPassword(password)
	if err != nil {
		return err
	}
	err = AddAccount(email, string(hashedPassword))
	if err != nil {
		return err
	}

	passcode, err := GenerateTotp(email)
	if err != nil {
		return err
	}
	*result = passcode

	err = sendOptMail(email, passcode)
	if err != nil {
		LogError.Error("send email error: ", err)
	}
	return nil
}

func validateEmailFormat(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (BigNumCompute) ValidateEmail(args []string, result *string) error {
	if len(args) != 2 {
		return ArgumentValidateEmailFailed
	}
	email := args[0]
	if len(email) == 0 {
		return EmailEmptyFailed
	}
	if validateEmailFormat(email) == false {
		return EmailFormatFailed
	}

	passcode := args[1]
	if len(passcode) == 0 {
		return EmptyPassCode
	}

	success, err := VerifyAccountEmail(email, passcode)
	if err != nil {
		return EmailVerifyFailed
	}
	if success == false {
		return EmailVerifyFailed
	}
	err = AccountEmailActivated(email)
	if err != nil {
		return EmailActivateFailed
	}
	*result = ResultSuccess
	return nil
}

func (BigNumCompute) LoginAccount(args []string, result *string) error {
	if len(args) != 2 {
		return ArgumentEmailPasswordFailed
	}
	email := args[0]
	if len(email) == 0 {
		return EmailEmptyFailed
	}
	if validateEmailFormat(email) == false {
		return EmailFormatFailed
	}

	password := args[1]
	if len(password) == 0 {
		return EmptyPasswordFailed
	}

	success, err := IsAccountEmailVerified(email)
	if err != nil {
		return EmailNotVerified
	}
	if success == false {
		return EmailNotVerified
	}
	token, err := GenerateAccountToken(email, password)
	if err != nil {
		return GenerateAccountTokenFailed
	}
	*result = token
	return nil
}

func (BigNumCompute) DeleteAccount(args []string, result *string) error {
	if len(args) != 2 {
		return ArgumentEmailTokenFailed
	}
	email := args[0]
	if len(email) == 0 {
		return EmailEmptyFailed
	}
	if validateEmailFormat(email) == false {
		return EmailFormatFailed
	}

	token := args[1]
	if len(token) == 0 {
		return TokenEmptyFailed
	}

	success, err := VerifiedAccountToken(email, token)
	if err != nil {
		return ValidateTokenFailed
	}
	if success == false {
		return ValidateTokenFailed
	}

	err = DeleteAccount(email)
	if err != nil {
		return DeleteAccountFailed
	}
	*result = ResultSuccess
	return nil
}

func Run() {
	err := rpc.Register(BigNumCompute{})
	if err != nil {
		LogError.Error("register big num error, ", err.Error())
		return
	}
	listen, err := net.Listen("tcp", ":"+BigNumComputeConf.Core.Port)
	if err != nil {
		LogError.Error("can't listen port: " + BigNumComputeConf.Core.Port + ", err:" + err.Error())
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			LogError.Error("accept error: " + err.Error())
		}
		QueueComputeWorker <- conn
	}
}
