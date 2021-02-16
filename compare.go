package assertion

import (
	"fmt"
	"reflect"
)

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
	rv, ro := reflect.ValueOf(value), reflect.ValueOf(other)
	if rv.Kind() != ro.Kind() {
		return false, buildError(fmt.Sprintf(errMsgNotSameType, value, other), msgArgs...)
	}

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
		v, o := rv.Bool(), ro.Bool()
		if op == cmpOpEqual && v == o {
			return true, nil
		}
	case int, int8, int16, int32, int64:
		v, o := rv.Int(), ro.Int()
		if (op == cmpOpEqual && v == o) || (op == cmpOpGreater && v > o) || (op == cmpOpLower && v < o) {
			return true, nil
		}
	case uint, uint8, uint16, uint32, uint64:
		v, o := rv.Uint(), ro.Uint()
		if (op == cmpOpEqual && v == o) || (op == cmpOpGreater && v > o) || (op == cmpOpLower && v < o) {
			return true, nil
		}
	case float32, float64:
		v, o := rv.Float(), ro.Float()
		if (op == cmpOpEqual && v == o) || (op == cmpOpGreater && v > o) || (op == cmpOpLower && v < o) {
			return true, nil
		}
	case string:
		v, o := rv.String(), ro.String()
		if (op == cmpOpEqual && v == o) || (op == cmpOpGreater && v > o) || (op == cmpOpLower && v < o) {
			return true, nil
		}
	}

	return false, buildError(fmt.Sprintf(errMsgByOp[op], value, other), msgArgs...)
}

// validateArgsLength panics if args length is lower than minLength
func validateArgsLength(minLength int, args ...interface{}) {
	if len(args) < minLength {
		panic(buildError(errMsgMissingArgs))
	}
}

// Nil returns true if a given bool value is equal to other bool value
func (a *Assertion) Nil(args ...interface{}) bool {
	validateArgsLength(1, args...)

	if args[0] == nil {
		return true
	}

	v := reflect.ValueOf(args[0])
	switch v.Kind() {
	case reflect.Chan, reflect.Func,
		reflect.Interface, reflect.Map,
		reflect.Ptr, reflect.Slice:
		if v.IsNil() {
			return true
		}
	}

	a.addError(buildError(fmt.Sprintf(errMsgNot, args[0], nil), args[1:]...))
	return false
}

// Equal returns true if a given value is equal to other value
func (a *Assertion) Equal(args ...interface{}) bool {
	validateArgsLength(2, args...)

	ok, err := compare(cmpOpEqual, args[0], args[1], args[2:]...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// True returns true if a given bool value is true
func (a *Assertion) True(value bool, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpEqual, value, true, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// False returns true if a given bool value is false
func (a *Assertion) False(value bool, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpEqual, value, false, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThan returns true if a given int64 value is greater than other int64 value
func (a *Assertion) GreaterThan(args ...interface{}) bool {
	validateArgsLength(2, args...)

	ok, err := compare(cmpOpGreater, args[0], args[1], args[2:]...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanOrEqual returns true if a given value is greater than or equal to other value
func (a *Assertion) GreaterThanOrEqual(args ...interface{}) bool {
	validateArgsLength(2, args...)

	ok, err := compare(cmpOpGreaterEqual, args[0], args[1], args[2:]...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThan returns true if a given value is lower than other value
func (a *Assertion) LowerThan(args ...interface{}) bool {
	validateArgsLength(2, args...)

	ok, err := compare(cmpOpLower, args[0], args[1], args[2:]...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanOrEqual returns true if a given value is lower than or equal to other value
func (a *Assertion) LowerThanOrEqual(args ...interface{}) bool {
	validateArgsLength(2, args...)

	ok, err := compare(cmpOpLowerEqual, args[0], args[1], args[2:]...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// Between returns true if a given value is between a lower and upper
// limit values (including both)
func (a *Assertion) Between(args ...interface{}) bool {
	validateArgsLength(3, args...)

	for op, v := range map[int]interface{}{cmpOpGreaterEqual: args[1], cmpOpLowerEqual: args[2]} {
		ok, _ := compare(op, args[0], v, args[3:]...)
		if !ok {
			a.addError(buildError(fmt.Sprintf(errMsgNotBetween, args[0], args[1], args[2]), args[3:]...))
			return false
		}
	}

	return true
}

// BetweenExclude returns true if a given value is between a lower and upper
// limit values (excluding both)
func (a *Assertion) BetweenExclude(args ...interface{}) bool {
	validateArgsLength(3, args...)

	for op, v := range map[int]interface{}{cmpOpGreater: args[1], cmpOpLower: args[2]} {
		ok, _ := compare(op, args[0], v, args[3:]...)
		if !ok {
			a.addError(buildError(fmt.Sprintf(errMsgNotBetweenExclude, args[0], args[1], args[2]), args[3:]...))
			return false
		}
	}

	return true
}
