package main

import (
	"fmt"
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty-transport/udp"
	"github.com/go-netty/go-netty/codec/format"
	"os"
	"web-programming-class-experiments/exp3-udp-chatroom/internal"
	"web-programming-class-experiments/exp3-udp-chatroom/internal/client"
)

func main() {
	setupCodec := func(channel netty.Channel) {
		//magic:wpcechat
		internal.WithProtocol(channel.Pipeline(), func(pipeline netty.Pipeline) {
			pipeline.AddLast(format.JSONCodec(false, false)).
				AddLast(&client.ClientMessageHandler{})
		})
	}
	// setup bootstrap & startup server.
	ch, err := netty.NewBootstrap(netty.WithClientInitializer(setupCodec),
		netty.WithTransport(udp.New())).Connect("udp://0.0.0.0:5750", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(ch)
	var fromID = os.Getenv("FROM_ID")
	fmt.Println("from id:", fromID)
	ch.Write(internal.TransferPacket{
		FromID: fromID,
		ToID:   "",
		Method: "auth",
		Data:   []byte("password"),
	})
	for {
		var toID = ""
		var data = ""
		_, err := fmt.Scanf("%s %s", &toID, &data)
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
