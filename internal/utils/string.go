package utils

import "strings"

// PolishString trims a string,
// removing unnecessary blank spaces
func PolishString(s string) string {
	return strings.Trim(s, " \t\n\v\f\r")
}
