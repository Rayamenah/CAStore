package p2p

import "net"

//RPC holds any arbitrary data that is being sent over the network
//
type RPC struct {
	From    net.Addr
	Payload []byte
}
