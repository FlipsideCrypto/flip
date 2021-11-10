package pointer

// MakeRunePointer returns a pointer to a rune of the given value
func MakeRunePointer(r rune) *rune {
	return &r
}
