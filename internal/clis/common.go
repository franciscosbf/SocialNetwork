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

package clis

import (
	"github.com/franciscosbf/micro-dwarf/internal/errorw"
)

// Error codes
const (
	ErrorCodeMissingConfig errorw.ErrorCode = iota
	ErrorCodeInvalidGetVar
	ErrorCodeMissingVar
	ErrorCodeUnacceptedVal
	ErrorCodeInvalidVarType
)

// VarInfo represents the base getters
// that each variable info must have
type VarInfo interface {
	Name() string
	Required() bool
}
