package service

import "errors"

const (
	AddOp      = "add"
	SubtractOp = "subtract"
	MultiplyOp = "multiply"
	DivideOp   = "divide"
)

const (
	ResultSuccess = "success"
	ResultFailed  = "failed"
	ResultZero    = "0"
)

const (
	Docker    = "docker"
	Localhost = "localhost"
)

const (
	Issuer = "example.com"
)

var ArgumentEmailPassword = errors.New("arguments should be [email:password")
var ArgumentEmailToken = errors.New("arguments should be [email:token")
var EmailNotVerified = errors.New("email not verified yet")
var ValidateCredentialsByEmailFailed = errors.New("email with credential failed")
var ValidatePasscodeFailed = errors.New("validate passcode failed")
var ValidateTokenFailed = errors.New("validate token failed")
var DeleteAccountFailed = errors.New("delete account failed")
