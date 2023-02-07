package utils

import "testing"

func TestSetContains(t *testing.T) {
	s := NewSet[string]()

	values := []string{
		"a", "b", "c", "", " ",
	}

	for _, v := range values {
		s.Put(v)
	}

	for _, v := range values {
		if !s.Contains(v) {
			t.Errorf("Missing value '%v' in set", v)
		}
	}

}

func TestSetValues(t *testing.T) {
	s := NewSet[string]()

	values := map[string]struct{}{
		"a": {}, "b": {}, "c": {}, "": {}, " ": {},
	}

	for v := range values {
		s.Put(v)
	}

	for _, v := range s.Values() {
		if _, ok := values[v]; !ok {
			t.Errorf("Missing value '%v' in set", v)
		}
	}
}

func TestSetCopy(t *testing.T) {
	type dummy struct {
		i int
	}

	s := NewSet[*dummy]()

	values := []*dummy{
		{i: 1}, {i: 2}, {i: 3}, {i: 4}, {i: 5},
	}

	for _, v := range values {
		s.Put(v)
	}

	newS := s.Copy()

	for _, v := range s.Values() {
		if !newS.Contains(v) {
			t.Errorf("Missing value '%v' in set", v)
		}
	}
}
