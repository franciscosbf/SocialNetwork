package config

import (
	"github.com/franciscosbf/micro-dwarf/internal/utils"
	"reflect"
	"strings"
)

// buildAcceptedTks splits accepted values into
// tokens and puts them into the tokens set
func buildAcceptedTks(raw string, tokensSet *utils.Set[string]) {
	tokens := strings.Split(raw, ",")

	for _, token := range tokens {
		tokensSet.Put(token)
	}
}

// Available tags
const (
	nameTag     = "name"
	requiredTag = "required"
	acceptsTag  = "accepts"
)

// tagParser represents the reader that will evaluate
// and get the content from a given struct field and set
// the result to the corresponding
type tagParser = func(v *variableInfo, field *reflect.StructField) error

// tagParsers contains each tag parser
var tagParsers = []tagParser{
	// Parsed name tag
	func(v *variableInfo, field *reflect.StructField) error {
		name, ok := field.Tag.Lookup(nameTag)
		if !ok || name == "" {
			return &MissingTagError{
				fieldName: field.Name,
				tagName:   nameTag,
			}
		}
		v.varName = name

		return nil
	},
	// Parses required tag
	func(v *variableInfo, field *reflect.StructField) error {
		if required, ok := field.Tag.Lookup(requiredTag); ok {
			switch required {
			case "yes", "true":
				v.required = true
			case "no", "false": // false by default
			default:
				return &InvalidTagValueError{
					fieldName:      field.Name,
					tagName:        requiredTag,
					acceptedValues: []string{"yes", "no", "true", "false"},
				}
			}
		}

		return nil
	},
	// Parses accepts tag
	func(v *variableInfo, field *reflect.StructField) error {
		v.accepted = utils.NewSet[string]() // Empty set means that accepts any value
		if accepts, ok := field.Tag.Lookup(acceptsTag); ok {
			if accepts == "" {
				return &InvalidTagValueError{
					fieldName:      field.Name,
					tagName:        acceptsTag,
					acceptedValues: []string{"any value separated by a comma, e.g. hi,hello,bye"},
				}
			}

			buildAcceptedTks(accepts, v.accepted)
		}

		return nil
	},
}
