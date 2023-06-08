package internal

import (
	"encoding/base64"
	"github.com/go-netty/go-netty"
)

type TransferPacket struct {
	FromID string
	ToID   string
	Method string
	Data   []byte
}

func FromMessage(msg netty.Message) *TransferPacket {
	var msgMap = msg.(map[string]interface{})
	var tp = new(TransferPacket)
	tp.FromID = msgMap["FromID"].(string)
	tp.ToID = msgMap["ToID"].(string)
	tp.Method = msgMap["Method"].(string)
	var dataBase64 = msgMap["Data"].(string)
	data, err := base64.StdEncoding.DecodeString(dataBase64)
	if err != nil {
		panic(err)
	}
	tp.Data = data
	return tp
}
