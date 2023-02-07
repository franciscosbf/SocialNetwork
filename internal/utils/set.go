package utils

type Set[V comparable] struct {
	m map[V]struct{}
}

// PutValue Inserts a given value
func (s *Set[V]) PutValue(value V) {
	s.m[value] = struct{}{}
}

// ContainsValue returns true if contains a given value
func (s *Set[V]) ContainsValue(value V) bool {
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

// NewSet returns a new set
func NewSet[V comparable]() *Set[V] {
	bucket := make(map[V]struct{})

	return &Set[V]{m: bucket}
}
