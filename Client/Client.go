package Client

import (
	"log"
	"net/rpc"
)

type Args struct{}

func Client() {

	var reply string
	args := Args{}

	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dialing error:", err)
	}

	err = client.Call("TimeServer.GetTime", args, &reply)
	if err != nil {
		log.Fatal("RPC error:", err)
	}

	log.Println("Server time:", reply)
}