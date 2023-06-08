package main

import (
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty-transport/udp"
	"github.com/go-netty/go-netty/codec/format"
	"web-programming-class-experiments/exp3-udp-chatroom/internal"
	"web-programming-class-experiments/exp3-udp-chatroom/internal/server"
)

type ChatRoomServer struct {
}

func main() {
	// child pipeline initializer.
	var serverMessageHandler = server.NewServerMessageHandler()
	setupCodec := func(channel netty.Channel) {
		//magic:wpcechat
		internal.WithProtocol(channel.Pipeline(), func(pipeline netty.Pipeline) {
			pipeline.AddLast(format.JSONCodec(false, false)).
				AddLast(serverMessageHandler)
		})
	}

	// setup bootstrap & startup server.
	if err := netty.NewBootstrap(netty.WithChildInitializer(setupCodec),
		netty.WithTransport(udp.New())).Listen("udp://0.0.0.0:5750").Sync(); err != nil {
		panic(err)
	}
}
