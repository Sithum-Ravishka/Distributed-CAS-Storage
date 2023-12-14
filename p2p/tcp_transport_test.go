package p2p

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestTCPTransport(t *testing.T) {
// 	listenAddr := ":4000"
// 	tr := NewTCPTransport(listenAddr)

// 	assert.Equal(t, tr.listenAddress, listenAddr)

// 	// Server
// 	// tr.Start()
// 	assert.Nil(t, tr.ListenAndAccept())

// 	select {}

// }

// Node 2
// package main

// import (
// 	"cas-storage/p2p"
// 	"log"
// 	"time"
// )

// func main() {
// 	// Node 2 configuration
// 	node2Addr := ":3001"
// 	tcpOpts := p2p.TCPTransportOpts{
// 		ListenAddr:    node2Addr,
// 		HandshakeFunc: p2p.NOPHandshakeFunc,
// 		Decoder:       p2p.DefaultDecoder{},
// 	}
// 	tr := p2p.NewTCPTransport(tcpOpts)

// 	if err := tr.ListenAndAccept(); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Node 1 configuration
// 	node1Addr := "localhost:3000"
// 	tr2 := p2p.NewTCPTransport(p2p.TCPTransportOpts{
// 		ListenAddr:    node1Addr,
// 		HandshakeFunc: p2p.NOPHandshakeFunc,
// 		Decoder:       p2p.DefaultDecoder{},
// 	})

// 	// Retry until Node 1 is reachable
// 	for {
// 		err := tr2.DialAndConnect(node1Addr)
// 		if err == nil {
// 			break
// 		}
// 		log.Printf("Connection to Node 1 failed, retrying in 2 seconds: %v", err)
// 		time.Sleep(2 * time.Second)
// 	}

// 	select {}
// }
