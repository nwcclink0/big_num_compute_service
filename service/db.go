package service

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math/big"
	"strconv"
)

type Number struct {
	gorm.Model
	Name   string `gorm:"primaryKey"`
	Number float64
}

const dsn = "host=localhost user=yt password=yt dbname=big_num port=5432 sslmode=disable"

func InitDb() {
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		LogError.Error(err)
		panic(err)
	}
	err = db.AutoMigrate(&Number{})
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
