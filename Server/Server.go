package Server

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	_"time"
)

type Args struct{}

type TimeServer int64

func (t *TimeServer) GetTime(args *Args, reply *string) error {
	*reply = "TEST SUCCEEDED"
	return nil
} 




func startHTTPServer(l net.Listener) {
	err := http.Serve(l, nil)
	if err != nil {
		log.Fatal("HTTP serve error:", err)
	}
}


func Server() string {

	timeserver := new(TimeServer)

	// register RPC
	rpc.Register(timeserver)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	go  startHTTPServer(l); 


	return "SERVER STARTED"
}