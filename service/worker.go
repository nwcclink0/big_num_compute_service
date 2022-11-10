package service

import (
	"big_num_compute_service/rpc/jsonrpc"
	"net"
	"strconv"
)

func InitWorker(workerNum int64, queueNum int64) {
	LogAccess.Debug("worker number is " + strconv.FormatInt(workerNum,
		10) + ", " +
		"queue number is " + strconv.FormatInt(queueNum, 10))
	QueueComputeWorker = make(chan net.Conn, queueNum)
	for i := int64(0); i < workerNum; i++ {
		go startWorker()
	}
}

func startWorker() {
	for {
		conn := <-QueueComputeWorker
		LogAccess.Debug("start client connection")
		jsonrpc.ServeConn(conn)
		LogAccess.Debug("finish client connection")
	}
}
