package main

import (
	"log"

	"github.com/go-daq/canbus"
)

func main() {

	sck, err := canbus.New()
	if err != nil {
		log.Fatal(err)
	}
	defer sck.Close()
	addr := "vcan0"

	err = sck.Bind(addr)
	if err != nil {
		log.Fatalf("error binding to [%s]: %v\n", addr, err)
	}

	//for {
	//	msg, err := sck.Recv()
	//}

}
