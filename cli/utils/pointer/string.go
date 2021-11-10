package pointer

// MakeStringPointer returns a pointer to a string of the given value
func MakeStringPointer(s string) *string {
	return &s
}
