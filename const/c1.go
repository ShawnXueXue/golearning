package consttest

import (
	"unsafe"
)

type pgid uint64

type page struct {
	id       pgid
	flags    uint16
	count    uint16
	overflow uint32
	ptr      uintptr
}

type leafPageElement struct {
	flags uint32
	pos   uint32
	ksize uint32
	vsize uint32
}

const pageHeaderSize = int(unsafe.Offsetof(((*page)(nil)).count))

var b = "b"
