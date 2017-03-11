package main

import (
	"flag"
	"log"
)

var (
	addr  = flag.String("addr", "localhost:6399", "net address")
	proto = flag.String("port", "tcp", "protocol")
)

func main() {
	flag.Parse()
	bigred, err := NewBigRed(*proto, *addr)
	if err != nil {
		log.Fatal(err)
	}
	if err := bigred.Run(); err != nil {
		log.Fatal(err)
	}
}
