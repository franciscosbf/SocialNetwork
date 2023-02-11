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

// Available tags
const (
	nameTagKey     = "name"
	requiredTagKey = "required"
	acceptsTagKey  = "accepts"
)

// parseTagKeyName fetches the variable name.
// Returns MissingTagKeyError if it is missing
func parseTagKeyName(field *reflect.StructField) (string, error) {
	name, ok := field.Tag.Lookup(nameTagKey)
	if !ok || name == "" {
		return "", &MissingTagKeyError{
			fieldName: field.Name,
			keyName:   nameTagKey,
		}
	}
	return name, nil
}

// parseTagKeyRequired returns true if set, otherwise false. If field is invalid
// or the value wasn't recognized returns false and InvalidTagKeyValueError
func parseTagKeyRequired(field *reflect.StructField) (bool, error) {
	if required, ok := field.Tag.Lookup(requiredTagKey); ok {
		switch strings.ToLower(required) {
		case "yes", "true":
			return true, nil
		case "no", "false": // false by default
		default:
			return false, &InvalidTagKeyValueError{
				fieldName:      field.Name,
				keyName:        requiredTagKey,
				acceptedValues: []string{"yes", "no", "true", "false"},
			}
		}
	}

	return false, nil
}

// parseTagKeyAccepts parses each accepted value and returns a set with them.
// If there isn't any value, returns an empty set. However, if the
// field is invalid returns InvalidTagKeyValueError. In the other hand,
// if the tag content is invalid (starting/ending with a comma, or
// some token is an empty string, then returns InvalidTagKeyValueFmtError
func parseTagKeyAccepts(field *reflect.StructField) (*utils.Set[string], error) {
	accepts, ok := field.Tag.Lookup(acceptsTagKey)
	if !ok {
		return utils.NewSet[string](), nil // Empty set means that accepts any value
	}

	if accepts == "" {
		return nil, &InvalidTagKeyValueError{
			fieldName:      field.Name,
			keyName:        acceptsTagKey,
			acceptedValues: []string{"any token separated by a comma, e.g. hi,hello,bye"},
		}
	}

	// Value can't start with a comma
	if strings.HasPrefix(accepts, ",") {
		return nil, &InvalidTagKeyValueFmtError{
			fieldName: field.Name,
			keyName:   acceptsTagKey,
			rawValue:  accepts,
			reason:    "value can't start with a comma",
		}
	}

	// Value can't end with a comma
	if strings.HasSuffix(accepts, ",") {
		return nil, &InvalidTagKeyValueFmtError{
			fieldName: field.Name,
			keyName:   acceptsTagKey,
			rawValue:  accepts,
			reason:    "value can't end with a comma",
		}
	}

	validTokens := utils.NewSet[string]()

	// Check each one and adds it to the accepted tokens set
	for _, token := range strings.Split(accepts, ",") {
		if token == "" {
			return nil, &InvalidTagKeyValueFmtError{
				fieldName: field.Name,
				keyName:   acceptsTagKey,
				rawValue:  accepts,
				reason:    "empty token",
			}
		}

		validTokens.Put(token)
	}

	return validTokens, nil
}
