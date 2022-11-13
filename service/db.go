package service

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math/big"
	"os"
	"strconv"
)

type Number struct {
	gorm.Model
	Name   string `gorm:"primaryKey"`
	Number float64
}

type Account struct {
	gorm.Model
	Id           uuid.UUID `gorm:"type:uuid;primary_key;"`
	Email        string
	Activated    bool
	HashPassword string
}

const dsnDocker = "host=db user=%s password=%s dbname=big_num port=5432 sslmode=disable"
const dsnLocalhost = "host=localhost user=%s password=%s dbname=big_num port=5432 sslmode=disable"

func InitDb() {
	var err error
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_USER")
	connectCmd := fmt.Sprintf(dsnLocalhost, dbUser, dbPassword)
	if BigNumComputeConf.Core.Mode == Docker {
		connectCmd = fmt.Sprintf(dsnDocker, dbUser, dbPassword)
	}
	db, err = gorm.Open(postgres.Open(connectCmd), &gorm.Config{})
	if err != nil {
		LogError.Error(err)
		panic(err)
	}
	err = db.AutoMigrate(&Number{})
	if err != nil {
		LogError.Error(err.Error())
	}
	err = db.AutoMigrate(&Account{})
	if err != nil {
		LogError.Error(err.Error())
	}
}

func CreateObj(name string, number float64) error {
	var numberObj Number
	result := db.First(&numberObj, Number{
		Name: name,
	})

	if result.Error == nil { // update
		LogAccess.Debug("name:", name, " with number: ", number, " is exist")
		return fmt.Errorf("name: " + name + " with number: " + fmt.Sprint(number) + " is exist")
	} else { // new insert
		db.Create(&Number{
			Name:   name,
			Number: number,
		})
		return nil
	}
}

func UpdateObj(name string, number float64) error {
	var numberObj Number
	result := db.First(&numberObj, Number{
		Name: name,
	})
	if result.Error != nil {
		LogAccess.Debug("name:", name, " with number: ", number, "isn't exist")
		return fmt.Errorf("name: " + name + " with number: " + fmt.Sprint(number) + " isn't exist")
	} else {
		db.Model(&numberObj).Updates(&Number{
			Name:   name,
			Number: number,
		})
		return nil
	}
}

func DeleteObj(name string) error {
	var numberObj Number
	result := db.First(&numberObj, Number{
		Name: name,
	})
	if result.Error != nil {
		LogAccess.Debug("name:", name, "isn't exist")
		return fmt.Errorf("name: " + name + " isn't exist" + ", error:" + result.Error.Error())
	} else {
		db.Delete(&numberObj)
		return nil
	}
}

func Compute(numberArg1 string, numberArg2 string, operation string) (float64, error) {
	var numberObj1 Number
	result := db.First(&numberObj1, Number{
		Name: numberArg1,
	})
	if result.Error != nil { // name didn't exist
		return 0, fmt.Errorf("name: " + numberArg1 + " didn't exist")
	}

	number, err := strconv.ParseFloat(numberArg2, 64)
	var isNum bool
	var numberObj2 Number
	if err != nil { //arg2 is name
		LogAccess.Debug("arg 2 is name")
		result = db.First(&numberObj2, Number{
			Name: numberArg2,
		})
		if result.Error != nil {
			return 0, fmt.Errorf("name: " + numberArg2 + "didn't exist")
		}
		isNum = false
	} else {
		LogAccess.Debug("arg 2 is number")
		isNum = true
	}

	bigNumber1 := big.NewFloat(numberObj1.Number)
	var bigNumber2 *big.Float
	if isNum {
		bigNumber2 = big.NewFloat(number)
	} else {
		bigNumber2 = big.NewFloat(numberObj2.Number)
	}

	var newBigNumber *big.Float
	if operation == AddOp {
		newBigNumber = bigNumber1.Add(bigNumber1, bigNumber2)
	} else if operation == SubtractOp {
		newBigNumber = bigNumber1.Sub(bigNumber1, bigNumber2)
	} else if operation == MultiplyOp {
		newBigNumber = bigNumber1.Mul(bigNumber1, bigNumber2)
	} else if operation == DivideOp {
		newBigNumber = bigNumber1.Quo(bigNumber1, bigNumber2)
	} else {
		return 0, nil
	}

	newVal, _ := newBigNumber.Float64()
	return newVal, nil
}

func AddAccount(email string, hashPassword string) error {
	var account Account
	result := db.First(&account, Account{
		Email: email,
	})
	if result.Error == nil { // name didn't exist
		return fmt.Errorf("email: " + email + " exist")
	}
	db.Create(&Account{
		Id:           uuid.New(),
		Email:        email,
		HashPassword: hashPassword,
		Activated:    false,
	})
	return nil
}

func DeleteAccount(email string) error {
	var account Account
	result := db.First(&account, Account{
		Email: email,
	})
	if result.Error != nil {
		return fmt.Errorf("account: " + email + " don't exist")
	}
	db.Delete(&account)
	return nil
}

func GetAccount(email string) (*Account, error) {
	var account Account
	result := db.First(&account, Account{
		Email: email,
	})
	if result.Error != nil { // name didn't exist
		return nil, fmt.Errorf("email: " + email + " exist")
	}
	return &account, nil
}

func UpdateAccount(account *Account) error {
	var checkAccount Account
	result := db.First(&checkAccount)
	if result.Error != nil { // name didn't exist
		return fmt.Errorf("Can't find account: " + account.Email)
	}
	db.Model(account).Updates(&Account{
		Email:        account.Email,
		HashPassword: account.HashPassword,
		Activated:    account.Activated,
	})
	return nil
}
