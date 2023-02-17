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

// typeRep is used to add new pointer types
func typeRefRep(t any) reflect.Type {
	return reflect.TypeOf(t)
}

// typeRep is used to add new types from a pointer
func typeRep(t any) reflect.Type {
	return typeRefRep(t).Elem()
}

// typeConverter represents the function that will try to convert the raw
// value from a given variable and set it to the corresponding field represented
// by vRep. vRep must be a value of a settable struct field
type typeConverter = func(vRep *reflect.Value, rawVal string) error

// typeConverterInfo contains the pretended
// type set with typeRep and the associated
// type converter
type typeConverterInfo struct {
	typeRep   reflect.Type
	converter typeConverter
}

var (
	// Type Parsers
	parseStringType,
	parseIntegerType,
	parseInteger32Type,
	parseUnsignedInteger16Type,
	parseDurationType,
	parseBoolType,
	parseAddrsRefType,
	_ *typeConverterInfo

	// typeConverters contains each type converter
	typeConverters []*typeConverterInfo
)

func init() {
	parseStringType = &typeConverterInfo{
		typeRep: typeRep((*string)(nil)),
		// Sets the raw value itself
		converter: func(vRep *reflect.Value, rawVal string) error {
			vRep.SetString(rawVal)

			return nil
		},
	}

	parseIntegerType = &typeConverterInfo{
		typeRep: typeRep((*int)(nil)),
		// Tries to get an int from the given raw
		// value and set it to the corresponding field
		converter: func(vRep *reflect.Value, rawVal string) error {
			val, err := strconv.Atoi(rawVal)
			if err == nil {
				vRep.SetInt(int64(val))
			}

			return err
		},
	}

	parseInteger32Type = &typeConverterInfo{
		typeRep: typeRep((*int32)(nil)),
		// Tries to get an int32 from the given raw
		// value and set it to the corresponding field
		converter: func(vRep *reflect.Value, rawVal string) error {
			val, err := strconv.ParseInt(rawVal, 10, 32)
			if err == nil {
				vRep.SetInt(val)
			}

			return err
		},
	}

	parseUnsignedInteger16Type = &typeConverterInfo{
		typeRep: typeRep((*uint16)(nil)),
		// Tries to get an uint16 from the given raw
		// value and set it to the corresponding field
		converter: func(vRep *reflect.Value, rawVal string) error {
			val, err := strconv.ParseUint(rawVal, 10, 16)
			if err == nil {
				vRep.SetUint(val)
			}

			return err
		},
	}

	parseDurationType = &typeConverterInfo{
		typeRep: typeRep((*time.Duration)(nil)),
		// Tries to get a valid one from the given raw
		// value and set it to the corresponding field
		converter: func(vRep *reflect.Value, rawVal string) error {
			val, err := time.ParseDuration(rawVal)
			if err == nil {
				vRep.Set(reflect.ValueOf(val))
			}

			return err
		},
	}

	parseBoolType = &typeConverterInfo{
		typeRep: typeRep((*bool)(nil)),
		// Tries to get a bool from the given raw
		// value and set it to the corresponding field
		converter: func(vRep *reflect.Value, rawVal string) error {
			val, err := strconv.ParseBool(rawVal)
			if err == nil {
				vRep.SetBool(val)
			}

			return err
		},
	}

	parseAddrsRefType = &typeConverterInfo{
		typeRep: typeRefRep((*utils.Addrs)(nil)),
		// Tries to get a *utils.Addrs from the given raw
		// value and set it to the corresponding field
		converter: func(vRep *reflect.Value, rawVal string) error {
			val, err := utils.ParseAddrs(rawVal)
			if err == nil {
				vRep.Elem().Set(reflect.ValueOf(val))
			}

			return err
		},
	}

	// Groups all converters
	typeConverters = []*typeConverterInfo{
		parseStringType,
		parseIntegerType,
		parseInteger32Type,
		parseUnsignedInteger16Type,
		parseDurationType,
		parseBoolType,
		parseAddrsRefType,
	}
}
