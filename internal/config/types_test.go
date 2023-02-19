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
	"fmt"
	"github.com/franciscosbf/micro-dwarf/internal/utils"
	"math"
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

func TestValidInteger32Parsing(t *testing.T) {
	var i int32

	v := reflect.ValueOf(&i).Elem()
	if err := parseInteger32Type.converter(&v, fmt.Sprintf("%v", math.MaxInt32)); err != nil {
		t.Errorf("Unexptected error %v", err)
	}

	if i != math.MaxInt32 {
		t.Errorf("Expecting assign value \"%v\", got: %v", math.MaxInt32, i)
	}
}

func TestInvalidInteger32Parsing(t *testing.T) {
	var i int32

	v := reflect.ValueOf(&i).Elem()
	if err := parseInteger32Type.converter(&v, fmt.Sprintf("%v", math.MaxInt)); err == nil {
		t.Error("Expecting getting an error")
	}
}

func TestValidUnsignedInteger16Parsing(t *testing.T) {
	var i uint16

	v := reflect.ValueOf(&i).Elem()
	if err := parseUnsignedInteger16Type.converter(&v, fmt.Sprintf("%v", math.MaxUint16)); err != nil {
		t.Errorf("Unexptected error %v", err)
	}

	if i != math.MaxUint16 {
		t.Errorf("Expecting assign value \"%v\", got: %v", math.MaxUint16, i)
	}
}

func TestInvalidUnsignedInteger16Parsing(t *testing.T) {
	var i uint16

	v := reflect.ValueOf(&i).Elem()
	if err := parseUnsignedInteger16Type.converter(&v, fmt.Sprintf("%v", math.MaxInt)); err == nil {
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

func TestValidBoolParsing(t *testing.T) {
	var b bool

	v := reflect.ValueOf(&b).Elem()
	if err := parseBoolType.converter(&v, "true"); err != nil {
		t.Errorf("Unexptected error %v", err)
	}

	if !b {
		t.Errorf("Expecting true")
	}
}

func TestInvalidBoolParsing(t *testing.T) {
	var b bool

	v := reflect.ValueOf(&b).Elem()
	if err := parseBoolType.converter(&v, "lol"); err == nil {
		t.Error("Expecting getting an error")
	}
}

func TestValidAddrsRefParsing(t *testing.T) {
	var aP *utils.Addrs

	v := reflect.ValueOf(&aP).Elem()
	if err := parseAddrsRefType.converter(&v, "localhost:123;lo.com:999"); err != nil {
		t.Errorf("Unexptected error %v", err)
	}

	if aP == nil {
		t.Error("Nil addrs pointer")
		return
	}

	if len(aP.Bucket) == 0 {
		t.Error("Doesn't contain any addr")
		return
	}

	if len(aP.Bucket) != 2 {
		t.Errorf("Contains more than the expected: %v", aP.Bucket)
		return
	}

	first := fmt.Sprintf("%v:%v", aP.Bucket[0].Host, aP.Bucket[0].Port)
	if first != "localhost:123" {
		t.Errorf("First position doesn't contain localhost:123, got %v", first)
	}

	second := fmt.Sprintf("%v:%v", aP.Bucket[1].Host, aP.Bucket[1].Port)
	if second != "lo.com:999" {
		t.Errorf("First position doesn't contain lo.com:999, got %v", second)
	}
}

func TestInvalidAddrsRefParsing(t *testing.T) {
	var aP *utils.Addrs

	v := reflect.ValueOf(&aP).Elem()
	if err := parseAddrsRefType.converter(&v, "lol"); err == nil {
		t.Error("Expecting getting an error")
	}
}
