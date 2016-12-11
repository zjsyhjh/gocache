package view

import "github.com/golang/protobuf/proto"

/*
 * stringSink implement all functions of Sink interface
 */
type stringSink struct {
	strptr   *string
	byteview ByteView
}

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
