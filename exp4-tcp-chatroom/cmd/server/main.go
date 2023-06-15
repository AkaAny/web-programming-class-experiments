package main

import (
	"github.com/go-netty/go-netty/transport/tcp"
	"web-programming-class-experiments/chatroom/startup/server"
)

func main() {
	server.ServerMain(tcp.New(), "tcp://0.0.0.0:5750")
}
