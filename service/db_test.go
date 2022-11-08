package service

import (
	"big_num_compute_service/config"
	"fmt"
	"log"
	"math/big"
	"testing"
)

func initTest() {
	BigNumComputeConf, _ = config.LoadConf("")
	if err := InitLog(); err != nil {
		log.Fatal(err)
	}
	BigNumComputeConf.Core.Mode = "test"
	InitDb()
}

func TestCreateObj(t *testing.T) {
	initTest()
	err := CreateObj("my_weight", 92)
	if err != nil {
		t.Errorf("can't create num obj")
		return
	}
	err = DeleteObj("my_weight")
	if err != nil {
		t.Errorf("can't delete num obj")
		return
	}
}

func TestDeleteObj(t *testing.T) {
	initTest()
	err := DeleteObj("my_weight")
	if err != nil {
		t.Errorf("can't delete num obj")
		return
	}
}

func TestUpdateObj(t *testing.T) {
	initTest()
	err := CreateObj("my_weight", 92)
	if err != nil {
		t.Errorf("can't create num obj")
		return
	}
	err = UpdateObj("my_weight", 81)
	if err != nil {
		t.Errorf("can't update num obj")
	}
	err = DeleteObj("my_weight")
	if err != nil {
		t.Errorf("can't delete num obj")
	}
}

func TestAddComputeWithNumber(t *testing.T) {
	initTest()
	myWeight := 92.23103102319203192
	err := CreateObj("my_weight", myWeight)
	if err != nil {
		t.Errorf("can't create num obj")
		return
	}
	t.Logf("Test add")
	testVal := 1.4399349834
	testValStr := fmt.Sprint(1.4399349834)
	result, err := Compute("my_weight", testValStr, AddOp)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf("Add compute result:%f", result)

	actualResult, _ := big.NewFloat(myWeight).Add(big.NewFloat(myWeight), big.NewFloat(testVal)).Float64()
	if actualResult != result {
		t.Errorf("Add compute result not same")
	}
	err = DeleteObj("my_weight")
	if err != nil {
		t.Errorf("can't delete num obj")
	}
}

func TestAddComputeWithName(t *testing.T) {
	initTest()
	myWeight := 92.23103102319203192
	err := CreateObj("my_weight", myWeight)
	if err != nil {
		t.Errorf("can't create num obj")
		return
	}
	t.Logf("Test add")
	testVal := 1.4399349834
	err = CreateObj("dog_weight", testVal)
	if err != nil {
		t.Errorf(err.Error())
	}
	result, err := Compute("my_weight", "dog_weight", AddOp)
	t.Logf("Add compute result:%f", result)

	actualResult, _ := big.NewFloat(myWeight).Add(big.NewFloat(myWeight), big.NewFloat(testVal)).Float64()
	if actualResult != result {
		t.Errorf("Add compute result not same")
	}
	err = DeleteObj("my_weight")
	if err != nil {
		t.Errorf("can't delete num obj")
	}
	err = DeleteObj("dog_weight")
	if err != nil {
		t.Errorf("can't delete num obj")
	}
}

func TestSubtractComputeWithNumber(t *testing.T) {
	initTest()
	myWeight := 92.23103102319203192
	err := CreateObj("my_weight", myWeight)
	if err != nil {
		t.Errorf("can't create num obj")
		return
	}
	t.Logf("Test add")
	testVal := 1.4399349834
	testValStr := fmt.Sprint(1.4399349834)
	result, err := Compute("my_weight", testValStr, SubtractOp)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf("Add compute result:%f", result)

	actualResult, _ := big.NewFloat(myWeight).Sub(big.NewFloat(myWeight), big.NewFloat(testVal)).Float64()
	if actualResult != result {
		t.Errorf("Add compute result not same")
	}
	err = DeleteObj("my_weight")
	if err != nil {
		t.Errorf("can't delete num obj")
	}
}

