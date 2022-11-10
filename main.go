package main

import (
	"big_num_compute_service/config"
	"big_num_compute_service/service"
	"flag"
	"log"
)

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

	if err = service.InitLog(); err != nil {
		log.Fatalf("Can't load log module, error: %v", err)
	}

	service.InitDb()
	service.InitWorker(service.BigNumComputeConf.Core.WorkerNum, service.BigNumComputeConf.Core.QueueNum)
	service.Run()
}
