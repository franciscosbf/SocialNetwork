package config

import (
	"reflect"
	"testing"
)

func TestValidTagNameParsing(t *testing.T) {
	type Dummy struct {
		F string `name:"ola"`
	}

	sf := reflect.TypeOf(&Dummy{}).Elem().Field(0)
	name, err := parseTagName(&sf)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if name != "ola" {
		t.Errorf("Expecting extract name \"ola\", got: %v", name)
	}
}

func TestInvalidTagNameParsing(t *testing.T) {
	// TODO
}

// TODO - remaining tests
