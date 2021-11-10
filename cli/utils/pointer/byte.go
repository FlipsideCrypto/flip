package pointer

// MakeBytePointer returns a pointer to a byte of the given value
func MakeBytePointer(b byte) *byte {
	return &b
}
