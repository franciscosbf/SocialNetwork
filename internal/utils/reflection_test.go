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

func TestValidSetAny(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("Unexpected panic: %v", err)
		}
	}()

	var i int

	SetAny(1, &i)

}

func TestInvalidSetAny(t *testing.T) {
	catchPanic := func(t *testing.T, from any, to AnyPtr) {
		defer func() {
			if err := recover(); err == nil {
				t.Error("Expecting panic")
			}
		}()

		SetAny(from, to)

	}

	testBattery := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "TestInvalidFrom",
			test: func(t *testing.T) {
				var j string
				l := "hello"
				catchPanic(t, &l, &j)
			},
		},
		{
			name: "TestInvalidTo",
			test: func(t *testing.T) {
				var j string
				catchPanic(t, "hello", j)
			},
		},
	}

	for _, pair := range testBattery {
		t.Run(pair.name, pair.test)
	}
}
