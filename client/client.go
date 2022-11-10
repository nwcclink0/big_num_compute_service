package main

import (
	"big_num_compute_service/config"
	"big_num_compute_service/rpc/jsonrpc"
	"big_num_compute_service/service"
	"flag"
	"fmt"
	"log"
	//"net/rpc/jsonrpc"
)

type Arg struct {
	A int
	B int
}

func main() {
	var (
		configFile string
	)
	var err error

	flag.StringVar(&configFile, "c", "", "Configuration file path")
	service.BigNumComputeConf, err = config.LoadConf(configFile)
	if err != nil {
		log.Fatalf("Load yaml config file error: '%v'", err)
		return
	}

	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:"+service.BigNumComputeConf.Core.Port)
	if err != nil {
		fmt.Println("can't dial to localhost with " + service.BigNumComputeConf.Core.Port)
	}
	defer conn.Close()
	var data string
	var args = []string{"dog", "1110"}
	//var args = []string{"dog",""}
	//var args = interface{"dog", "10"}
	//var arg = Arg{}
	//arg.A = 1
	//arg.B = 2
	err = conn.Call("BigNumCompute.Create", args, &data)
	if err != nil {
		fmt.Println("can't call rpc, reason:", err)
	}
	fmt.Println(data)
}
