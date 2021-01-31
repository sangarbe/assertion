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

// EqualBool returns true if a given bool value is equal to other bool value
func (a *Assertion) EqualBool(value, other bool, msgArgs ...interface{}) bool {
	if other == value {
		return true
	}

	a.appendError(fmt.Sprintf("%t is not %t", value, other))
	return false
}

// True returns true if a given bool value is true
func (a *Assertion) True(value bool, msgArgs ...interface{}) bool {
	return a.EqualBool(value, true)
}

// False returns true if a given bool value is false
func (a *Assertion) False(value bool, msgArgs ...interface{}) bool {
	return a.EqualBool(value, false)
}

// EqualInt64 returns true if a given int64 value is equal to other int64 value
func (a *Assertion) EqualInt64(value, other int64, msgArgs ...interface{}) bool {
	if other == value {
		return true
	}

	a.appendError(fmt.Sprintf("%d is not equal %d", value, other))
	return false
}

// EqualInt32 returns true if a given int32 value is equal to other int32 value
func (a *Assertion) EqualInt32(value, other int32, msgArgs ...interface{}) bool {
	return a.EqualInt64(int64(value), int64(other))
}

// EqualInt16 returns true if a given int16 value is equal to other int16 value
func (a *Assertion) EqualInt16(value, other int16, msgArgs ...interface{}) bool {
	return a.EqualInt64(int64(value), int64(other))
}

// EqualInt8 returns true if a given int8 value is equal to other int8 value
func (a *Assertion) EqualInt8(value, other int8, msgArgs ...interface{}) bool {
	return a.EqualInt64(int64(value), int64(other))
}

// EqualInt returns true if a given int value is equal to other int value
func (a *Assertion) EqualInt(value, other int, msgArgs ...interface{}) bool {
	return a.EqualInt64(int64(value), int64(other))
}

// EqualFloat64 returns true if a given float64 value is equal to other float64 value
func (a *Assertion) EqualFloat64(value, other float64, msgArgs ...interface{}) bool {
	if other == value {
		return true
	}

	a.appendError(fmt.Sprintf("%f is not equal %f", value, other))
	return false
}

// EqualFloat32 returns true if a given float32 value is equal to other float32 value
func (a *Assertion) EqualFloat32(value, other float32, msgArgs ...interface{}) bool {
	if other == value {
		return true
	}

	a.appendError(fmt.Sprintf("%f is not equal %f", value, other))
	return false
}

// GreaterThanInt returns true if a given int value is greater than a min value
func (a *Assertion) GreaterThanInt(value, min int64, msgArgs ...interface{}) bool {
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
