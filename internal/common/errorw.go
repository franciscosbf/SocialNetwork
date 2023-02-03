package common

import "fmt"

// ErrorWrap contains an error with a nice message along
// with a code which give some meaning about its nature
type ErrorWrap struct {
	code   ErrorCode
	origin error
	msg    string
}

// ErrorCode represents the error nature. Error codes
// are implementation specific, giving the freedom to
// shape them according to their environment
type ErrorCode uint

// Error returns the given message passed to WrapErrorf
// if origin is nil, otherwise returns a formatted string
// containing both
func (e *ErrorWrap) Error() string {
	if e.origin != nil {
		return fmt.Sprintf("%v: %v", e.msg, e.origin)
	}

	return e.msg
}

// Unwrap returns the origin error. It can be nil,
// depending on what was given to WrapErrorf
func (e *ErrorWrap) Unwrap() error {
	return e.origin
}

// Code returns the error code
func (e *ErrorWrap) Code() ErrorCode {
	return e.code
}

// WrapErrorf wraps around and returns an error with a given status
// code and a msg to be formatted with optional parameters. There isn't
// any verification about the nullability of each parameter
func WrapErrorf(code ErrorCode, origin error, format string, fmtArgs ...any) error {
	msg := fmt.Sprintf(format, fmtArgs)

	return &ErrorWrap{
		code:   code,
		origin: origin,
		msg:    msg,
	}
}
