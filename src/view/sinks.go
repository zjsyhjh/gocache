package view

import (
	"github.com/golang/protobuf/proto"
)

/*
 * A Sink receive data from a Get call
 */
type Sink interface {
	SetBytes(bytes []byte) error
	SetString(s string) error
	SetProto(m proto.Message) error
	view() (ByteView, error)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
