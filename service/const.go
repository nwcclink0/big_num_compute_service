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

var ArgumentEmailPasswordFailed = errors.New("arguments should be [email:password")
var ArgumentEmailTokenFailed = errors.New("arguments should be [email:token")
var ArgumentNumberCreateFailed = errors.New("argument length should be 4 with object [name, number, email, account_token]")
var ArgumentNumberDeleteFailed = errors.New("argument length should be 3 with object [name, email, account_token]")
var ArgumentNumberUpdateFailed = errors.New("argument length should be 4 with object [name, number, email, account_token]")
var ArgumentNumberAddFailed = errors.New("argument length should be 4 with object [name, number/name, email, account_token]")
var ArgumentNumberSubtractFailed = errors.New("argument length should be 4 with object [name, number/name, email, account_token]")
var ArgumentNumberMultiplyFailed = errors.New("argument length should be 4 with object [name, number/name, email, account_token]")
var ArgumentNumberDivideFailed = errors.New("argument length should be 4 with object [name, number/name, email, account_token]")
var ArgumentValidateEmailFailed = errors.New("arguments should be [email:passcode]")
var NameEmptyFailed = errors.New("name is empty")
var NumberEmptyFailed = errors.New("number is empty")
var EmailEmptyFailed = errors.New("email is empty")
var TokenEmptyFailed = errors.New("token is empty")
var NumberParsingFailed = errors.New("parsing number failed")
var EmailVerifyFailed = errors.New("email verify failed")
var EmailActivateFailed = errors.New("email activate failed")
var EmailNotVerified = errors.New("email not verified yet")
var EmailFormatFailed = errors.New("incorrect email format")
var ValidateCredentialsByEmailFailed = errors.New("email with credential failed")
var GenerateAccountTokenFailed = errors.New("generate account token failed")
var EmptyPassCode = errors.New("empty passcode")
var ValidatePasscodeFailed = errors.New("validate passcode failed")
var EmptyPasswordFailed = errors.New("empty password failed")
var ValidateTokenFailed = errors.New("validate token failed")
var DeleteAccountFailed = errors.New("delete account failed")
