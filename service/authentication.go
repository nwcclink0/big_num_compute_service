package service

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base32"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/argon2"
	"os"
	"strings"
	"time"
)

type Argon2Params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

func InitArgon2Params() {
	Argon2Conf = &Argon2Params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
}

func GenerateHashPassword(password string) (string, error) {
	salt, err := generateRandomBytes(Argon2Conf.saltLength)
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password), salt, Argon2Conf.iterations, Argon2Conf.memory,
		Argon2Conf.parallelism, Argon2Conf.keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version,
		Argon2Conf.memory, Argon2Conf.iterations, Argon2Conf.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func VerifyAccountEmail(email string, passcode string) (bool, error) {
	secret := base32.StdEncoding.EncodeToString([]byte(email))
	validate, err := totp.ValidateCustom(passcode, secret, time.Now(), totp.ValidateOpts{
		Period:    120,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA512,
	})
	if err != nil {
		return false, err
	}
	return validate, nil
}

func GenerateTotp(email string) (string, error) {
	secret := base32.StdEncoding.EncodeToString([]byte(email))
	passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    120,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA512,
	})
	if err != nil {
		panic(err)
	}
	return passcode, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func ValidateCredentialsByEmail(account *Account, password string) (bool, error) {
	return compareHashPasswordAndPassword(account.HashPassword, password)
}

func compareHashPasswordAndPassword(hashPassword string, password string) (bool, error) {
	params, salt, hash, err := decodeHash(hashPassword)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey([]byte(password), salt, params.iterations,
		params.memory, params.parallelism, params.keyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func decodeHash(encodedHash string) (param *Argon2Params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	param = &Argon2Params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &param.memory,
		&param.iterations, &param.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	param.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	param.keyLength = uint32(len(hash))

	return param, salt, hash, nil
}

func GenerateToken(account *Account) (string, error) {
	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["account"] = account.Email
	claims["id"] = account.Id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", fmt.Errorf("something went wrong: %s", err.Error())
	}
	return tokenString, nil
}

func ValidateToken(token string, account *Account) (bool, error) {
	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, fmt.Errorf("there was an error in parsing")
		}
		return jwtKey, nil
	})
	if err != nil {
		return false, err
	}

	if parsedToken == nil {
		return false, fmt.Errorf("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return false, fmt.Errorf("cant parse claims")
	}
	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		return false, fmt.Errorf("token expired")
	}

	id := claims["id"].(string)
	if id != account.Id.String() {
		return false, fmt.Errorf("id didn't match")
	}
	return true, nil
}
