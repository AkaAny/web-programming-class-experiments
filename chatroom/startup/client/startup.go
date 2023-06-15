package client

import (
	"fmt"
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/codec/format"
	"github.com/go-netty/go-netty/transport"
	"web-programming-class-experiments/chatroom/internal"
	"web-programming-class-experiments/chatroom/internal/client"
)

func ClientMain(transportFactory netty.TransportFactory, address string, fromID string, option ...transport.Option) {
	setupCodec := func(channel netty.Channel) {
		internal.WithProtocol(channel.Pipeline(), func(pipeline netty.Pipeline) {
			pipeline.AddLast(format.JSONCodec(false, false)).
				AddLast(&client.ClientMessageHandler{})
		})
	}
	ch, err := netty.NewBootstrap(netty.WithClientInitializer(setupCodec),
		netty.WithTransport(transportFactory)).Connect(address, option...)
	if err != nil {
		panic(err)
	}
	fmt.Println(ch)
	ch.Write(internal.TransferPacket{
		FromID: fromID,
		ToID:   "",
		Method: "auth",
		Data:   []byte("password"),
	})
	for {
		var toID = ""
		var data = ""
		fmt.Println("input toID:")
		_, err := fmt.Scanln(&toID)
		if err != nil {
			panic(err)
		}
		fmt.Println("input msg:")
		_, err = fmt.Scanf("%s", &data)
		if err != nil {
			panic(err)
		}
		fmt.Println("send to id:", toID, "data:", data)
		ch.Write(internal.TransferPacket{
			FromID: fromID,
			ToID:   toID,
			Method: "transfer",
			Data:   []byte(data),
		})
	}
}
