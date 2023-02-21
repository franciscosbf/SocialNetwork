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

func TestSetLen(t *testing.T) {
	s := NewSet[string]()

	if s.Size() != 0 {
		t.Error("Expecting set with zero size")
	}

	values := []string{
		"a", "b", "c", "", " ",
	}

	for _, v := range values {
		s.Put(v)
	}

	if s.Size() != len(values) {
		t.Errorf("Expecting set to have size %v", len(values))
	}
}

func TestEmptySet(t *testing.T) {
	s := NewSet[string]()

	if !s.Empty() {
		t.Errorf("Expecting empty set")
	}

	s.Put("a")

	if s.Empty() {
		t.Errorf("Empty call should have returned false on set with one element")
	}
}

func TestSetDelete(t *testing.T) {
	s := NewSet[string]()

	s.Put("a")
	s.Put("b")

	s.Delete("a")
	s.Delete("b")

	if !s.Empty() {
		t.Errorf("Expecting empty set, contains %v", s.Values())
	}
}
