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

// StructPtr represents any struct pointer
type StructPtr = any

// Error codes
const (
	ErrorCodeInvalidGetVar errorw.ErrorCode = iota
	ErrorCodeMissingVar
	ErrorCodeUnacceptedVal
	ErrorCodeInvalidVarType
)

// ConfParser represents a client config that
// fetches variables from a given reader
type ConfParser struct {
	reader *envvars.VarReader
}

// variableInfo contains all parsed info from a
// struct field. It's used to evaluate a variable
type variableInfo struct {
	varName        string
	required       bool
	acceptedValues *utils.Set[string]
	val            *reflect.Value
	setValue       typeConverter
}

// isValidKeyword checks if a given value matches one of the accepted
// keywords. If it doesn't have any keyword, then returns true.
func (v *variableInfo) isValidKeyword(val string) bool {
	return v.acceptedValues.Empty() || v.acceptedValues.Contains(val)
}

// validKeywords returns a slice of accepted keywords
func (v *variableInfo) validKeywords() []string {
	return v.acceptedValues.Values()
}

// extractStrVal returns a ready to evaluate struct. If it doesn't respect the
// expected type returns the error InvalidPointerError or InvalidValuePointedError
func extractStrVal(possibleSrt StructPtr) (*reflect.Value, error) {
	// Check if is a pointer
	value := reflect.ValueOf(possibleSrt)
	if value.Kind() != reflect.Pointer {
		return nil, InvalidPointerError
	}

	// Check if points to a struct
	extractedValue := value.Elem()
	if extractedValue.Kind() != reflect.Struct {
		return nil, InvalidValuePointedError
	}

	return &extractedValue, nil
}

// isAssignable tells if a given
// struct field is exported or not
func isAssignable(field *reflect.Value) bool {
	return field.CanSet()
}

// selectTypeConverter searches in the types converter repository if someone
// matches the field type. If not, then means that is an unsupported type,
// returning an error
func selectTypeConverter(v *variableInfo, field *reflect.StructField) error {
	// Searches for the corresponding converter
	for _, pair := range typeConverters {
		if field.Type.AssignableTo(pair.typeRep) {
			v.setValue = pair.converter

			return nil
		}
	}

	return &UnsupportedTypeError{
		fieldName: field.Name,
		typeName:  field.Type.Name(),
	}
}

// parseFieldTagKeys tries to parse each tag key.
// Returns an error on the first invalid tag element.
// The error that may come directly from this func
// is RepeatedVarNameError, in case of a variable's
// name was already assigned to another struct field
func parseFieldTagKeys(
	v *variableInfo,
	field *reflect.StructField,
	parsedNames map[string]string,
) (err error) {
	if v.varName, err = parseTagKeyName(field); err != nil {
		return
	}

	if assignedField, ok := parsedNames[v.varName]; ok {
		return &RepeatedVarNameError{
			assignedFieldName: assignedField,
			varName:           v.varName,
		}
	}
	// Otherwise, caches it
	parsedNames[v.varName] = field.Name

	if v.required, err = parseTagKeyRequired(field); err != nil {
		return
	}

	if v.acceptedValues, err = parseTagKeyAccepts(field); err != nil {
		return
	}

	return
}

// validateKeywords checks if each one matches
// the field type by converting it with the
// required parser. Returns TypeInconsistencyError
// if some value has a different type
func validateKeywords(
	fieldT reflect.Type,
	parser typeConverter,
	accepts *utils.Set[string],
) error {
	// Used to try to assign each value
	dummyField := reflect.New(fieldT).Elem()

	for _, v := range accepts.Values() {
		if err := parser(&dummyField, v); err != nil {
			return &TypeInconsistencyError{
				fieldName: fieldT.Name(),
				typeName:  fieldT.Name(),
				rawValue:  v,
			}
		}
	}

	return nil
}

// parseFields iterates over each struct field, evaluating its type
// and tag elements. Returns a slice containing info of all struct
// variables. Upon some error while evaluating a field, it's returned
// immediately after have received it. Errors from this function (not
// the ones that might be returned from other calls) are WithoutFieldsError
// and PrivateFieldError.
func parseFields(strInfo *reflect.Value) ([]*variableInfo, error) {
	sType := strInfo.Type()
	fieldsNum := strInfo.NumField()

	if fieldsNum == 0 {
		return nil, WithoutFieldsError
	}

	// Cache to look for repeated variable names
	parsedNames := make(map[string]string)

	var fields []*variableInfo

	// Extracts info from each struct field
	for i := 0; i < fieldsNum; i++ {
		fieldV := strInfo.Field(i)
		fieldSt := sType.Field(i)

		if !isAssignable(&fieldV) {
			return nil, &PrivateFieldError{fieldName: fieldSt.Name}
		}

		newVar := &variableInfo{}

		// Set elements according to Type representation
		if err := selectTypeConverter(newVar, &fieldSt); err != nil {
			return nil, err
		}
		if err := parseFieldTagKeys(newVar, &fieldSt, parsedNames); err != nil {
			return nil, err
		}

		fieldT := fieldSt.Type
		converter := newVar.setValue
		keywords := newVar.acceptedValues
		if err := validateKeywords(fieldT, converter, keywords); err != nil {
			return nil, err
		}

		// Set value representation
		newVar.val = &fieldV

		fields = append(fields, newVar)
	}

	return fields, nil
}

