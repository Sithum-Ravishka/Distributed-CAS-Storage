package p2p

// Message holds any arbitray data that is being sent over the
// each transport between two nodes in the network.
type Message struct {
	Payload []byte
}
