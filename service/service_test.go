package service

import (
	"fmt"
	"math/big"
	"testing"
)

func TestBigNumCompute_Create_And_Delete(t *testing.T) {
	initTest()
	var arg = []string{"dog", "10"}
	var result string
	bigNumberCompute := BigNumCompute{}
	err := bigNumberCompute.Create(arg, &result)
	if err != nil {
		t.Errorf("can't create object, error: %s", err)
	}
	if result != ResultSuccess {
		t.Errorf("incorrect reuslt:%s", result)
	}

	arg = []string{"dog"}
	err = bigNumberCompute.Delete(arg, &result)
	if err != nil {
		t.Errorf("can't deleteobject, error: %s", err)
	}
	if result != ResultSuccess {
		t.Errorf("incorrect result: %s", result)
	}
}

//// This test is for manual test, uncomment it if necessary
//func TestBigNumCompute_Delete(t *testing.T) {
//	initTest()
//	var arg = []string{"dog", "10"}
//	var result string
//	bigNumberCompute := BigNumCompute{}
//
//	arg = []string{"dog"}
//	err := bigNumberCompute.Delete(arg, &result)
//	if err != nil {
//		t.Errorf("can't deleteobject, error: %s", err)
//	}
//	if result != ResultSuccess {
//		t.Errorf("incorrect result: %s", result)
//	}
//}

func TestBigNumCompute_Update(t *testing.T) {
	initTest()
	var args = []string{"dog", "10"}
	var result string
	bigNumberCompute := BigNumCompute{}
	err := bigNumberCompute.Create(args, &result)
	if err != nil {
		t.Errorf("can't create object, error: %s", err)
	}
	if result != ResultSuccess {
		t.Errorf("incorrect reuslt:%s", result)
	}

	args = []string{"dog", "20"}
	err = bigNumberCompute.Update(args, &result)
	if err != nil {
		t.Errorf("can't update object, error: %s", err)
	}
	if result != ResultSuccess {
		t.Errorf("incorrect reuslt:%s", result)
	}

	args = []string{"dog"}
	err = bigNumberCompute.Delete(args, &result)
	if err != nil {
		t.Errorf("can't deleteobject, error: %s", err)
	}
	if result != ResultSuccess {
		t.Errorf("incorrect result: %s", result)
	}
}

func TestBigNumCompute_Add(t *testing.T) {
	initTest()

	dogWeight := big.NewFloat(float64(10))
	var args = []string{"dog", fmt.Sprint(dogWeight)}
	var result string
	bigNumberCompute := BigNumCompute{}
	err := bigNumberCompute.Create(args, &result)
	if err != nil {
		t.Errorf("can't create object, error: %s", err)
		return
	}
	if result != ResultSuccess {
		t.Errorf("incorrect reuslt:%s", result)
		return
	}

	catWeight := big.NewFloat(float64(20))
	args = []string{"cat", fmt.Sprint(catWeight)}
	err = bigNumberCompute.Create(args, &result)
	if err != nil {
		t.Errorf("can't create object, error: %s", err)
		return
	}
	if result != ResultSuccess {
		t.Errorf("incorrect reuslt:%s", result)
		return
	}

	dogWeight1 := big.NewFloat(10.11340123)
	args = []string{"dog", fmt.Sprint(dogWeight1)}
	err = bigNumberCompute.Add(args, &result)
	if err != nil {
		t.Errorf("can't create object, error: %s", err)
		return
	}

	sum, _ := dogWeight.Add(dogWeight, dogWeight1).Float64()
	if result != fmt.Sprint(sum) {
		t.Errorf("incorrect reuslt:%s", result)
		return
	}
	dogWeight.Sub(dogWeight, dogWeight1)

	args = []string{"dog", "cat"}
	err = bigNumberCompute.Add(args, &result)
	if err != nil {
		t.Errorf("can't create object, error: %s", err)
		return
	}
	sum, _ = dogWeight.Add(dogWeight, catWeight).Float64()
	if result != fmt.Sprint(sum) {
		t.Errorf("incorrect reuslt:%s", result)
		return
	}

	args = []string{"dog"}
	err = bigNumberCompute.Delete(args, &result)
	if result != ResultSuccess {
		t.Errorf("cant delete %s, error: %s", args[0], err)
	}

	args = []string{"cat"}
	err = bigNumberCompute.Delete(args, &result)
	if result != ResultSuccess {
		t.Errorf("cant delete %s, error: %s", args[0], err)
	}
}

