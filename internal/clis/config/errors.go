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
	"reflect"
	"strings"
)

var MissingPublicFieldsError = errors.New("struct doesn't have any public field")
var InvalidPointerError = errors.New("expecting pointer to struct")
var InvalidValuePointedError = errors.New("expecting pointer referencing a non-nil struct")

type MissingTagError struct {
	fieldName string
	tagName   string
}

func (e *MissingTagError) Error() string {
	return fmt.Sprintf(
		"missing tag %v in struct field %v",
		e.tagName, e.fieldName)
}

type InvalidTagValueError struct {
	fieldName      string
	tagName        string
	acceptedValues []string
}

func (e *InvalidTagValueError) Error() string {
	return fmt.Sprintf(
		"invalid value in tag %v of struct field %v; accepted values %v",
		strings.Join(e.acceptedValues, ","), e.tagName, e.fieldName)
}

type UnsupportedTypeError struct {
	fieldName string
	fieldType reflect.Type
}

func (e *UnsupportedTypeError) Error() string {
	return fmt.Sprintf(
		"struct field %v contains unsupported type %v",
		e.fieldName, e.fieldType.Name())
}
