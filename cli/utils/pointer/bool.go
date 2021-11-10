package pointer

// MakeBoolPointer returns a pointer to a boolean of the given value
func MakeBoolPointer(b bool) *bool {
	return &b
}
