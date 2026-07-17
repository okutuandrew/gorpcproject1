package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

type Values struct {
	A, B int
}

type Quotient struct {  // Fixed spelling to match client
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Values, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Values, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	
	fmt.Println("Server starting on port 1234...")
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println("Server error:", err.Error())
	}
}