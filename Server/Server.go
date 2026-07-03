package Server

import (
	"log"
	"net/http"
	"net/rpc"
)


type Args struct{}


type TimeServer int


func (t *TimeServer) GetTime(args *Args, reply *string) error {
	*reply = "SERVER 1: Time request processed successfully!"
	return nil
}


func runServer(port string, mux *http.ServeMux) {
	log.Printf("Microservice listening on port %s...", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed to start on %s: %v", port, err)
	}
}



func StartServer1() string {
	// Create an isolated RPC registry instance instead of using the global default
	server := rpc.NewServer()
	
	// Register our TimeServer service to this isolated instance
	err := server.Register(new(TimeServer))
	if err != nil {
		log.Fatalf("Registration error: %v", err)
	}

	// Create an isolated HTTP router
	mux := http.NewServeMux()
	
	// Bind our isolated RPC server to this specific router path
	mux.Handle(rpc.DefaultRPCPath, server)

	go runServer(":1234", mux)

	return "SERVER 1 INITIALIZED"
}
