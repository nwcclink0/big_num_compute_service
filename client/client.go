package main

import (
	"big_num_compute_service/config"
	"big_num_compute_service/rpc"
	"big_num_compute_service/rpc/jsonrpc"
	"big_num_compute_service/service"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	//"net/rpc/jsonrpc"
)

type Arg struct {
	A int
	B int
}

const email = "yuantingwei@pm.me"
const password = "yt"

func CreateAccount() {
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:"+service.BigNumComputeConf.Core.Port)
	if err != nil {
		fmt.Println("can't dial to localhost with " + service.BigNumComputeConf.Core.Port)
	}
	defer func(conn *rpc.Client) {
		err := conn.Close()
		if err != nil {
			fmt.Println("can't close connection, error: " + err.Error())
		}
	}(conn)
	var args = []string{email, password}
	var data string
	err = conn.Call("createaccount", args, &data)
	if err != nil {
		fmt.Println("can't call rpc, reason: ", err)
		return
	}
	passcode := data
	args = []string{email, passcode}
	err = conn.Call("validateemail", args, &data)
	if err != nil {
		fmt.Println("can't call rpc, reason: ", err)
		return
	}
	if data != "success" {
		fmt.Println("can't verify email")
		return
	}
	fmt.Println(data)
}

func LoginAccount() (string, error) {
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:"+service.BigNumComputeConf.Core.Port)
	if err != nil {
		return "", fmt.Errorf("can't dial to localhost with " + service.BigNumComputeConf.Core.Port)
	}
	defer func(conn *rpc.Client) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	var args = []string{email, password}
	var data string
	err = conn.Call("loginaccount", args, &data)
	if err != nil {
		return "", fmt.Errorf("can't call rpc, reason:" + err.Error())
	}
	token := data
	return token, nil
}

func DeleteAccount() {
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:"+service.BigNumComputeConf.Core.Port)
	if err != nil {
		fmt.Println("can't dial to localhost with " + service.BigNumComputeConf.Core.Port)
	}
	defer func(conn *rpc.Client) {
		err := conn.Close()
		if err != nil {
			fmt.Println("can't close connection, error: " + err.Error())
		}
	}(conn)
	var args = []string{email, password}
	var data string
	err = conn.Call("loginaccount", args, &data)
	if err != nil {
		fmt.Println("can't call rpc, reason: ", err)
	}
	token := data
	args = []string{email, token}
	err = conn.Call("deleteaccount", args, &data)
	if err != nil {
		fmt.Println("can't call rpc, reason: ", err)
	}
	if data != "success" {
		fmt.Println("can't verify email")
		return
	}
	fmt.Println(data)
}

