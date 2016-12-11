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
	*s.dst = ByteView{bytes: cloneBytes(bytes)}
	return nil
}

func (s *byteViewSink) SetString(str string) error {
	*s.dst = ByteView{str: str}
	return nil
}

func (s *byteViewSink) SetProto(m proto.Message) error {
	bytes, err := proto.Marshal(m)
	if err != nil {
		return err
	}
	*s.dst = ByteView{bytes: bytes}
	return nil
}

func (s *byteViewSink) view() (ByteView, error) {
	return *s.dst, nil
}

func (s *byteViewSink) setView(view ByteView) error {
	*s.dst = view
	return nil
}
