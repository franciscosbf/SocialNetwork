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

// NewSet returns a new set
func NewSet[V comparable]() *Set[V] {
	bucket := make(map[V]struct{})

	return &Set[V]{m: bucket}
}
