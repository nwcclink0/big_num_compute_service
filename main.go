package main

import (
	"big_num_compute_service/config"
	"big_num_compute_service/service"
	"flag"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main() {
	var (
		configFile string
	)
	var err error
	flag.Usage = usage
	flag.StringVar(&configFile, "c", "", "Configuration file path")
	flag.Parse()
	service.BigNumComputeConf, err = config.LoadConf(configFile)
	if err != nil {
		log.Fatalf("Load yaml config file error: '%v'", err)
		return
	}

	if err = service.InitLog(); err != nil {
		log.Fatalf("Can't load log module, error: %v", err)
	}

	service.InitDb()
	service.InitWorker(service.BigNumComputeConf.Core.WorkerNum, service.BigNumComputeConf.Core.QueueNum)
	service.InitArgon2Params()

	service.Run()
}

var usageStr = `
Usage: big_number_compute service [options]
Server Options:
	-c, --config <file>
`

func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}