func main() {
	var (
		configFile string
		op         string
	)
	var err error
	flag.Usage = usage
	flag.StringVar(&configFile, "c", "", "Configuration file path")
	flag.StringVar(&op, "o", "", "operation mode: "+
		"create/delete/update/create_me/create_cat/delete_me/delete_cat/\n"+
		"compute_me_cat_add/compute_me_cat_sub/compute_me_cat_mul/compute_me_cat_div/\n"+
		"compute_me_add_random/compute_me_sub_random/compute_me_mul_random/compute_me_div_random\n")
	flag.Parse()

	if len(op) == 0 {
		log.Fatalf("please provide op mode")
		return
	}
	service.BigNumComputeConf, err = config.LoadConf(configFile)
	if err != nil {
		log.Fatalf("Load yaml config file error: '%v'", err)
		return
	}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Nanosecond)
	service.BigNumComputeConf.Core.Mode = service.Localhost
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:"+service.BigNumComputeConf.Core.Port)
	if err != nil {
		fmt.Println("can't dial to localhost with " + service.BigNumComputeConf.Core.Port)
	}
	defer func(conn *rpc.Client) {
		err := conn.Close()
		if err != nil {
			fmt.Println("can't close connection, error: " + err.Error())
		}
	}(conn)

	var data string
	min := -100000000000000000.123
	max := 1000000000000000000.123
	val := min + rand.Float64()*(max-min)
	if op == "create" {
		token, err := LoginAccount()
		var args = []string{"dog", fmt.Sprint(val), email, token}
		err = conn.Call("create", args, &data)
		if err != nil {
			fmt.Println("can't call rpc, reason:", err)
		}
	} else if op == "delete" {
		token, err := LoginAccount()
		var args = []string{"dog", email, token}
		err = conn.Call("delete", args, &data)
		if err != nil {
			fmt.Println("can't call rpc, reason:", err)
		}
	} else if op == "update" {
		token, err := LoginAccount()
		var args = []string{"dog", fmt.Sprint(val), email, token}
		err = conn.Call("update", args, &data)
		if err != nil {
			fmt.Println("can't call rpc, reason:", err)
		}
	} else if op == "create_me" {
		token, err := LoginAccount()
		var args = []string{"me", fmt.Sprint(val), email, token}
		err = conn.Call("create", args, &data)
		if err != nil {
			fmt.Println("can't call rpc, reason:", err)
		}
	} else if op == "create_cat" {
		token, err := LoginAccount()
		var args = []string{"cat", fmt.Sprint(val), email, token}
		err = conn.Call("create", args, &data)
		if err != nil {
			fmt.Println("can't call rpc, reason:", err)
		}
	} else if op == "delete_me" {
		token, err := LoginAccount()
		var args = []string{"me", email, token}
		err = conn.Call("delete", args, &data)
		if err != nil {
			fmt.Println("can't call rpc, reason:", err)
		}
	} else if op == "delete_cat" {
		token, err := LoginAccount()
		var args = []string{"cat", email, token}
		err = conn.Call("delete", args, &data)
		if err != nil {
			fmt.Println("can't call rpc, reason:", err)
		}
	} else if op == "compute_me_cat_add" {
		token, err := LoginAccount()
		var args = []string{"me", "cat", email, token}
		err = conn.Call("add", args, &data)
		if err != nil {
			fmt.Println("can't call rpc, reason:", err)
		}
	} else if op == "compute_me_cat_sub" {
		token, err := LoginAccount()
		var args = []string{"me", "cat", email, token}
		err = conn.Call("subtract", args, &data)
		if err != nil {
			fmt.Println("can't call rpc, reason:", err)
		}
	} else if op == "compute_me_cat_mul" {
		token, err := LoginAccount()
		var args = []string{"me", "cat", email, token}
		err = conn.Call("multiply", args, &data)
		if err != nil {
			fmt.Println("can't call rpc, reason:", err)
		}
	} else if op == "compute_me_cat_div" {
		token, err := LoginAccount()
		var args = []string{"me", "cat", email, token}
		err = conn.Call("divide", args, &data)
		if err != nil {
			fmt.Println("can't call rpc, reason:", err)
		}
	} else if op == "compute_me_add_random" {
		token, err := LoginAccount()
		var args = []string{"me", fmt.Sprint(val), email, token}
		err = conn.Call("add", args, &data)
		if err != nil {
			fmt.Println("can't call rpc, reason:", err)
		}
	} else if op == "compute_me_sub_random" {
		token, err := LoginAccount()
		var args = []string{"me", fmt.Sprint(val), email, token}
		err = conn.Call("subtract", args, &data)
		if err != nil {
			fmt.Println("can't call rpc, reason:", err)
		}
	} else if op == "compute_me_mul_random" {
		token, err := LoginAccount()
		var args = []string{"me", fmt.Sprint(val), email, token}
		err = conn.Call("multiply", args, &data)
		if err != nil {
			fmt.Println("can't call rpc, reason:", err)
		}
	} else if op == "compute_me_div_random" {
		token, err := LoginAccount()
		var args = []string{"me", fmt.Sprint(val), email, token}
		err = conn.Call("divide", args, &data)
		if err != nil {
			fmt.Println("can't call rpc, reason:", err)
		}
	} else if op == "create_account" {
		CreateAccount()
	} else if op == "delete_account" {
		DeleteAccount()
	}
	fmt.Println("result: " + data)
}

var usageStr = `
Usage: big number compute service jsonrpc client [options]
Client Options:
	-c, <file>
	-o, mode, create/delete/update/
            create_me/create_cat/delete_me/delete_cat/
            compute_me_cat_add/compute_me_cat_sub/compute_me_cat_mul/compute_me_cat_div/
            compute_me_add_random/compute_me_sub_random/compute_me_mul_random/compute_me_div_random
`

func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}