func TestSubtractComputeWithName(t *testing.T) {
	initTest()
	myWeight := 92.23103102319203192
	err := CreateObj("my_weight", myWeight)
	if err != nil {
		t.Errorf("can't create num obj")
		return
	}

	t.Logf("Test sub")
	testVal := 1.4399349834
	err = CreateObj("dog_weight", testVal)
	if err != nil {
		t.Errorf("can't create num obj")
		return
	}
	result, err := Compute("my_weight", "dog_weight", SubtractOp)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf("Add compute result:%f", result)

	actualResult, _ := big.NewFloat(myWeight).Sub(big.NewFloat(myWeight), big.NewFloat(testVal)).Float64()
	if actualResult != result {
		t.Errorf("Add compute result not same")
	}
	err = DeleteObj("my_weight")
	if err != nil {
		t.Errorf("can't delete num obj")
	}
	err = DeleteObj("dog_weight")
	if err != nil {
		t.Errorf("can't delete num obj")
	}
}

func TestMultiplyComputeWithNumber(t *testing.T) {
	initTest()
	myWeight := 92.23103102319203192
	err := CreateObj("my_weight", myWeight)
	if err != nil {
		t.Errorf("can't create num obj")
		return
	}
	t.Logf("Test multiply")
	testVal := 1.4399349834
	testValStr := fmt.Sprint(1.4399349834)
	result, err := Compute("my_weight", testValStr, MultiplyOp)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf("Multiply compute result:%f", result)

	actualResult, _ := big.NewFloat(myWeight).Mul(big.NewFloat(myWeight), big.NewFloat(testVal)).Float64()
	if actualResult != result {
		t.Errorf("Multiply compute result not same")
	}
	err = DeleteObj("my_weight")
	if err != nil {
		t.Errorf("can't delete num obj")
	}
}

func TestMultiplyComputeWithName(t *testing.T) {
	initTest()
	myWeight := 92.23103102319203192
	err := CreateObj("my_weight", myWeight)
	if err != nil {
		t.Errorf("can't create num obj")
		return
	}

	t.Logf("Test mul")
	testVal := 1.4399349834
	err = CreateObj("dog_weight", testVal)
	if err != nil {
		t.Errorf("can't create num obj")
		return
	}

	result, err := Compute("my_weight", "dog_weight", MultiplyOp)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf("Add compute result:%f", result)

	actualResult, _ := big.NewFloat(myWeight).Mul(big.NewFloat(myWeight), big.NewFloat(testVal)).Float64()
	if actualResult != result {
		t.Errorf("Add compute result not same")
	}
	err = DeleteObj("my_weight")
	if err != nil {
		t.Errorf("can't delete num obj")
	}
	err = DeleteObj("dog_weight")
	if err != nil {
		t.Errorf("can't delete num obj")
	}
}

func TestDivideComputeNumber(t *testing.T) {
	initTest()
	myWeight := 92.23103102319203192
	err := CreateObj("my_weight", myWeight)
	if err != nil {
		t.Errorf("can't create num obj")
		return
	}
	t.Logf("Test multiply")
	testVal := 1.4399349834
	testValStr := fmt.Sprint(1.4399349834)
	result, err := Compute("my_weight", testValStr, DivideOp)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf("Multiply compute result:%f", result)

	actualResult, _ := big.NewFloat(myWeight).Quo(big.NewFloat(myWeight), big.NewFloat(testVal)).Float64()
	if actualResult != result {
		t.Errorf("Multiply compute result not same")
	}
	err = DeleteObj("my_weight")
	if err != nil {
		t.Errorf("can't delete num obj")
	}
}

func TestDivideComputeName(t *testing.T) {
	initTest()
	myWeight := 92.23103102319203192
	err := CreateObj("my_weight", myWeight)
	if err != nil {
		t.Errorf("can't create num obj")
		return
	}

	t.Logf("Test mul")
	testVal := 1.4399349834
	err = CreateObj("dog_weight", testVal)
	if err != nil {
		t.Errorf("can't create num obj")
		return
	}

	result, err := Compute("my_weight", "dog_weight", DivideOp)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf("Add compute result:%f", result)

	actualResult, _ := big.NewFloat(myWeight).Quo(big.NewFloat(myWeight), big.NewFloat(testVal)).Float64()
	if actualResult != result {
		t.Errorf("Add compute result not same")
	}
	err = DeleteObj("my_weight")
	if err != nil {
		t.Errorf("can't delete num obj")
	}
	err = DeleteObj("dog_weight")
	if err != nil {
		t.Errorf("can't delete num obj")
	}
}
