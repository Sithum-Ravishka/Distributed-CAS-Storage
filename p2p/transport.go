package p2p

// Peer is an interface the represent the remote node.
type Peer interface {
}

// Transport is anything that handles the communication
// between the nodes in the network. This can be of the
// from (TCP, UDP, websockets, ...)
type Transport interface {
}
