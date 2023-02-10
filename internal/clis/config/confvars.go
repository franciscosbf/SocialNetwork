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
	"github.com/franciscosbf/micro-dwarf/internal/envvars"
	"github.com/franciscosbf/micro-dwarf/internal/errorw"
	"github.com/franciscosbf/micro-dwarf/internal/utils"
	"reflect"
	"strings"
)

type StructPtr = any

// Error codes
const (
	ErrorCodeInvalidConf errorw.ErrorCode = iota
	ErrorCodeInvalidField
	ErrorCodeInvalidGetVar
	ErrorCodeMissingVar
	ErrorCodeUnacceptedVal
	ErrorCodeInvalidVarType
)

type VarsConf struct {
	reader *envvars.Config
}

type structInfo struct {
	strType reflect.Type
	strVal  *reflect.Value
}

type variableInfo struct {
	varName  string
	required bool
	accepted *utils.Set[string]
	val      *reflect.Value
	rType    reflect.Type
	setValue typeConverter
}

// accepts checks if a given value is valid
func (v *variableInfo) accepts(val string) bool {

	return v.accepted.Contains(val)
}

// validValues returns a slice of accepted values
func (v *variableInfo) validValues() []string {
	return v.accepted.Values()
}

// TODO - "(...) from a reflect.Value it’s easy to get to the corresponding reflect.Type (...)"
//  Credits: https://go.dev/blog/laws-of-reflection

// extractStr returns a ready to evaluate struct. If it doesn't respect the
// expected type returns the error InvalidPointerError or InvalidValuePointedError
func extractStr(possibleStr StructPtr) (*structInfo, error) {
	// Check if is a pointer
	value := reflect.TypeOf(possibleStr)
	if value.Kind() != reflect.Pointer {
		return nil, InvalidPointerError
	}

	// Check if points to a struct
	extractedType := value.Elem()
	if extractedType.Kind() != reflect.Struct {
		return nil, InvalidValuePointedError
	}

	// Get struct representation
	extractedValue := reflect.
		ValueOf(possibleStr).
		Elem()

	return &structInfo{
		strType: extractedType,
		strVal:  &extractedValue,
	}, nil
}

// setTypeConverter searches in the types converter repository if someone
// matches the field type. If not, then means that is an unsupported type,
// returning an error
func setTypeConverter(v *variableInfo, field reflect.StructField) error {
	var selectedConverter typeConverter

	// Searches for
	for _, pair := range typeConverters {
		if field.Type.AssignableTo(pair.typeRep) {
			selectedConverter = pair.converter
			break
		}
	}

	if selectedConverter == nil {
		return &UnsupportedTypeError{
			fieldName: field.Name,
			typeName:  field.Type.Name(),
		}
	}

	v.setValue = selectedConverter

	return nil
}

// parseFieldTags tries to parse each field tag.
// Returns an error on the first invalid tag
func parseFieldTags(v *variableInfo, field *reflect.StructField) error {
	for _, parser := range tagParsers {
		if err := parser(v, field); err != nil {
			return err
		}
	}

	return nil
}

// parseFields iterates over each struct field, evaluating its type
// and tags, collecting its type and value representations from
// reflect pkg. Returns a slice containing info of all struct variables
func parseFields(info *structInfo) ([]*variableInfo, error) {
	sType := info.strType
	fieldsNum := info.strType.NumField()

	var fields []*variableInfo

	for i := 0; i < fieldsNum; i++ {
		newVar := &variableInfo{}

		// Set elements according to Type representation
		fieldT := sType.Field(i)
		if err := setTypeConverter(newVar, fieldT); err != nil {
			return nil, err
		}
		if err := parseFieldTags(newVar, &fieldT); err != nil {
			return nil, err
		}
		newVar.rType = fieldT.Type

		// Set value representation
		fieldV := info.strVal.Field(i)
		newVar.val = &fieldV

		fields = append(fields, newVar)
	}

	return fields, nil
}

// fillFields iterates over each field and evaluates the read
// value from the config reader, according to the parsed info.
// Lastly, tries to parse the raw value and set it into the field
func (vc *VarsConf) fillFields(vars []*variableInfo) error {
	for _, v := range vars {
		rawVal, err := vc.reader.Get(v.varName)
		if err != nil {
			return errorw.WrapErrorf(
				ErrorCodeInvalidGetVar, err, "Invalid Redis config var Fetch")
		}

		if rawVal == "" {
			if v.required {
				return errorw.WrapErrorf(
					ErrorCodeMissingVar, nil,
					"Missing Redis required variableInfo %v", v.varName)
			}
			return nil
		}

		if v.accepts(rawVal) {
			return errorw.WrapErrorf(
				ErrorCodeUnacceptedVal, nil,
				"Unaccepted Redis value %v of variableInfo %v. Only accepts: %v",
				rawVal, v.varName, strings.Join(v.validValues(), ", "))
		}

		if err := v.setValue(v.val, rawVal); err != nil {
			return errorw.WrapErrorf(
				ErrorCodeInvalidVarType, err,
				"Invalid Redis type of variableInfo %v", v.varName)
		}
	}

	return nil
}

// ParseConf TODO - comment this - don't forget to specify the valid representation of duration
func (vc *VarsConf) ParseConf(from StructPtr) error {
	str, err := extractStr(from)
	if err != nil {
		return errorw.WrapErrorf(
			ErrorCodeInvalidConf, err, "Invalid config provided")
	}

	variables, err := parseFields(str)
	if err != nil {
		return errorw.WrapErrorf(
			ErrorCodeInvalidField, err, "Invalid parsed val")
	}

	return vc.fillFields(variables)
}

// NewVarsConf TODO - comment this
func NewVarsConf(varReader *envvars.Config) *VarsConf {
	return &VarsConf{
		reader: varReader,
	}
}
