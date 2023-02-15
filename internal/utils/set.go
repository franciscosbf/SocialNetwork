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

type Set[V comparable] struct {
	m map[V]struct{}
}

// Put Inserts a given value
func (s *Set[V]) Put(value V) {
	s.m[value] = struct{}{}
}

// Contains returns true if contains a given value
func (s *Set[V]) Contains(value V) bool {
	_, ok := s.m[value]

	return ok
}

// Values returns all values
func (s *Set[V]) Values() []V {
	values := make([]V, len(s.m))

	c := 0
	for v := range s.m {
		values[c] = v
		c++
	}

	return values
}

// Copy returns a new set with the same values.
// Warning: it doesn't do deep copy of values
func (s *Set[V]) Copy() (newS *Set[V]) {
	newS = NewSet[V]()

	for v := range s.m {
		newS.Put(v)
	}

	return
}

// Size returns the number
// of stored values
func (s *Set[V]) Size() int {
	return len(s.m)
}

func (s *Set[V]) Empty() bool {
	return s.Size() == 0
}

// NewSet returns a new set
func NewSet[V comparable]() *Set[V] {
	bucket := make(map[V]struct{})

	return &Set[V]{m: bucket}
}
