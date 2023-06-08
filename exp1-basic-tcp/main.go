package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
)

type ProtocolData []byte

func FromReader(reader io.Reader) (ProtocolData, error) {
	var dataLen64 int64 = 0
	if err := binary.Read(reader, binary.BigEndian, &dataLen64); err != nil {
		return nil, fmt.Errorf("read len err:%w", err)
	}
	var protocolData = ProtocolData(make([]byte, dataLen64))
	if _, err := reader.Read(protocolData); err != nil {
		panic(err)
	}
	return protocolData, nil
}

func (x ProtocolData) Write(writer io.Writer) (n int, err error) {
	var dataLen64 = int64(len(x))
	if err := binary.Write(writer, binary.BigEndian, dataLen64); err != nil {
		panic(err)
	}
	return writer.Write(x)
}

func main() {
	const bindAddr = "localhost:30000"
	listener, err := net.Listen("tcp", bindAddr)
	if err != nil {
		panic(err)
	}
	var serverReady = make(chan bool)
	go func() {
		for {
			serverReady <- true
			clientConn, err := listener.Accept()
			if err != nil {
				panic(err)
			}
			var serverHelloData = ProtocolData("server hello\n")
			if _, err := serverHelloData.Write(clientConn); err != nil {
				panic(fmt.Errorf("[s] write server hello err:%w", err))
			}
			clientResp, err := FromReader(clientConn)
			if err != nil {
				panic(fmt.Errorf("[s] read client resp err:%w", err))
			}
			fmt.Println("[s] client resp:", string(clientResp))
			if err := clientConn.Close(); err != nil {
				panic(fmt.Errorf("[s] close client conn err:%w", err))
			}
			fmt.Println("[s] close client conn")
		}
	}()
	<-serverReady
	serverConn, err := net.Dial("tcp", bindAddr)
	if err != nil {
		panic(err)
	}
	serverHelloData, err := FromReader(serverConn)
	if err != nil {
		panic(fmt.Errorf("[c] read server hello err:%w", err))
	}
	fmt.Println("[c] server hello:", string(serverHelloData))
	var clientRespData = ProtocolData("client ack\n")
	if _, err := clientRespData.Write(serverConn); err != nil {
		panic(fmt.Errorf("[c] write client resp err:%w", err))
	}
	if err := serverConn.Close(); err != nil {
		panic(fmt.Errorf("[c] close server conn err:%w", err))
	}
	fmt.Println("[c] close server conn")
	var sigChan = make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
}
