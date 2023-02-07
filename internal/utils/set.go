package utils

type Bucket = map[any]struct{}

type Set[V comparable] struct {
	m Bucket
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

// NewSet returns a new set
func NewSet[V comparable]() *Set[V] {
	bucket := make(Bucket)

	return &Set[V]{m: bucket}
}
