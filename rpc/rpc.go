package rpc

import (
	"log"
	"net/http"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

type Cal int

func (cal *Cal) Square(nums int) *Result {
	return &Result{Num: nums, Ans: nums * nums}
}

func (cal *Cal) SquareRpc(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}

func Calc() {
	result := new(Cal).Square(12)
	log.Printf("%d^2 = %d", result.Num, result.Ans)

	cal := new(Cal)
	var resultRpc Result
	err := cal.SquareRpc(11, &resultRpc)
	if err != nil {
		return
	}
	log.Printf("%d^2 = %d", resultRpc.Num, resultRpc.Ans)
}

func RpcServer() {

	rpc.Register(new(Cal))
	rpc.HandleHTTP()

	log.Printf("Serving RPC server on port %d", 8088)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatal("Error serving: ", err)
	}

}
