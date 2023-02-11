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
