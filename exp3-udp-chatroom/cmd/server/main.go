package main

import (
	"github.com/go-netty/go-netty-transport/quic"
	"web-programming-class-experiments/chatroom/startup/server"
	"web-programming-class-experiments/exp3-udp-chatroom/common"
)

func main() {
	server.ServerMain(quic.New(), "quic://0.0.0.0:5750", quic.WithOptions(&quic.Options{
		TLS: common.GenerateTLSConfig(),
	}))
}
