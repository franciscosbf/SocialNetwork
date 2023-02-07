package utils

import "testing"

func TestSetContains(t *testing.T) {
	s := NewSet[string]()

	values := []string{
		"a", "b", "c", "", " ",
	}

	for _, v := range values {
		s.PutValue(v)
	}

	for _, v := range values {
		if !s.ContainsValue(v) {
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
		s.PutValue(v)
	}

	for _, v := range s.Values() {
		if _, ok := values[v]; !ok {
			t.Errorf("Missing value '%v' in set", v)
		}
	}
}
