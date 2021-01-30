package assertion

import (
	"errors"
	"fmt"
)

// Assertion represents a data assertion process. It provides several methods
// to execute common assertions, and stores and gives access to the corresponding
// errors result of assertion failures
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

// IntGreaterThan returns true if a given integer value is greater than a min value
func (a *Assertion) IntGreaterThan(value, min int64, msgArgs ...interface{}) bool {
	if value > min {
		return true
	}

	a.appendError(fmt.Sprintf("%d is not greater than %d", value, min))
	return false
}

func (a *Assertion) appendError(defaultMsg string, msgArgs ...interface{}) {
	errMsg := defaultMsg
	if len(msgArgs) == 1 {
		errMsg = fmt.Sprintf("%+v", msgArgs[0])
	}

	if len(msgArgs) > 1 {
		errMsg = fmt.Sprintf(msgArgs[0].(string), msgArgs[1:]...)
	}

	a.errors = append(a.errors, errors.New(errMsg))
}