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
	"strings"
)

// buildAcceptedTks splits by a comma all
// tokens and puts them into the tokens set
func buildAcceptedTks(raw string, s *utils.Set[string]) {
	tokens := strings.Split(raw, ",")

	for _, token := range tokens {
		s.Put(token)
	}
}

// Available tags
const (
	nameTag     = "name"
	requiredTag = "required"
	acceptsTag  = "accepts"
)

// parseTagName fetches the variable name.
// Returns MissingTagError if it is missing
func parseTagName(field *reflect.StructField) (string, error) {
	name, ok := field.Tag.Lookup(nameTag)
	if !ok || name == "" {
		return "", &MissingTagError{
			fieldName: field.Name,
			tagName:   nameTag,
		}
	}
	return name, nil
}

// parseTagRequired returns true if set, otherwise false. If field is invalid
// or the value wasn't recognized returns false and InvalidTagValueError
func parseTagRequired(field *reflect.StructField) (bool, error) {
	if required, ok := field.Tag.Lookup(requiredTag); ok {
		switch required {
		case "yes", "true":
			return true, nil
		case "no", "false": // false by default
		default:
			return false, &InvalidTagValueError{
				fieldName:      field.Name,
				tagName:        requiredTag,
				acceptedValues: []string{"yes", "no", "true", "false"},
			}
		}
	}

	return false, nil
}

// parseTagAccepts parses each accepted value and returns a set with them.
// If there isn't any value, returns an empty set. However, if the
// field is invalid returns InvalidTagValueError
func parseTagAccepts(field *reflect.StructField) (*utils.Set[string], error) {
	tokensS := utils.NewSet[string]() // Empty set means that accepts any value

	if accepts, ok := field.Tag.Lookup(acceptsTag); ok {
		if accepts == "" {
			return nil, &InvalidTagValueError{
				fieldName:      field.Name,
				tagName:        acceptsTag,
				acceptedValues: []string{"any values separated by a comma, e.g. hi,hello,bye"},
			}
		}

		buildAcceptedTks(accepts, tokensS)
	}

	return tokensS, nil
}
