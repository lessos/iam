// Code generated by github.com/hooto/protobuf_slice
// DO NOT EDIT!

package iamapi

import "bytes"

func PbBytesSliceEqual(ls, ls2 []byte) bool {
	if len(ls) != len(ls2) {
		return false
	}
	return bytes.Compare(ls, ls2) == 0
}

func PbArrayBytesSliceEqual(ls, ls2 [][]byte) bool {
	if len(ls) != len(ls2) {
		return false
	}
	for i, v := range ls {
		if !PbBytesSliceEqual(v, ls2[i]) {
			return false
		}
	}
	return true
}