// fillFields iterates over each field and evaluates the read
// value from the config reader, according to the parsed info.
// Lastly, tries to parse the raw value and set it into the field
func (cp *ConfParser) fillFields(vars []*variableInfo) error {
	for _, v := range vars {
		vName := v.varName

		rawVal, err := cp.reader.Get(vName)
		if err != nil {
			return errorw.WrapErrorf(
				ErrorCodeInvalidGetVar, err,
				"Error while trying to get value from variable %v", vName)
		}

		if rawVal == "" {
			if v.required {
				return errorw.WrapErrorf(
					ErrorCodeMissingVar, nil, "Missing variable %v", vName)
			}

			continue // struct field value isn't changed
		}

		if !v.isValidKeyword(rawVal) {
			return errorw.WrapErrorf(
				ErrorCodeUnacceptedVal, nil,
				"Unaccepted value \"%v\" of variable %v. Valid keywords: %v",
				rawVal, vName, strings.Join(v.validKeywords(), ", "))
		}

		if err := v.setValue(v.val, rawVal); err != nil {
			return errorw.WrapErrorf(
				ErrorCodeInvalidVarType, err,
				"Invalid value type of variable %v", vName)
		}
	}

	return nil
}

// ParseConf reads the struct contents (field type, tag and its proper name)
// and maps each field with the associated variable name specified in its tag,
// among other configuration elements. Then, extracts and converts the contents
// fetched by the variables reader provided to New. Finally, assigns
// the resulting value to the corresponding struct field. Errors related to
// invalid struct pointer or invalid structured fields are returned immediately.
// In the other hand, errors related to the content returned by the variables
// reader are wrapped by errorw.Wrapper. Tag element value is trimmed,
// e.g. name:" VARIABLE_1  "  results in "VARIABLE_1"
//
//	Restrictions:
//
//	1. Valid types and their conversions:
//
//			- string <- directly parsed
//			- int <- strconv.Atoi(raw)
//			- int32 <- strconv.ParseInt(raw, 10, 32)
//			- uint16 <- strconv.ParseUint(raw, 10, 16)
//			- time.Duration <- time.ParseDuration(raw)
//			- bool <- strconv.ParseBool(raw)
//			- *utils.Addrs <- utils.ParseAddrs(raw)
//
//	2. Valid tag elements:
//
//	    	- name: variable's name (required) in the
//	     	format `[a-zA-Z]\w+`
//
//	    	- required: valid keywords are true,false,
//	    	yes and no (field is optional by default)
//
//	    	- accepts: accepted keywords, separated by
//	    	a comma, that can be passed to a given variable
//	    	(if omitted, it means that are all accepted)
//
//	Struct example:
//
//	type S struct {
//		A int `name:"VAR_1" accepts:"1,2,  3   "`
//		B string `name:"    VAR_2" required:"yes   ", accepts:"bark,meow"`
//		C time.Duration `name:"VAR_3" required:"   false"`
//	}
//
//	// Defined variables: VAR_1=2, VAR_2=bark, VAR_3=1h30m
//
//	var varReader envvars.VarReader
//	// varReader setup ...
//
//	s := &S{}
//	parser, _ := config.New(varReader)
//	_ := parser.ParseConf(s)
//	fmt.Println(s)
func (cp *ConfParser) ParseConf(from StructPtr) error {
	srtVal, err := extractStrVal(from)
	if err != nil {
		return err
	}

	variables, err := parseFields(srtVal)
	if err != nil {
		return err
	}

	return cp.fillFields(variables)
}

// New returns a new config parser with
// a variables reader associated to it. Returns
// MissingVariablesReader if varReader is nil
func New(varReader *envvars.VarReader) (*ConfParser, error) {
	if varReader == nil {
		return nil, MissingVariablesReader
	}

	return &ConfParser{
		reader: varReader,
	}, nil
}
