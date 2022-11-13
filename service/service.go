package service

import (
	"big_num_compute_service/rpc"
	"fmt"
	"net"
	"net/smtp"
	"os"

	//"net/rpc/jsonrpc"
	"strconv"
)

type BigNumCompute struct {
}

func (BigNumCompute) Create(args []string, result *string) error {
	if len(args) != 4 {
		return fmt.Errorf("argument length shoue be 2 with object name and related number")
	}

	name := args[0]
	if len(name) == 0 {
		return fmt.Errorf("name is empty")
	}
	LogAccess.Debug("create number object name: " + name)
	numberStr := args[1]
	if len(numberStr) == 0 {
		return fmt.Errorf("naumber is empty")
	}
	number, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		return fmt.Errorf("parsing number failed: %s", err.Error())
	}

	email := args[2]
	if len(email) == 0 {
		return fmt.Errorf("email is empty")
	}

	token := args[3]
	if len(token) == 0 {
		return fmt.Errorf("token is empty")
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
		return fmt.Errorf("argument should be a name that want to delete with")
	}
	name := args[0]
	LogAccess.Debug("delete object name")

	email := args[1]
	if len(email) == 0 {
		return fmt.Errorf("email is empty")
	}

	token := args[2]
	if len(token) == 0 {
		return fmt.Errorf("token is empty")
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
		return fmt.Errorf("argument length error, it should be [name:number]")
	}
	name := args[0]
	LogAccess.Debug("create number object name: " + name)
	numberStr := args[1]
	number, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		return fmt.Errorf("parsing number failed: %s", err)
	}

	email := args[2]
	if len(email) == 0 {
		return fmt.Errorf("email is empty")
	}

	token := args[3]
	if len(token) == 0 {
		return fmt.Errorf("token is empty")
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
		return fmt.Errorf("argument length error, it should be [name:number] or [name1:name2]")
	}

	email := args[2]
	if len(email) == 0 {
		return fmt.Errorf("email is empty")
	}

	token := args[3]
	if len(token) == 0 {
		return fmt.Errorf("token is empty")
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
		return fmt.Errorf("argument length error, it should be [name:number] or [name1:name2]")
	}

	email := args[2]
	if len(email) == 0 {
		return fmt.Errorf("email is empty")
	}

	token := args[3]
	if len(token) == 0 {
		return fmt.Errorf("token is empty")
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
		return fmt.Errorf("argument length error, it should be [name:number] or [name1:name2]")
	}

	email := args[2]
	if len(email) == 0 {
		return fmt.Errorf("email is empty")
	}

	token := args[3]
	if len(token) == 0 {
		return fmt.Errorf("token is empty")
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
		return fmt.Errorf("argument length error, it should be [name:number] or [name1:name2]")
	}

	email := args[2]
	if len(email) == 0 {
		return fmt.Errorf("email is empty")
	}

	token := args[3]
	if len(token) == 0 {
		return fmt.Errorf("token is empty")
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
	from := os.Getenv("GMAIL_ACCOUNT")
	if len(from) == 0 {
		return fmt.Errorf("mail sender failed")
	}
	password := os.Getenv("GMAIL_AUTH")
	if len(password) == 0 {
		return fmt.Errorf("mail sender failed")
	}

	toEmailAddress := email
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
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

func (BigNumCompute) ValidateEmail(args []string, result *string) error {
	if len(args) != 2 {
		return fmt.Errorf("arguments should be [email:passcode]")
	}
	email := args[0]
	passcode := args[1]
	success, err := VerifyAccountEmail(email, passcode)
	if err != nil {
		return err
	}
	if success == false {
		return err
	}
	err = AccountEmailActivated(email)
	if err != nil {
		return err
	}
	*result = ResultSuccess
	return nil
}

func (BigNumCompute) LoginAccount(args []string, result *string) error {
	if len(args) != 2 {
		return ArgumentEmailPassword
	}
	email := args[0]
	password := args[1]

	success, err := IsAccountEmailVerified(email)
	if err != nil {
		return err
	}
	if success == false {
		return EmailNotVerified
	}
	token, err := GenerateAccountToken(email, password)
	if err != nil {
		return err
	}
	*result = token
	return nil
}

func (BigNumCompute) DeleteAccount(args []string, result *string) error {
	if len(args) != 2 {
		return ArgumentEmailToken
	}
	email := args[0]
	token := args[1]

	success, err := VerifiedAccountToken(email, token)
	if err != nil {
		return err
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
