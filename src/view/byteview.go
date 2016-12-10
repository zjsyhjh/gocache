package view

import (
	"bytes"
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
