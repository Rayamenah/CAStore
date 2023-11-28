package p2p

import "errors"

var ErrINvalidHandshake = errors.New("Invalid handshake")

// HandshakeFunc
type HandshakeFunc func(any) error

func NOPHandshakeFunc(any) error { return nil }
