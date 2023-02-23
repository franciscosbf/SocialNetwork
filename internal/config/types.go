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
	"github.com/franciscosbf/micro-dwarf/internal/utils"
	"reflect"
	"strconv"
	"time"
)

// typeConverter represents the function that will try to convert the raw
// value from a given variable and set it to the corresponding field represented
// by vRep. vRep must be a value of a settable struct field
type typeConverter = func(vRep *reflect.Value, rawVal string) error

// parseString sets the raw value itself
func parseString(vRep *reflect.Value, rawVal string) error {
	vRep.SetString(rawVal)

	return nil
}

// parseInt tries to obtain an int from the given
// raw value and set it to the corresponding field
func parseInt(vRep *reflect.Value, rawVal string) error {
	val, err := strconv.Atoi(rawVal)
	if err == nil {
		vRep.SetInt(int64(val))
	}

	return err
}

// parseInt32 tries to obtain an int32 from the given
// raw value and set it to the corresponding field
func parseInt32(vRep *reflect.Value, rawVal string) error {
	val, err := strconv.ParseInt(rawVal, 10, 32)
	if err == nil {
		vRep.SetInt(val)
	}

	return err
}

// parseUnsignedInt16 tries to obtain an uint16 from the
// given raw value and set it to the corresponding field
func parseUnsignedInt16(vRep *reflect.Value, rawVal string) error {
	val, err := strconv.ParseUint(rawVal, 10, 16)
	if err == nil {
		vRep.SetUint(val)
	}

	return err
}

// parseDuration tries to obtain a valid one from the
// given raw value and set it to the corresponding field
func parseDuration(vRep *reflect.Value, rawVal string) error {
	val, err := time.ParseDuration(rawVal)
	if err == nil {
		vRep.Set(reflect.ValueOf(val))
	}

	return err
}

// parseBool tries to obtain a bool from the given
// raw value and set it to the corresponding field
func parseBool(vRep *reflect.Value, rawVal string) error {
	val, err := strconv.ParseBool(rawVal)
	if err == nil {
		vRep.SetBool(val)
	}

	return err
}

// parseAddrsRef tries to obtain a *utils.Addrs from the
// given raw value and set it to the corresponding field
func parseAddrsRef(vRep *reflect.Value, rawVal string) error {
	val, err := utils.ParseAddrs(rawVal)
	if err == nil {
		vRep.Set(reflect.ValueOf(val))
	}

	return err
}

// selectConverter returns a converter if field type is supported
func selectConverter(field *reflect.Value) typeConverter {
	inter := field.Interface()

	switch inter.(type) {
	case string:
		return parseString
	case int:
		return parseInt
	case int32:
		return parseInt32
	case uint16:
		return parseUnsignedInt16
	case time.Duration:
		return parseDuration
	case bool:
		return parseBool
	case *utils.Addrs:
		return parseAddrsRef
	default:
		return nil
	}
}
