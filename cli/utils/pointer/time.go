package pointer

import "time"

func MakeTimePointer(t time.Time) *time.Time {
	return &t
}
