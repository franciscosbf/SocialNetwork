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
	"time"
)

func TestValidStringParsing(t *testing.T) {
	var s string

	v := reflect.ValueOf(&s).Elem()
	if err := parseStringType.converter(&v, "hello"); err != nil {
		t.Errorf("Unexptected error %v", err)
	}

	if s != "hello" {
		t.Errorf("Expecting assign value \"hello\", got: %v", s)
	}
}

func TestValidIntegerParsing(t *testing.T) {
	var i int

	v := reflect.ValueOf(&i).Elem()
	if err := parseIntegerType.converter(&v, "212"); err != nil {
		t.Errorf("Unexptected error %v", err)
	}

	if i != 212 {
		t.Errorf("Expecting assign value \"212\", got: %v", i)
	}
}

func TestInvalidIntegerParsing(t *testing.T) {
	var i int

	v := reflect.ValueOf(&i).Elem()
	if err := parseIntegerType.converter(&v, "212s"); err == nil {
		t.Error("Expecting getting an error")
	}
}

func TestValidDurationParsing(t *testing.T) {
	var d time.Duration

	v := reflect.ValueOf(&d).Elem()
	if err := parseDurationType.converter(&v, "1h3m"); err != nil {
		t.Errorf("Unexptected error %v", err)
	}

	expected, _ := time.ParseDuration("1h3m")
	if d != expected {
		t.Errorf("Expecting assign value \"%v\", got: %v", expected, d)
	}
}

func TestInvalidDurationParsing(t *testing.T) {
	var d time.Duration

	v := reflect.ValueOf(&d).Elem()
	if err := parseDurationType.converter(&v, "1h2mm"); err == nil {
		t.Error("Expecting getting an error")
	}
}
