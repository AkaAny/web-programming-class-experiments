package server

import (
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/codec/format"
	"web-programming-class-experiments/chatroom/internal"
	"web-programming-class-experiments/chatroom/internal/server"
)

func ServerMain(transportFactory netty.TransportFactory, address string) { //5750
	// child pipeline initializer.
	var serverMessageHandler = server.NewServerMessageHandler()
	setupCodec := func(channel netty.Channel) {
		internal.WithProtocol(channel.Pipeline(), func(pipeline netty.Pipeline) {
			pipeline.AddLast(format.JSONCodec(false, false)).
				AddLast(serverMessageHandler)
		})
	}

	// setup bootstrap & startup server.
	if err := netty.NewBootstrap(netty.WithChildInitializer(setupCodec),
		netty.WithTransport(transportFactory)).Listen(address).Sync(); err != nil {
		panic(err)
	}
}
