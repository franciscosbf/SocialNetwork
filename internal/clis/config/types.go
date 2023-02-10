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
	"reflect"
	"strconv"
	"time"
)

// typeRep is used to add new types
func typeRep(t any) reflect.Type {
	return reflect.TypeOf(t).Elem()
}

// typeConverter represents the function that will try to convert the
// raw value from a given variable and set it to the corresponding
// field represented by vRep
type typeConverter = func(vRep *reflect.Value, rawVal string) error

// typeConverters contains each type converter
var typeConverters = []struct {
	typeRep   reflect.Type
	converter typeConverter
}{
	// String conversion only sets the given raw string
	{
		typeRep: typeRep((*string)(nil)),
		converter: func(vRep *reflect.Value, rawVal string) error {
			vRep.SetString(rawVal)

			return nil
		},
	},
	// Integer conversion tries to get an int from the given
	// raw value and set it to the corresponding field
	{
		typeRep: typeRep((*int)(nil)),
		converter: func(vRep *reflect.Value, rawVal string) error {
			val, err := strconv.Atoi(rawVal)
			if err == nil {
				vRep.SetInt(int64(val))
			}

			return err
		},
	},
	// Duration conversion tries to get a valid one from the
	// given raw value and set it to the corresponding field
	{
		typeRep: typeRep((*time.Duration)(nil)),
		converter: func(vRep *reflect.Value, rawVal string) error {
			val, err := time.ParseDuration(rawVal)
			if err == nil {
				vRep.Set(reflect.ValueOf(val))
			}

			return err
		},
	},
}