package pointer

// MakeUintPointer returns an int pointer of the given argument
func MakeUintPointer(i uint) *uint {
	return &i
}

// MakeUintPointer returns an int8 pointer of the given argument
func MakeUint8Pointer(i uint8) *uint8 {
	return &i
}

// MakeUintPointer returns an int16 pointer of the given argument
func MakeUint16Pointer(i uint16) *uint16 {
	return &i
}

// MakeUintPointer returns an int32 pointer of the given argument
func MakeUint32Pointer(i uint32) *uint32 {
	return &i
}

// MakeUintPointer returns an int64 pointer of the given argument
func MakeUint64Pointer(i uint64) *uint64 {
	return &i
}
