package main

import (
	"crypto/tls"
	"fmt"
	"github.com/go-netty/go-netty-transport/quic"
	"os"
	"web-programming-class-experiments/chatroom/startup/client"
)

func main() {
	var fromID = os.Getenv("FROM_ID")
	fmt.Println("from id:", fromID)
	client.ClientMain(quic.New(), "quic://localhost:5750", fromID, quic.WithOptions(&quic.Options{
		TLS: &tls.Config{
			InsecureSkipVerify: true,
			NextProtos:         []string{"quic-echo-example"},
		},
	}))
}
