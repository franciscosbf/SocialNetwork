package utils

type BucketValues = map[any]struct{}

type Set[V any] struct {
	bucket BucketValues
}

// PutValue Inserts a given value
func (s *Set[V]) PutValue(value V) {
	s.bucket[value] = struct{}{}
}

// ContainsValue returns true if contains a given value
func (s *Set[V]) ContainsValue(value V) bool {
	_, ok := s.bucket[value]

	return ok
}

// NewSet returns a new set
func NewSet[V any]() *Set[V] {
	bucket := make(BucketValues)

	return &Set[V]{bucket: bucket}
}
