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
	"github.com/franciscosbf/micro-dwarf/internal/envvars"
	"github.com/franciscosbf/micro-dwarf/internal/envvars/providers"
	"github.com/franciscosbf/micro-dwarf/internal/errorw"
	"github.com/franciscosbf/micro-dwarf/internal/utils"
	"os"
	"reflect"
	"testing"
	"time"
)

func setVars(vs map[string]string) {
	for k, v := range vs {
		_ = os.Setenv(k, v)
	}
}

func unsetVars(vs ...string) {
	for _, v := range vs {
		_ = os.Unsetenv(v)
	}
}

type FakeProvider struct {
}

func (fp *FakeProvider) Get(key string) (string, error) {
	if key == "error" {
		return "", fmt.Errorf("some error")
	}

	return os.Getenv(key), nil
}

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

		f := sT.Field(i)
		fValue := sV.Field(i)

		v := &variableInfo{val: &fValue}

		if err := selectTypeConverter(v, &f); err != nil {
			t.Errorf("Unnexpected getting error in field %v: %v", f.Name, err)
		}
		if v.setValue == nil {
			t.Errorf("Expecting converter defined for field %v", f.Name)
		}

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

	fV := reflect.ValueOf(s).Elem().Field(0)

	v := &variableInfo{val: &fV}

	err := selectTypeConverter(v, &f)
	if _, ok := err.(*UnsupportedTypeError); !ok {
		t.Errorf("Expecting error UnsupportedTypeError, got: %v", err)
	}
	if v.setValue != nil {
		t.Error("Expecting nil converter")
	}
}

func TestValidTagParser(t *testing.T) {
	type Dummy struct {
		S string `name:"hello" required:"yes" accepts:"pasta,chicken"`
	}

	f := reflect.
		TypeOf(&Dummy{}).
		Elem().
		Field(0)

	v := &variableInfo{}

	fakeM := make(map[string]string)

	if err := parseFieldTagKeys(v, &f, fakeM); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if v.name != "hello" {
		t.Errorf("Expecting variable name hello, got: %v", v.name)
	}

	if !v.required {
		t.Error("Expecting variable as required")
	}

	if !(v.acceptedValues.Contains("pasta") &&
		v.acceptedValues.Contains("chicken")) {
		t.Error("Expecting accepted tokens pasta and chicken: got", v.acceptedValues.Values())
	}
}

func TestInvalidTagParser(t *testing.T) {
	checkError := func(t *testing.T, srtPtr any) {
		f := reflect.
			TypeOf(srtPtr).
			Elem().
			Field(0)

		fakeM := make(map[string]string)

		if err := parseFieldTagKeys(&variableInfo{}, &f, fakeM); err == nil {
			t.Errorf("Expecting getting an error")
		}
	}

	testBattery := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "TestWithoutName",
			test: func(t *testing.T) {
				checkError(t, &struct {
					I int `required:"YES"`
				}{})
			},
		},
		{
			name: "TestWithInvalidRequired",
			test: func(t *testing.T) {
				checkError(t, &struct {
					I int `name:"hi" required:"sa"`
				}{})
			},
		},
		{
			name: "TestWithInvalidAcceptsFmt",
			test: func(t *testing.T) {
				checkError(t, &struct {
					I int `name:"hi" accepts:","`
				}{})
			},
		},
	}

	for _, pair := range testBattery {
		t.Run(pair.name, pair.test)
	}
}

func TestRepeatedVarName(t *testing.T) {
	type Dummy struct {
		I int `name:"Joe"`
	}

	f := reflect.
		TypeOf(&Dummy{}).
		Elem().
		Field(0)

	fakeM := make(map[string]string)
	fakeM["Joe"] = ""

	err := parseFieldTagKeys(&variableInfo{}, &f, fakeM)
	if _, ok := err.(*RepeatedVarNameError); !ok {
		t.Errorf("Expecting error RepeatedVarNameError. Instead got: %v", err)
	}
}

func TestInvalidAccepts(t *testing.T) {
	ft := reflect.
		TypeOf((*int)(nil)).
		Elem()

	accepts := utils.NewSet[string]()
	accepts.Put("1")
	accepts.Put("str")

	err := validateAccepted(ft, parseInt, accepts)
	if _, ok := err.(*TypeInconsistencyError); !ok {
		t.Errorf("Expecting getting error TypeInconsistencyError, got %v", err)
	}
}

func TestValidParseFields(t *testing.T) {
	type Dummy struct {
		S string `name:"hello" required:"yes"`
		I int    `name:"bye" accepts:"1,2"`
	}

	sV := reflect.ValueOf(&Dummy{}).Elem()
	vs, err := parseFields(&sV)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	nVar := sV.NumField()
	if len(vs) != nVar {
		t.Errorf("Expecting %v parsed variables", nVar)
	}

	checkVar := func(v *variableInfo, name string, required bool, accepts ...string) {
		if name != v.name {
			t.Errorf("Invalid variable name %v. Was expecting %v", v.name, name)
			return
		}

		if required != v.required {
			t.Errorf("Expecting required to be %v of variable %v", required, name)
			return
		}

		for _, a := range accepts {
			if !v.isValidKeyword(a) {
				t.Errorf("Accepted keyword %v is missing in variable %v", a, name)
			}
		}
	}

	// Struct fields are evaluated from top to bottom
	checkVar(vs[0], "hello", true)
	checkVar(vs[1], "bye", false, "1", "2")
}