func TestBigNumCompute_Subtract(t *testing.T) {
	initTest()

	dogWeight := big.NewFloat(float64(10))
	var args = []string{"dog", fmt.Sprint(dogWeight)}
	var result string
	bigNumberCompute := BigNumCompute{}
	err := bigNumberCompute.Create(args, &result)
	if err != nil {
		t.Errorf("can't create object, error: %s", err)
		return
	}
	if result != ResultSuccess {
		t.Errorf("incorrect reuslt:%s", result)
		return
	}

	catWeight := big.NewFloat(float64(20))
	args = []string{"cat", fmt.Sprint(catWeight)}
	err = bigNumberCompute.Create(args, &result)
	if err != nil {
		t.Errorf("can't create object, error: %s", err)
		return
	}
	if result != ResultSuccess {
		t.Errorf("incorrect reuslt:%s", result)
		return
	}

	dogWeight1 := big.NewFloat(10.11340123)
	args = []string{"dog", fmt.Sprint(dogWeight1)}
	err = bigNumberCompute.Subtract(args, &result)
	if err != nil {
		t.Errorf("can't create object, error: %s", err)
		return
	}

	sum, _ := dogWeight.Sub(dogWeight, dogWeight1).Float64()
	if result != fmt.Sprint(sum) {
		t.Errorf("incorrect reuslt:%s", result)
		return
	}
	dogWeight.Add(dogWeight, dogWeight1)

	args = []string{"dog", "cat"}
	err = bigNumberCompute.Subtract(args, &result)
	if err != nil {
		t.Errorf("can't create object, error: %s", err)
		return
	}
	sum, _ = dogWeight.Sub(dogWeight, catWeight).Float64()
	if result != fmt.Sprint(sum) {
		t.Errorf("incorrect reuslt:%s", result)
		return
	}

	args = []string{"dog"}
	err = bigNumberCompute.Delete(args, &result)
	if result != ResultSuccess {
		t.Errorf("cant delete %s, error: %s", args[0], err)
	}

	args = []string{"cat"}
	err = bigNumberCompute.Delete(args, &result)
	if result != ResultSuccess {
		t.Errorf("cant delete %s, error: %s", args[0], err)
	}
}

func TestBigNumCompute_Multiply(t *testing.T) {
	initTest()

	dogWeight := big.NewFloat(float64(10))
	var args = []string{"dog", fmt.Sprint(dogWeight)}
	var result string
	bigNumberCompute := BigNumCompute{}
	err := bigNumberCompute.Create(args, &result)
	if err != nil {
		t.Errorf("can't create object, error: %s", err)
		return
	}
	if result != ResultSuccess {
		t.Errorf("incorrect reuslt:%s", result)
		return
	}

	catWeight := big.NewFloat(float64(20))
	args = []string{"cat", fmt.Sprint(catWeight)}
	err = bigNumberCompute.Create(args, &result)
	if err != nil {
		t.Errorf("can't create object, error: %s", err)
		return
	}
	if result != ResultSuccess {
		t.Errorf("incorrect reuslt:%s", result)
		return
	}

	dogWeight1 := big.NewFloat(10.11340123)
	args = []string{"dog", fmt.Sprint(dogWeight1)}
	err = bigNumberCompute.Multiply(args, &result)
	if err != nil {
		t.Errorf("can't create object, error: %s", err)
		return
	}

	sum, _ := dogWeight.Mul(dogWeight, dogWeight1).Float64()
	if result != fmt.Sprint(sum) {
		t.Errorf("incorrect reuslt:%s", result)
		return
	}
	dogWeight.Quo(dogWeight, dogWeight1)

	args = []string{"dog", "cat"}
	err = bigNumberCompute.Multiply(args, &result)
	if err != nil {
		t.Errorf("can't create object, error: %s", err)
		return
	}
	sum, _ = dogWeight.Mul(dogWeight, catWeight).Float64()
	if result != fmt.Sprint(sum) {
		t.Errorf("incorrect reuslt:%s", result)
		return
	}

	args = []string{"dog"}
	err = bigNumberCompute.Delete(args, &result)
	if result != ResultSuccess {
		t.Errorf("cant delete %s, error: %s", args[0], err)
	}

	args = []string{"cat"}
	err = bigNumberCompute.Delete(args, &result)
	if result != ResultSuccess {
		t.Errorf("cant delete %s, error: %s", args[0], err)
	}
}

func TestBigNumCompute_Divide(t *testing.T) {
	initTest()

	dogWeight := big.NewFloat(float64(10))
	var args = []string{"dog", fmt.Sprint(dogWeight)}
	var result string
	bigNumberCompute := BigNumCompute{}
	err := bigNumberCompute.Create(args, &result)
	if err != nil {
		t.Errorf("can't create object, error: %s", err)
		return
	}
	if result != ResultSuccess {
		t.Errorf("incorrect reuslt:%s", result)
		return
	}

	catWeight := big.NewFloat(float64(20))
	args = []string{"cat", fmt.Sprint(catWeight)}
	err = bigNumberCompute.Create(args, &result)
	if err != nil {
		t.Errorf("can't create object, error: %s", err)
		return
	}
	if result != ResultSuccess {
		t.Errorf("incorrect reuslt:%s", result)
		return
	}

	dogWeight1 := big.NewFloat(10.11340123)
	args = []string{"dog", fmt.Sprint(dogWeight1)}
	err = bigNumberCompute.Divide(args, &result)
	if err != nil {
		t.Errorf("can't create object, error: %s", err)
		return
	}

	sum, _ := dogWeight.Quo(dogWeight, dogWeight1).Float64()
	if result != fmt.Sprint(sum) {
		t.Errorf("incorrect reuslt:%s", result)
		return
	}
	dogWeight.Mul(dogWeight, dogWeight1)

	args = []string{"dog", "cat"}
	err = bigNumberCompute.Divide(args, &result)
	if err != nil {
		t.Errorf("can't create object, error: %s", err)
		return
	}
	sum, _ = dogWeight.Quo(dogWeight, catWeight).Float64()
	if result != fmt.Sprint(sum) {
		t.Errorf("incorrect reuslt:%s", result)
		return
	}

	args = []string{"dog"}
	err = bigNumberCompute.Delete(args, &result)
	if result != ResultSuccess {
		t.Errorf("cant delete %s, error: %s", args[0], err)
	}

	args = []string{"cat"}
	err = bigNumberCompute.Delete(args, &result)
	if result != ResultSuccess {
		t.Errorf("cant delete %s, error: %s", args[0], err)
	}
}
