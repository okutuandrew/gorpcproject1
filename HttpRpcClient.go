package main

import (
	"net/rpc"
	"fmt"
	"log"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {  // Fixed spelling from "Quitent" to "Quotient"
	Quo, Rem int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "server")
		os.Exit(1)
	}
	serverAddress := os.Args[1]

	// Fixed: DailHTTP -> DialHTTP (spelling correction)
	client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	
	args := Args{17, 8}
	var reply int
	var quot Quotient  // Added this variable declaration
	
	// Call Multiply
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
	
	// Call Divide
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d = %d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}