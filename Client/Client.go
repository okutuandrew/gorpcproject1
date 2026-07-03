package Client

import (
	"log"
	"net/rpc"
)

// CallAllServices connects to both microservices and prints their responses
func CallAllServices() {
	
	// ------------------------------------------------------------
	// 1. TALK TO MICROSERVICE 1 (Port 1234)
	// ------------------------------------------------------------
	// Dial the HTTP RPC server running on port 1234
	client1, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("Client failed to connect to Server 1: ", err)
	}
	defer client1.Close() // Clean up the connection when finished

	var reply1 string
	// Execute the "GetTime" method belonging to the "TimeServer" struct
	err = client1.Call("TimeServer.GetTime", struct{}{}, &reply1)
	if err != nil {
		log.Println("Server 1 RPC Error: ", err)
	} else {
		log.Println("Client Success -> ", reply1)
	}

	// ------------------------------------------------------------
	// 2. TALK TO MICROSERVICE 2 (Port 8082)
	// ------------------------------------------------------------
	// Dial the second HTTP RPC server running on port 8082
	client2, err := rpc.DialHTTP("tcp", "127.0.0.1:8082")
	if err != nil {
		log.Fatal("Client failed to connect to Server 2: ", err)
	}
	defer client2.Close() // Clean up the connection when finished

	var reply2 string
	// Execute the "GetProfile" method belonging to the "UserData" struct
	err = client2.Call("UserData.GetProfile", struct{}{}, &reply2)
	if err != nil {
		log.Println("Server 2 RPC Error: ", err)
	} else {
		log.Println("Client Success -> ", reply2)
	}
}
