package main

import (
	"cas-storage/p2p"
	"log"
)

func main() {
	// Create TCP transport for port 3000
	tcpOpts3000 := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tr3000 := p2p.NewTCPTransport(tcpOpts3000)

	// Create TCP transport for port 4000
	tcpOpts4000 := p2p.TCPTransportOpts{
		ListenAddr:    ":4000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tr4000 := p2p.NewTCPTransport(tcpOpts4000)

	// Listen on both ports concurrently
	go func() {
		if err := tr3000.ListenAndAccept(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := tr4000.ListenAndAccept(); err != nil {
			log.Fatal(err)
		}
	}()

	// Keep the main goroutine alive
	select {}
}
