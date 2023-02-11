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

package config

import (
	"reflect"
	"testing"
)

func TestValidTagKeyNameParsing(t *testing.T) {
	type Dummy struct {
		F string `name:"ola"`
	}

	sf := reflect.TypeOf(&Dummy{}).Elem().Field(0)
	name, err := parseTagKeyName(&sf)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if name != "ola" {
		t.Errorf("Expecting extract name \"ola\", got: %v", name)
	}
}

func TestInvalidTagKeyNameParsing(t *testing.T) {
	checkError := func(t *testing.T, srtPtr any) {
		sf := reflect.TypeOf(srtPtr).Elem().Field(0)

		name, err := parseTagKeyName(&sf)
		if _, ok := err.(*MissingTagKeyError); !ok {
			t.Errorf("Expecting error MissingTagKeyError, got: %v", err)
		}

		if name != "" {
			t.Errorf("Expecting empty returned string, got %v", name)
		}
	}

	testBattery := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "TestMissingTagName",
			test: func(t *testing.T) {
				checkError(t, &struct {
					F string ``
				}{})
			},
		},
		{
			name: "TestWithInvalidValue",
			test: func(t *testing.T) {
				checkError(t, &struct {
					F string `name:"`
				}{})
			},
		},
		{
			name: "TestWithEmptyValue",
			test: func(t *testing.T) {
				checkError(t, &struct {
					F string `name:""`
				}{})
			},
		},
	}

	for _, pair := range testBattery {
		t.Run(pair.name, pair.test)
	}
}

func TestValidTagKeyRequiredParsing(t *testing.T) {
	checkReturn := func(t *testing.T, required bool, srtPtr any) {
		sf := reflect.TypeOf(srtPtr).Elem().Field(0)

		requiresField, err := parseTagKeyRequired(&sf)
		if err != nil {
			t.Errorf("Expecting nil err, got %v in tag %v", err, sf.Tag)
		}

		if requiresField != required {
			t.Errorf("Expecting required to be %v in tag %v", required, sf.Tag)
		}
	}

	testBattery := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "TestWithTrue",
			test: func(t *testing.T) {
				checkReturn(t, true, &struct {
					I int `required:"tRUe"`
				}{})
			},
		},
		{
			name: "TestWithYes",
			test: func(t *testing.T) {
				checkReturn(t, true, &struct {
					I int `required:"YeS"`
				}{})
			},
		},
		{
			name: "TestWithFalse",
			test: func(t *testing.T) {
				checkReturn(t, false, &struct {
					I int `required:"FalSE"`
				}{})
			},
		},
		{
			name: "TestWithNo",
			test: func(t *testing.T) {
				checkReturn(t, false, &struct {
					I int `required:"nO"`
				}{})
			},
		},
		{
			name: "TestEmpty",
			test: func(t *testing.T) {
				checkReturn(t, false, &struct {
					I int
				}{})
			},
		},
	}

	for _, pair := range testBattery {
		t.Run(pair.name, pair.test)
	}
}

func TestInvalidTagKeyRequiredParsing(t *testing.T) {
	checkError := func(t *testing.T, srtPtr any) {
		sf := reflect.TypeOf(srtPtr).Elem().Field(0)

		b, err := parseTagKeyRequired(&sf)
		if _, ok := err.(*InvalidTagKeyValueError); !ok {
			t.Errorf("Expecting error InvalidTagKeyValueError, got: %v", err)
		}

		if b {
			t.Errorf("Expecting false by default when an error occurrs")
		}
	}

	testBattery := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "TestWithEmptyValue",
			test: func(t *testing.T) {
				checkError(t, &struct {
					F string `required:""`
				}{})
			},
		},
		{
			name: "TestWithOtherValue",
			test: func(t *testing.T) {
				checkError(t, &struct {
					F string `required:"fsa"`
				}{})
			},
		},
	}

	for _, pair := range testBattery {
		t.Run(pair.name, pair.test)
	}
}

func TestValidTagKeyAcceptsParsing(t *testing.T) {
	checkReturn := func(t *testing.T, accepts []string, srtPtr any) {
		sf := reflect.TypeOf(srtPtr).Elem().Field(0)

		vs, err := parseTagKeyAccepts(&sf)
		if err != nil {
			t.Errorf("Expecting nil err, got %v", err)
		}

		for _, v := range accepts {
			if !vs.Contains(v) {
				t.Errorf("Missing accepted value %v", v)
			}
		}
	}

	testBattery := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "TestWithoutTag",
			test: func(t *testing.T) {
				checkReturn(t, []string{}, &struct {
					S int
				}{})
			},
		},
		{
			name: "TestWithOneValue",
			test: func(t *testing.T) {
				checkReturn(t, []string{"one"}, &struct {
					S int `accepts:"one"`
				}{})
			},
		},
		{
			name: "TestWithSomeValues1",
			test: func(t *testing.T) {
				checkReturn(t, []string{"one", "two"}, &struct {
					S int `accepts:"one,two"`
				}{})
			},
		},
		{
			name: "TestWithSomeValues2",
			test: func(t *testing.T) {
				checkReturn(t, []string{"one", "two", "three"}, &struct {
					S int `accepts:"one,two,three"`
				}{})
			},
		},
	}

	for _, pair := range testBattery {
		t.Run(pair.name, pair.test)
	}
}

func TestInvalidTagKeyAcceptsParsing(t *testing.T) {
	checkInvalidFormat := func(t *testing.T, srtPtr any) {
		sf := reflect.TypeOf(srtPtr).Elem().Field(0)

		vs, err := parseTagKeyAccepts(&sf)
		if _, ok := err.(*InvalidTagKeyValueFmtError); !ok {
			t.Errorf("Expecting err InvalidTagKeyValueFmtError, got %v", err)
		}

		if vs != nil {
			t.Errorf("Expecting nil return, got %v", vs)
		}
	}

	testBattery := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "TestWithEmptyValue",
			test: func(t *testing.T) {
				sf := reflect.TypeOf(&struct {
					I int `accepts:""`
				}{}).Elem().Field(0)

				vs, err := parseTagKeyAccepts(&sf)
				if _, ok := err.(*InvalidTagKeyValueError); !ok {
					t.Errorf("Expecting err InvalidTagKeyValueError, got %v", err)
				}

				if vs != nil {
					t.Errorf("Expecting nil return, got %v", vs)
				}
			},
		},
		{
			name: "TestWithInvalidSuffix",
			test: func(t *testing.T) {
				checkInvalidFormat(t, &struct {
					I int `accepts:",a,d,b"`
				}{})
			},
		},
		{
			name: "TestWithInvalidPrefix",
			test: func(t *testing.T) {
				checkInvalidFormat(t, &struct {
					I int `accepts:"a,d,b,"`
				}{})
			},
		},
		{
			name: "TestWithInvalidTokens",
			test: func(t *testing.T) {
				checkInvalidFormat(t, &struct {
					I int `accepts:"a,,,e"`
				}{})
			},
		},
	}

	for _, pair := range testBattery {
		t.Run(pair.name, pair.test)
	}
}
