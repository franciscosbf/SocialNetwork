package utils

import "reflect"

// AnyPtr represents the location
// where the value will be assigned
type AnyPtr any

// SetAny assigns a value to a given pointer if
// from isn't zero/nil. If to isn't a pointer it
// panics. Moreover, if from isn't an interface
// value resulted of a pass-by-value, it also panics
func SetAny(from any, to AnyPtr) {
	fromV := reflect.ValueOf(from)

	// Check if preset
	if !fromV.IsZero() {
		reflect.
			ValueOf(to).
			Elem().
			Set(fromV)
	}
}
