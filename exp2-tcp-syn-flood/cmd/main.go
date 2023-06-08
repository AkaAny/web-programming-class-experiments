package main

import (
	"os"
	"os/signal"
	"web-programming-class-experiments/exp2-tcp-syn-flood/internal/raw"
)

func main() {
	var sigChan = make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	var stopChan = make(chan bool)
	if err := raw.StartFlooding(stopChan,
		"localhost", 30000,
		14000,
		"syn"); err != nil {
		panic(err)
	}
	<-sigChan
	<-stopChan
}
