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
	"errors"
	"fmt"
	"strings"
)

var MissingVariablesReader = errors.New("received nil varReader")
var WithoutFieldsError = errors.New("struct doesn't have any field")
var InvalidPointerError = errors.New("expecting pointer to struct")
var InvalidValuePointedError = errors.New("expecting pointer referencing a non-nil struct")

// PrivateFieldError represents a struct field
// that is private which means is impossible
// to assign a value to it
type PrivateFieldError struct {
	fieldName string
}

func (e *PrivateFieldError) Error() string {
	return fmt.Sprintf(
		"field %v is private. It's impossible to assign a value", e.fieldName)
}

// MissingTagKeyError represents a struct field
// with a given missing key in its tag
type MissingTagKeyError struct {
	fieldName string
	keyName   string
}

func (e *MissingTagKeyError) Error() string {
	return fmt.Sprintf(
		"missing tag %v in struct field %v",
		e.keyName, e.fieldName)
}

// UnsupportedTypeError represents a type
// that the parser couldn't recognize
type UnsupportedTypeError struct {
	fieldName string
	typeName  string
}

func (e *UnsupportedTypeError) Error() string {
	return fmt.Sprintf(
		"struct field %v contains unsupported type %v",
		e.fieldName, e.typeName)
}

// TypeInconsistencyError represents an accepted
// value that doesn't match the field's type
type TypeInconsistencyError struct {
	fieldName string
	typeName  string
	rawValue  string
}

func (e *TypeInconsistencyError) Error() string {
	return fmt.Sprintf(
		"struct field %v has type %v which "+
			"doesn't match the accepted value %v",
		e.fieldName, e.typeName, e.rawValue)
}

// RepeatedVarNameError represents a variable's name
// that was already assigned to another field
type RepeatedVarNameError struct {
	assignedFieldName string
	varName           string
}

func (e *RepeatedVarNameError) Error() string {
	return fmt.Sprintf(
		"variable %v was already assigned to struct field %v",
		e.varName, e.assignedFieldName)
}

// InvalidTagKeyValueError represents
// an unrecognized or missing value
type InvalidTagKeyValueError struct {
	fieldName      string
	keyName        string
	acceptedValues []string
}

func (e *InvalidTagKeyValueError) Error() string {
	return fmt.Sprintf(
		"invalid value in tag key %v of struct field %v; accepted values %v",
		strings.Join(e.acceptedValues, ","), e.keyName, e.fieldName)
}

// InvalidTagKeyValueFmtError represents
// a bad formatted value
type InvalidTagKeyValueFmtError struct {
	fieldName string
	keyName   string
	rawValue  string
	reason    string
}

func (e *InvalidTagKeyValueFmtError) Error() string {
	return fmt.Sprintf(
		"struct field %v contains invalid value \"%v\" in tag key %v: %v",
		e.fieldName, e.rawValue, e.keyName, e.reason)
}
