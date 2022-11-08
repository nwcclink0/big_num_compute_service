package service

import (
	"big_num_compute_service/config"
	"log"
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
