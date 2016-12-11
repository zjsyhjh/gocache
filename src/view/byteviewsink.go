package view

import "github.com/golang/protobuf/proto"

/*
 * byteViewSink
 */
type byteViewSink struct {
	dst *ByteView
}

func ByteViewSink(dst *ByteView) Sink {
	if dst == nil {
		panic("dst nil")
	}
	return &byteViewSink{dst: dst}
}

func (s *byteViewSink) SetBytes(bytes []byte) error {

}

func (s *byteViewSink) SetString(str string) error {

}

func (s *byteViewSink) SetProto(m proto.Message) error {

}

func (s *byteViewSink) view() (ByteView, error) {

}
