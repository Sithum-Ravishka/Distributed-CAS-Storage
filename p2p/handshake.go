package p2p

// HandshakeFunc is 
type HandshakeFunc func(any) error

func NOPHandshakeFunc(any) error { return nil }
