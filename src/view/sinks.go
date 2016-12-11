package view

import "github.com/golang/protobuf/proto"

/*
 * A Sink receive data from a Get call
 */
type Sink interface {
	SetBytes(bytes []byte) error
	SetString(str string) error
	SetProto(m proto.Message) error
	view() (ByteView, error)
}

type stringSink struct {
	strptr   *string
	byteview ByteView
}

/*
 * stringSink implement all functions of Sink interface
 */
func StringSink(strptr *string) Sink {
	return &stringSink{strptr: strptr}
}

func (s *stringSink) SetBytes(bytes []byte) error {
	return s.SetString(string(bytes))
}

func (s *stringSink) SetString(str string) error {
	s.byteview.bytes = nil
	s.byteview.str = str
	*s.strptr = str
	return nil
}

func (s *stringSink) SetProto(m proto.Message) error {
	bytes, err := proto.Marshal(m)
	if err != nil {
		return err
	}
	s.byteview.bytes = bytes
	*s.strptr = string(bytes)
	return nil
}

func (s *stringSink) view() (ByteView, error) {
	return s.byteview, nil
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
