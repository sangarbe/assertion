package assertion

import (
	"errors"
	"fmt"
)

const (
	errMsgMissingArgs       = `missing required arguments`
	errMsgNotSameType       = `%v and %v are not of the same type`
	errMsgNot               = `%v is not %v`
	errMsgNotEqual          = `%v is not equal %v`
	errMsgNotValid          = `%v is not a valid %v`
	errMsgNotGreater        = `%v is not greater than %v`
	errMsgNotLower          = `%v is not lower than %v`
	errMsgNotGreaterEqual   = `%v is not greater than or equal %v`
	errMsgNotLowerEqual     = `%v is not lower than or equal %v`
	errMsgNotDifferent      = `%v is not different %v`
	errMsgNotBetween        = `%v is not between %v and %v`
	errMsgNotBetweenExclude = `%v is not between %v and %v both excluded`
	errMsgNotStartsWith     = `%v does not start with %v`
	errMsgNotEndsWith       = `%v does not end with %v`
)

// Assertion represents a data assertion process. It provides several methods
// to execute common assertions, and stores and gives access to the corresponding
// errors result of assertion failures.
//
// Every assertion method admits message arguments (msgArgs) on their signature
// to allow the customization of error messages. If this arguments are provided
// they will form the error message in case of failure of the corresponding method.
type Assertion struct {
	errors []error
}

// New creates and returns a new Assertion
func New() Assertion {
	return Assertion{errors: make([]error, 0)}
}

// HasErrors returns if current Assertion stores some error
func (a *Assertion) HasErrors() bool {
	return len(a.errors) > 0
}

// CountErrors returns the number of current errors
func (a *Assertion) CountErrors() int {
	return len(a.errors)
}

// ErrorAt returns the error at given index. Negative indexes will be considered
// as reverse order, that is indexes from the last error element
func (a *Assertion) ErrorAt(index int) error {
	if index > len(a.errors)-1 {
		return nil
	}

	if len(a.errors)+index < 0 {
		return nil
	}

	if index >= 0 {
		return a.errors[index]
	}

	return a.errors[len(a.errors)+index]
}

// addError adds an error to Assertion
func (a *Assertion) addError(err error) {
	a.errors = append(a.errors, err)
}

// addErrorMsg adds the default error message to Assertion or a formatted error
// message if msgArgs are provided
func (a *Assertion) addErrorMsg(defaultMsg string, msgArgs ...interface{}) {
	a.errors = append(a.errors, buildError(defaultMsg, msgArgs...))
}

// buildError returns an error with a default message or a message built from
// message arguments if msgArgs are provided
func buildError(defaultMsg string, msgArgs ...interface{}) error {
	errMsg := defaultMsg
	if len(msgArgs) == 1 {
		errMsg = fmt.Sprintf("%+v", msgArgs[0])
	}

	if len(msgArgs) > 1 {
		errMsg = fmt.Sprintf(msgArgs[0].(string), msgArgs[1:]...)
	}

	return errors.New(errMsg)
}
