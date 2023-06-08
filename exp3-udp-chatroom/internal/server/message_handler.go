package server

import (
	"fmt"
	"github.com/go-netty/go-netty"
	"github.com/samber/lo"
	"web-programming-class-experiments/exp3-udp-chatroom/internal"
)

type ServerMessageHandler struct {
	idChannelMap map[int64]netty.Channel
	clientMap    map[string][]int64
}

func NewServerMessageHandler() *ServerMessageHandler {
	return &ServerMessageHandler{
		idChannelMap: make(map[int64]netty.Channel),
		clientMap:    make(map[string][]int64),
	}
}

func (m *ServerMessageHandler) HandleActive(ctx netty.ActiveContext) {

}

func (m *ServerMessageHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	fmt.Println(message)
	var transferPacket = internal.FromMessage(message)
	switch transferPacket.Method {
	case "auth":
		//var authRequest = new(auth.AuthRequest)
		//if err := proto.Unmarshal(transferPacket.Data, authRequest); err != nil {
		//	panic(err)
		//}
		var userID = transferPacket.FromID
		channelIDs, _ := m.clientMap[userID]
		var clientChannel = ctx.Channel()
		channelIDs = append(channelIDs, clientChannel.ID())
		m.clientMap[userID] = channelIDs
		m.idChannelMap[clientChannel.ID()] = clientChannel
	case "transfer":
		channelIDs, _ := m.clientMap[transferPacket.ToID]
		for _, channelID := range channelIDs {
			m.idChannelMap[channelID].Write(message)
		}
	case "close":
		channelIDs, _ := m.clientMap[transferPacket.FromID]
		var currentChannelID = ctx.Channel().ID()
		var newChannelIDs = lo.Filter(channelIDs, func(channelID int64, _ int) bool {
			return channelID != currentChannelID
		})
		m.clientMap[transferPacket.FromID] = newChannelIDs
		delete(m.idChannelMap, currentChannelID)
	}
	//ctx.Write(message)
}

func (m *ServerMessageHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {

}
