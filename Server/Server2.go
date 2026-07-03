package Server

import (
	"log"
	"net/http"
	"net/rpc"
)

// 1. Define the input arguments for this specific service
type Args2 struct{}

// 2. Define the second service object
type UserData struct{}

// 3. Define the RPC method for retrieving profile information
func (u *UserData) GetProfile(args *Args2, reply *string) error {
	*reply = "SERVER 2: User profile 'aokutu' retrieved!"
	return nil
}


func runHttpListener(mux *http.ServeMux) {
	log.Println("Microservice 2 listening on port :8082...")
	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatalf("Server 2 failed to start: %v", err)
	}
}


// 4. Create the function that main.go will call to boot this second server
func StartServer2() string {
	// Create a second isolated, private RPC registry instance
	server := rpc.NewServer()
	
	// Register UserData to this second isolated instance
	err := server.Register(new(UserData))
	if err != nil {
		log.Fatalf("Registration error for Server 2: %v", err)
	}

	// Create a second isolated HTTP router
	mux := http.NewServeMux()
	
	// Bind this second RPC server to its own private router path
	mux.Handle(rpc.DefaultRPCPath, server)

	go runHttpListener(mux) 
	

	return "SERVER 2 INITIALIZED"
}
