package main 

import (
	"fmt"
	"RpC/Server"
	"RpC/Client"
)

func main(){
	fmt.Println(Server.Server())

	Client.Client()
}