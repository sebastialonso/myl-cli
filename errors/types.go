package errors

import (
	"fmt"
	"myl/utils"
)

type baseError struct {
	code string
	message string
	errs []error
}

func newBaseError(code, message string, origErrs []error) *baseError {
	e := &baseError{
		code: code,
		message: message,
		errs: origErrs,
	}
	return e
}

// baseError implements the error interface
func (e baseError) Error() string {
	if len(e.errs) > 0 {
		return utils.SprintError(e.code, e.message, "", e.errs[0])
	}
	return utils.SprintError(e.code, e.message, "", nil)
}

// baseError implements the Stringer interface
func (e baseError) String() string {
	return e.Error()
}

// with the following 3 functions, baseError implements the MylError interface
// defined in errors.go
func (e baseError) Code() string {
	return e.code
}

func (e baseError) Message() string {
	return e.message
}

func (e baseError) OrigErrs() []error {
	return e.errs
}

type mylerror MylError

type commandError struct {
	mylerror
	token string
}

func newCommandError(err mylerror, token string) *commandError {
	return &commandError{
		mylerror: err,
		token: token,

	}
}

// commandError implements the error interface
func (e commandError) Error() string {
	return utils.SprintError(e.Code(), e.Message(), "", nil)
}

// commandError implements Stringer interface
func (e commandError) String() string {
	return e.Error()
}

func (e commandError) Message() string {
	return e.mylerror.Message() + fmt.Sprintf(" %s", e.token)
}

func (e commandError) Token() string {
	return e.token
}