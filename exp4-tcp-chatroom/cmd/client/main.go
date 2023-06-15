package main

import (
	"fmt"
	"github.com/go-netty/go-netty/transport/tcp"
	"os"
	"web-programming-class-experiments/chatroom/startup/client"
)

func main() {
	var fromID = os.Getenv("FROM_ID")
	fmt.Println("from id:", fromID)
	client.ClientMain(tcp.New(), "tcp://0.0.0.0:5750", fromID)
}
