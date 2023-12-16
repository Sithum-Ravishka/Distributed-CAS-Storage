package main

import (
	"cas-storage/p2p"
	"fmt"
	"log"
)

func main() {
	// Create TCP transport for port 3000
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%v\n", msg)
		}
	}()

	// Listen on both ports concurrently

	go func() {
		if err := tr.ListenAndAccept(); err != nil {
			log.Fatal(err)
		}
	}()

	// Keep the main goroutine alive
	select {}
}
