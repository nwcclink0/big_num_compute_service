package service

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Number struct {
	gorm.Model
	Name   string
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
		Name:   name,
		Number: number,
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
