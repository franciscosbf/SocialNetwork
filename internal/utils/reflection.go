/*
Copyright 2023 Francisco Simões Braço-Forte

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
