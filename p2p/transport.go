package p2p

//peer is the interface that represents the remote node
type Peer interface {
}

//transport is anything that handles commnication between
//the nodes in the network. can be of the form
//(TCP, UCP, websockets...)
type Transport interface {
	ListenAndAccept() error
}
