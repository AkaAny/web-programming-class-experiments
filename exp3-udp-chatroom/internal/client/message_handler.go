package client

import (
	"fmt"
	"github.com/go-netty/go-netty"
	"web-programming-class-experiments/exp3-udp-chatroom/internal"
)

type ClientMessageHandler struct {
}

func (c *ClientMessageHandler) HandleActive(ctx netty.ActiveContext) {

}

func (c *ClientMessageHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	fmt.Println(message)
	var transferPacket = internal.FromMessage(message)

	fmt.Println("[", transferPacket.FromID, "]", "says:", string(transferPacket.Data))
}

func (c *ClientMessageHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {

}
