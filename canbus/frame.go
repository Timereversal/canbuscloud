package canbus

import "unsafe"

type Frame struct {
	ID   uint32
	Len  byte
	_    [3]byte
	Data [8]byte
}

const frameSize = unsafe.Sizeof(
	// this is a can_frame.
	struct {
		ID   uint32
		Len  byte
		_    [3]byte
		Data [8]byte
	}{},
)
