package view

import (
	"bytes"
	"errors"
	"io"
	"strings"
)

/*
 * 对byte数组以及string的包装
 * 如果bytes != nil, 则使用bytes，否则使用str
 */
type ByteView struct {
	bytes []byte
	str   string
}

/*
 * return a copy of the data as a byte slice
 */
func (view ByteView) ByteSlice() []byte {
	if view.bytes != nil {
		return cloneBytes(view.bytes)
	}
	return []byte(view.str)
}

/*
 * return the data as string
 */
func (view ByteView) String() string {
	if view.bytes != nil {
		return string(view.bytes)
	}
	return view.str
}

/*
 * return the byte by the provided index
 */
func (view ByteView) At(i int) byte {
	if view.bytes != nil {
		return view.bytes[i]
	}
	return view.str[i]
}

/*
 * return the length of ByteView
 */
func (view ByteView) Len() int {
	if view.bytes != nil {
		return len(view.bytes)
	}
	return len(view.str)
}

/*
 * return slices the view
 */
func (view ByteView) Slice(from, to int) ByteView {
	if view.bytes != nil {
		return ByteView{bytes: view.bytes[from:to]}
	}
	return ByteView{str: view.str[from:to]}
}

/*
 * return slices the view from the provided index
 */
func (view ByteView) SliceFrom(from int) ByteView {
	if view.bytes != nil {
		return ByteView{bytes: view.bytes[from:]}
	}
	return ByteView{str: view.str[from:]}
}

func (view ByteView) Copy(dst []byte) int {
	if view.bytes != nil {
		return copy(dst, view.bytes)
	}
	return copy(dst, view.str)
}

/*
 * return ture when the bytes in view are the same as bytes in other
 */
func (view ByteView) Equals(other ByteView) bool {
	if other.bytes != nil {
		return view.equalBytes(other.bytes)
	}
	return view.equalString(other.str)
}

func (view ByteView) equalBytes(b []byte) bool {
	if view.bytes != nil {
		return bytes.Equal(view.bytes, b)
	}
	if len(view.str) != len(b) {
		return false
	}
	for i, bi := range b {
		if bi != view.str[i] {
			return false
		}
	}
	return true
}

func (view ByteView) equalString(s string) bool {
	if view.bytes == nil {
		return view.str == s
	}
	if len(view.bytes) != len(s) {
		return false
	}
	for i, bi := range view.bytes {
		if bi != s[i] {
			return false
		}
	}
	return true
}

/*
 * return io.ReadSeeker
 */
func (view ByteView) Reader() io.ReadSeeker {
	if view.bytes != nil {
		return bytes.NewReader(view.bytes)
	}
	return strings.NewReader(view.str)
}

/*
 * ReadAt implements io.ReadAt on the bytes in view
 */
func (view ByteView) ReadAt(b []byte, offset int64) (n int, err error) {
	if offset < 0 {
		return 0, errors.New("invalid offset")
	}
	if offset >= int64(view.Len()) {
		return 0, io.EOF
	}
	n = view.SliceFrom(int(offset)).Copy(b)
	if n < len(b) {
		err = io.EOF
	}
	return
}
