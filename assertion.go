package assertion

import (
	"errors"
	"fmt"
	"reflect"
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

const (
	cmpOpEqual = iota
	cmpOpNotEqual
	cmpOpGreater
	cmpOpLowerEqual
	cmpOpLower
	cmpOpGreaterEqual
)

var errMsgByOp = map[int]string{
	cmpOpEqual:        errMsgNotEqual,
	cmpOpNotEqual:     errMsgNotDifferent,
	cmpOpGreater:      errMsgNotGreater,
	cmpOpLowerEqual:   errMsgNotLowerEqual,
	cmpOpLower:        errMsgNotLower,
	cmpOpGreaterEqual: errMsgNotGreaterEqual,
}

// compare returns true if a given value and other operand satisfy the compare
// operation determined by the operator. If the operation is not satisfied, this
// function also returns an error (only comparable types allowed)
func compare(op int, value, other interface{}, msgArgs ...interface{}) (bool, error) {
	switch op {
	case cmpOpNotEqual, cmpOpGreaterEqual, cmpOpLowerEqual:
		ok, err := compare(op-1, value, other, msgArgs...)
		if ok {
			err = buildError(fmt.Sprintf(errMsgByOp[op], value, other), msgArgs...)
		}
		return !ok, err
	}

	switch value.(type) {
	case bool:
		v, o := value.(bool), other.(bool)
		if op == cmpOpEqual && v == o {
			return true, nil
		}
	case int, int8, int16, int32, int64:
		v, o := reflect.ValueOf(value).Int(), reflect.ValueOf(other).Int()
		if (op == cmpOpEqual && v == o) || (op == cmpOpGreater && v > o) || (op == cmpOpLower && v < o) {
			return true, nil
		}
	case uint, uint8, uint16, uint32, uint64:
		v, o := reflect.ValueOf(value).Uint(), reflect.ValueOf(other).Uint()
		if (op == cmpOpEqual && v == o) || (op == cmpOpGreater && v > o) || (op == cmpOpLower && v < o) {
			return true, nil
		}
	case float32, float64:
		v, o := reflect.ValueOf(value).Float(), reflect.ValueOf(other).Float()
		if (op == cmpOpEqual && v == o) || (op == cmpOpGreater && v > o) || (op == cmpOpLower && v < o) {
			return true, nil
		}
	case string:
		v, o := value.(string), other.(string)
		if (op == cmpOpEqual && v == o) || (op == cmpOpGreater && v > o) || (op == cmpOpLower && v < o) {
			return true, nil
		}
	}

	return false, buildError(fmt.Sprintf(errMsgByOp[op], value, other), msgArgs...)
}