func TestInvalidParseFields(t *testing.T) {
	testBattery := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "TestEmptyStruct",
			test: func(t *testing.T) {
				st := reflect.ValueOf(&struct{}{}).Elem()
				vs, err := parseFields(&st)
				if err != WithoutFieldsError {
					t.Errorf("Expecting error WithoutFieldsError, got: %v", err)
				}

				if vs != nil {
					t.Errorf("Expecting nil slice, got: %v", vs)
				}
			},
		},
		{
			name: "TestAnonymousField",
			test: func(t *testing.T) {
				st := reflect.ValueOf(&struct{ int }{}).Elem()
				vs, err := parseFields(&st)
				if err, ok := err.(*AnonymousFieldError); !ok {
					t.Errorf("Expecting error AnonymousFieldError, got: %v", err)
				}

				if vs != nil {
					t.Errorf("Expecting nil slice, got: %v", vs)
				}
			},
		},
		{
			name: "TestStructWithOnlyPrivateFields",
			test: func(t *testing.T) {
				st := reflect.ValueOf(&struct {
					I int    `name:"hello"`
					s string `name:"hi"`
				}{}).Elem()
				vs, err := parseFields(&st)
				if _, ok := err.(*PrivateFieldError); !ok {
					t.Errorf("Expecting error PrivateFieldError, got: %v", err)
				}

				if vs != nil {
					t.Errorf("Expecting nil slice, got: %v", vs)
				}
			},
		},
		{
			name: "TestWithMissingTag",
			test: func(t *testing.T) {
				st := reflect.ValueOf(&struct {
					I int    `name:"joe"`
					S string `nae:"hi"`
				}{}).Elem()
				vs, err := parseFields(&st)
				if _, ok := err.(*MissingTagKeyError); !ok {
					t.Errorf("Expecting error MissingTagKeyError, got: %v", err)
				}

				if vs != nil {
					t.Errorf("Expecting nil slice, got: %v", vs)
				}
			},
		},
	}

	for _, pair := range testBattery {
		t.Run(pair.name, pair.test)
	}
}

func TestValidParseConf(t *testing.T) {
	type Dummy struct {
		I int           `name:"var1" required:"yes" accepts:"1,2,3"`
		S string        `name:"var2" required:"no"`
		T time.Duration `name:"var3" required:"yes" accepts:"1s,2h3m"`
	}

	ev := providers.NewEnvVariables()
	c := envvars.New(ev)

	cp, err := New(c)
	if err != nil {
		t.Errorf("Unexptected error from conf parser initializer: %v", err)
		return
	}

	d := &Dummy{}

	setVars(map[string]string{
		"var1": "1",
		"var2": "ola",
		"var3": "1s",
	})
	defer unsetVars("var1", "var2", "var3")

	if err := cp.ParseConf(d); err != nil {
		t.Errorf("Unexpected error from config parser: %v", err)
		return
	}

	if d.I != 1 {
		t.Errorf("Expecting I field with 1, got %v", d.I)
	}
	if d.S != "ola" {
		t.Errorf("Expecting I field with ola, got %v", d.S)
	}
	expectedT, _ := time.ParseDuration("1s")
	if d.T != expectedT {
		t.Errorf("Expecting T field with 1s, got %v", d.T)
	}
}

func TestInvalidParseConf(t *testing.T) {
	checkErrorCode := func(t *testing.T, srtPtr any, code errorw.ErrorCode, errName string) {
		c := envvars.New(&FakeProvider{})

		cp, _ := New(c)
		pErr := cp.ParseConf(srtPtr)
		if pErr == nil {
			t.Error("Expecting error, got nil")
		} else if err, ok := pErr.(*errorw.Wrapper); !ok {
			t.Errorf("Expecting error of type errorw.Wrapper, got %v", err)
		} else {
			if err.Code() != code {
				t.Errorf("Expecting error code %v, got %v", errName, err.String())
			}
		}
	}

	testBattery := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "TestInvalidConf",
			test: func(t *testing.T) {
				c := envvars.New(&FakeProvider{})

				cp, _ := New(c)
				pErr := cp.ParseConf(struct{}{})
				if pErr != InvalidPointerError {
					t.Errorf("Expecting error InvalidPointerError, got: %v", pErr)
				}
			},
		},
		{
			name: "TestInvalidField",
			test: func(t *testing.T) {
				c := envvars.New(&FakeProvider{})

				cp, _ := New(c)
				pErr := cp.ParseConf(&struct {
					i int `name:"hello"`
				}{})
				if _, ok := pErr.(*PrivateFieldError); !ok {
					t.Errorf("Expecting error of type PrivateFieldError, got: %v", pErr)
				}
			},
		},
		{
			name: "TestGetVariableError",
			test: func(t *testing.T) {
				checkErrorCode(t, &struct {
					L string `name:"error"`
				}{}, ErrorCodeInvalidGetVar, "ErrorCodeInvalidGetVar")
			},
		},
		{
			name: "TestUnsetVariable",
			test: func(t *testing.T) {
				checkErrorCode(t, &struct {
					I string `name:"aa" required:"true"`
				}{}, ErrorCodeMissingVar, "ErrorCodeMissingVar")
			},
		},
		{
			name: "TestInvalidKeyword",
			test: func(t *testing.T) {
				setVars(map[string]string{
					"aa": "a",
					"bb": "c",
				})
				defer unsetVars("aa", "bb")
				checkErrorCode(t, &struct {
					I string `name:"aa"`
					J string `name:"bb" accepts:"hello,bye"`
				}{}, ErrorCodeUnacceptedVal, "ErrorCodeUnacceptedVal")
			},
		},
		{
			name: "TestInvalidType",
			test: func(t *testing.T) {
				setVars(map[string]string{
					"aa": "1",
				})
				defer unsetVars("aa")
				checkErrorCode(t, &struct {
					I time.Duration `name:"aa"`
				}{}, ErrorCodeInvalidVarType, "ErrorCodeInvalidVarType")
			},
		},
	}

	for _, pair := range testBattery {
		t.Run(pair.name, pair.test)
	}
}
