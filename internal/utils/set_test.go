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
