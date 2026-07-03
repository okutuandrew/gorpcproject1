package main

import (
	"fmt"
	"time"
	"RpC/Client"
	"RpC/Server"
)

func main() {
	fmt.Println("=== INITIALIZING MICROSERVICES ===")
	
	// 1. Start Server 1 (runs on port :1234 in the background)
	fmt.Println(Server.StartServer1())
	
	// 2. Start Server 2 (runs on port :8082 in the background)
	fmt.Println(Server.StartServer2())

	// 3. CRUCIAL PAUSE: Give the background goroutines 250 milliseconds 
	// to successfully bind to ports :1234 and :8082 before the client dials.
	time.Sleep(250 * time.Millisecond)

	fmt.Println("\n=== CLIENT RUNNING NETWORK CALLS ===")
	
	// 4. Trigger the client orchestrator to call both servers
	Client.CallAllServices()
}
