package config

import (
	"reflect"
	"testing"
	"time"
)

func TestExtractionOfValidStruct(t *testing.T) {
	type DummyStruct struct {
		A int
	}

	sV, err := extractStrVal(&DummyStruct{})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if sV.Kind() != reflect.Struct {
		t.Errorf("Expecting %v of type struct", sV)
	}
}

func TestExtractionOfNonPointer(t *testing.T) {
	_, err := extractStrVal(3)
	if err != InvalidPointerError {
		t.Errorf("Expecting error %v", InvalidPointerError)
	}
}

func TestExtractionOfNilStruct(t *testing.T) {
	_, err := extractStrVal((*struct{})(nil))
	if err != InvalidValuePointedError {
		t.Errorf("Expecting error %v", InvalidValuePointedError)
	}
}

func TestValidTypeConverter(t *testing.T) {
	type DummyStruct struct {
		A int           `v:"1"`
		B string        `v:"hello"`
		C time.Duration `v:"1s"`
	}

	s := &DummyStruct{}

	sT := reflect.TypeOf(s).Elem()
	sV := reflect.ValueOf(s).Elem()

	type fieldCheck struct {
		tC  typeConverter
		v   *reflect.Value
		f   *reflect.StructField
		raw string
	}

	var fChecks []*fieldCheck

	for i := 0; i < sT.NumField(); i++ {
		v := &variableInfo{}

		f := sT.Field(i)

		if err := selectTypeConverter(v, &f); err != nil {
			t.Errorf("Unnexpected getting error in field %v: %v", f.Name, err)
		}
		if v.setValue == nil {
			t.Errorf("Expecting converter defined for field %v", f.Name)
		}

		fValue := sV.Field(i)

		fChecks = append(fChecks, &fieldCheck{v.setValue, &fValue, &f, f.Tag.Get("v")})
	}

	tryAssignment := func(fc *fieldCheck) {
		err := fc.tC(fc.v, fc.raw)
		if err != nil {
			t.Errorf(
				"Unexpected error while trying to covert value %v on field %v: %v",
				fc.raw, fc.f.Name, err)
		}

		defer func() {
			if err := recover(); err != nil {
				t.Errorf(
					"Expecting valid assignment of value %v on field %v. got panic: %v",
					fc.raw, fc.f.Name, err)
			}

		}()
	}

	for _, check := range fChecks {
		tryAssignment(check)
	}
}

func TestInvalidTypeConverter(t *testing.T) {
	type DummyStruct struct {
		A struct{}
	}

	s := &DummyStruct{}

	sT := reflect.TypeOf(s).Elem()
	f := sT.Field(0)

	v := &variableInfo{}

	err := selectTypeConverter(v, &f)
	if _, ok := err.(*UnsupportedTypeError); !ok {
		t.Errorf("Expecting error UnsupportedTypeError, got: %v", err)
	}
	if v.setValue != nil {
		t.Error("Expecting nil converter")
	}
}

func TestValidTagParser(t *testing.T) {
	// TODO
}

func TestInvalidTagParser(t *testing.T) {
	// TODO
}

// TODO - remaining tests
