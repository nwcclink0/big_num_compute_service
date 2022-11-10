package service

import (
	"big_num_compute_service/rpc"
	"big_num_compute_service/rpc/jsonrpc"
	"fmt"
	"net"
	//"net/rpc/jsonrpc"
	"strconv"
)

type BigNumCompute struct {
}

func (BigNumCompute) Create(args []string, result *string) error {
	*result = ResultFailed
	if len(args) != 2 {
		return fmt.Errorf("argument length shoue be 2 with object name and related number")
	}

	name := args[0]
	LogAccess.Debug("create number object name: " + name)
	numberStr := args[1]
	number, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		return fmt.Errorf("parsing number failed: %s", err.Error())
	}
	LogAccess.Debug("create number object number: ", number)

	err = CreateObj(name, number)
	if err != nil {
		return fmt.Errorf("create number object failed: %s", err)
	}

	*result = ResultSuccess
	return nil
}

func (BigNumCompute) Delete(args []string, result *string) error {
	*result = ResultFailed
	if len(args) != 1 {
		return fmt.Errorf("argument should be a name that want to delete with")
	}
	name := args[0]
	LogAccess.Debug("delete object name")
	err := DeleteObj(name)
	if err != nil {
		return fmt.Errorf("delete number object failed: %s", err)
	}
	*result = ResultSuccess
	return nil
}

func (BigNumCompute) Update(args []string, result *string) error {
	*result = ResultFailed
	if len(args) != 2 {
		return fmt.Errorf("argument length error, it should be [name:number]")
	}
	name := args[0]
	LogAccess.Debug("create number object name: " + name)
	numberStr := args[1]
	number, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		return fmt.Errorf("parsing number failed: %s", err)
	}
	err = UpdateObj(name, number)
	if err != nil {
		return fmt.Errorf("can't update %s with %s, error: %s", args[0], args[1], err)
	}
	*result = ResultSuccess
	return nil
}

func (BigNumCompute) Add(args []string, result *string) error {
	*result = ResultZero
	if len(args) != 2 {
		return fmt.Errorf("argument length error, it should be [name:number] or [name1:name2]")
	}
	computeResult, err := Compute(args[0], args[1], AddOp)
	if err != nil {
		return fmt.Errorf("can't multiply of %s and %s, error: %s", args[0], args[1], err)
	}
	*result = fmt.Sprint(computeResult)
	return nil
}

func (BigNumCompute) Subtract(args []string, result *string) error {
	*result = ResultZero
	if len(args) != 2 {
		return fmt.Errorf("argument length error, it should be [name:number] or [name1:name2]")
	}
	computeResult, err := Compute(args[0], args[1], SubtractOp)
	if err != nil {
		return fmt.Errorf("can't multiply of %s and %s, error: %s", args[0], args[1], err)
	}
	*result = fmt.Sprint(computeResult)
	return nil
}

func (BigNumCompute) Multiply(args []string, result *string) error {
	*result = ResultZero
	if len(args) != 2 {
		return fmt.Errorf("argument length error, it should be [name:number] or [name1:name2]")
	}
	computeResult, err := Compute(args[0], args[1], MultiplyOp)
	if err != nil {
		return fmt.Errorf("can't multiply of %s and %s, error: %s", args[0], args[1], err)
	}
	*result = fmt.Sprint(computeResult)
	return nil
}

func (BigNumCompute) Divide(args []string, result *string) error {
	*result = ResultZero
	if len(args) != 2 {
		return fmt.Errorf("argument length error, it should be [name:number] or [name1:name2]")
	}
	computeResult, err := Compute(args[0], args[1], DivideOp)
	if err != nil {
		return fmt.Errorf("can't multiply of %s and %s, error: %s", args[0], args[1], err)
	}
	*result = fmt.Sprint(computeResult)
	return nil
}

func Run() {
	err := rpc.Register(BigNumCompute{})
	if err != nil {
		LogError.Error("register big num error, ", err.Error())
		return
	}
	listen, err := net.Listen("tcp", ":"+BigNumComputeConf.Core.Port)
	if err != nil {
		LogError.Error("can't listen port: " + BigNumComputeConf.Core.Port + ", err:" + err.Error())
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			LogError.Error("accept error: " + err.Error())
		}
		go jsonrpc.ServeConn(conn)
	}
}
