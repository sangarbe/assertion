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

// EqualBool returns true if a given bool value is equal to other bool value
func (a *Assertion) EqualBool(value, other bool, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpEqual, value, other, msgArgs...)
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

// EqualInt64 returns true if a given int64 value is equal to other int64 value
func (a *Assertion) EqualInt64(value, other int64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// EqualInt32 returns true if a given int32 value is equal to other int32 value
func (a *Assertion) EqualInt32(value, other int32, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// EqualInt16 returns true if a given int16 value is equal to other int16 value
func (a *Assertion) EqualInt16(value, other int16, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// EqualInt8 returns true if a given int8 value is equal to other int8 value
func (a *Assertion) EqualInt8(value, other int8, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// EqualInt returns true if a given int value is equal to other int value
func (a *Assertion) EqualInt(value, other int, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// EqualUint64 returns true if a given uint64 value is equal to other uint64 value
func (a *Assertion) EqualUint64(value, other uint64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// EqualUint32 returns true if a given uint32 value is equal to other uint32 value
func (a *Assertion) EqualUint32(value, other uint32, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// EqualUint16 returns true if a given uint16 value is equal to other uint16 value
func (a *Assertion) EqualUint16(value, other uint16, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// EqualUint8 returns true if a given uint8 value is equal to other uint8 value
func (a *Assertion) EqualUint8(value, other uint8, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// EqualUint returns true if a given uint value is equal to other uint value
func (a *Assertion) EqualUint(value, other uint, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// EqualFloat64 returns true if a given float64 value is equal to other float64 value
func (a *Assertion) EqualFloat64(value, other float64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// EqualFloat32 returns true if a given float32 value is equal to other float32 value
func (a *Assertion) EqualFloat32(value, other float32, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// EqualString returns true if a given string value is equal to other string value
func (a *Assertion) EqualString(value, other string, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanInt64 returns true if a given int64 value is greater than other int64 value
func (a *Assertion) GreaterThanInt64(value, other int64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreater, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanInt32 returns true if a given int32 value is greater than other int32 value
func (a *Assertion) GreaterThanInt32(value, other int32, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreater, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanInt16 returns true if a given int16 value is greater than other int16 value
func (a *Assertion) GreaterThanInt16(value, other int16, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreater, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanInt8 returns true if a given int8 value is greater than other int8 value
func (a *Assertion) GreaterThanInt8(value, other int8, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreater, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanInt returns true if a given int value is greater than other int value
func (a *Assertion) GreaterThanInt(value, other int64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreater, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanUint64 returns true if a given uint64 value is greater than other uint64 value
func (a *Assertion) GreaterThanUint64(value, other uint64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreater, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanUint32 returns true if a given uint32 value is greater than other uint32 value
func (a *Assertion) GreaterThanUint32(value, other uint32, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreater, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanUint16 returns true if a given uint16 value is greater than other uint16 value
func (a *Assertion) GreaterThanUint16(value, other uint16, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreater, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanUint8 returns true if a given uint8 value is greater than other uint8 value
func (a *Assertion) GreaterThanUint8(value, other uint8, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreater, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanUint returns true if a given uint value is greater than other uint value
func (a *Assertion) GreaterThanUint(value, other uint, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreater, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanFloat64 returns true if a given float64 value is greater than other float64 value
func (a *Assertion) GreaterThanFloat64(value, other float64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreater, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanFloat32 returns true if a given float32 value is greater than other float32 value
func (a *Assertion) GreaterThanFloat32(value, other float32, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreater, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanString returns true if a given string value is greater than other string value
func (a *Assertion) GreaterThanString(value, other string, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreater, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanOrEqualInt64 returns true if a given int64 value is greater than or equal to other int64 value
func (a *Assertion) GreaterThanOrEqualInt64(value, other int64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreaterEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanOrEqualInt32 returns true if a given int32 value is greater than or equal to other int32 value
func (a *Assertion) GreaterThanOrEqualInt32(value, other int32, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreaterEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanOrEqualInt16 returns true if a given int16 value is greater than or equal to other int16 value
func (a *Assertion) GreaterThanOrEqualInt16(value, other int16, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreaterEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanOrEqualInt8 returns true if a given int8 value is greater than or equal to other int8 value
func (a *Assertion) GreaterThanOrEqualInt8(value, other int8, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreaterEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanOrEqualInt returns true if a given int value is greater than or equal to other int value
func (a *Assertion) GreaterThanOrEqualInt(value, other int, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreaterEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanOrEqualUint64 returns true if a given uint64 value is greater than or equal to other uint64 value
func (a *Assertion) GreaterThanOrEqualUint64(value, other uint64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreaterEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanOrEqualUint32 returns true if a given uint32 value is greater than or equal to other uint32 value
func (a *Assertion) GreaterThanOrEqualUint32(value, other uint32, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreaterEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanOrEqualUint16 returns true if a given uint16 value is greater than or equal to other uint16 value
func (a *Assertion) GreaterThanOrEqualUint16(value, other uint16, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreaterEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanOrEqualUint8 returns true if a given uint8 value is greater than or equal to other uint8 value
func (a *Assertion) GreaterThanOrEqualUint8(value, other uint8, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreaterEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanOrEqualUint returns true if a given uint value is greater than or equal to other uint value
func (a *Assertion) GreaterThanOrEqualUint(value, other uint, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreaterEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanOrEqualFloat64 returns true if a given float64 value is greater than or equal to other float64 value
func (a *Assertion) GreaterThanOrEqualFloat64(value, other float64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreaterEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanOrEqualFloat32 returns true if a given float32 value is greater than or equal to other float32 value
func (a *Assertion) GreaterThanOrEqualFloat32(value, other float32, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreaterEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanOrEqualString returns true if a given string value is greater than or equal to other string value
func (a *Assertion) GreaterThanOrEqualString(value, other string, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreaterEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanInt64 returns true if a given int64 value is lower than other int64 value
func (a *Assertion) LowerThanInt64(value, other int64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLower, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanInt32 returns true if a given int32 value is lower than other int32 value
func (a *Assertion) LowerThanInt32(value, other int32, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLower, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanInt16 returns true if a given int16 value is lower than other int16 value
func (a *Assertion) LowerThanInt16(value, other int16, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLower, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanInt8 returns true if a given int8 value is lower than other int8 value
func (a *Assertion) LowerThanInt8(value, other int8, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLower, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanInt returns true if a given int value is lower than other int value
func (a *Assertion) LowerThanInt(value, other int, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLower, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanUint64 returns true if a given uint64 value is lower than other uint64 value
func (a *Assertion) LowerThanUint64(value, other uint64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLower, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanUint32 returns true if a given uint32 value is lower than other uint32 value
func (a *Assertion) LowerThanUint32(value, other uint32, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLower, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanUint16 returns true if a given uint16 value is lower than other uint16 value
func (a *Assertion) LowerThanUint16(value, other uint16, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLower, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanUint8 returns true if a given uint8 value is lower than other uint8 value
func (a *Assertion) LowerThanUint8(value, other uint8, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLower, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanUint returns true if a given uint value is lower than other uint value
func (a *Assertion) LowerThanUint(value, other uint, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLower, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanFloat64 returns true if a given float64 value is lower than other float64 value
func (a *Assertion) LowerThanFloat64(value, other float64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLower, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanFloat32 returns true if a given float32 value is lower than other float32 value
func (a *Assertion) LowerThanFloat32(value, other float32, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLower, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanString returns true if a given string value is lower than other string value
func (a *Assertion) LowerThanString(value, other string, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLower, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanOrEqualInt64 returns true if a given int64 value is lower than or equal to other int64 value
func (a *Assertion) LowerThanOrEqualInt64(value, other int64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLowerEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanOrEqualInt32 returns true if a given int32 value is lower than or equal to other int32 value
func (a *Assertion) LowerThanOrEqualInt32(value, other int32, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLowerEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanOrEqualInt16 returns true if a given int16 value is lower than or equal to other int16 value
func (a *Assertion) LowerThanOrEqualInt16(value, other int16, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLowerEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanOrEqualInt8 returns true if a given int8 value is lower than or equal to other int8 value
func (a *Assertion) LowerThanOrEqualInt8(value, other int8, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLowerEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanOrEqualInt returns true if a given int value is lower than or equal to other int value
func (a *Assertion) LowerThanOrEqualInt(value, other int, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLowerEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanOrEqualUint64 returns true if a given uint64 value is lower than or equal to other uint64 value
func (a *Assertion) LowerThanOrEqualUint64(value, other uint64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLowerEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanOrEqualUint32 returns true if a given uint32 value is lower than or equal to other uint32 value
func (a *Assertion) LowerThanOrEqualUint32(value, other uint32, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLowerEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanOrEqualUint16 returns true if a given uint16 value is lower than or equal to other uint16 value
func (a *Assertion) LowerThanOrEqualUint16(value, other uint16, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLowerEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanOrEqualUint8 returns true if a given uint8 value is lower than or equal to other uint8 value
func (a *Assertion) LowerThanOrEqualUint8(value, other uint8, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLowerEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanOrEqualUint returns true if a given uint value is lower than or equal to other uint value
func (a *Assertion) LowerThanOrEqualUint(value, other uint, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLowerEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanOrEqualFloat64 returns true if a given float64 value is lower than or equal to other float64 value
func (a *Assertion) LowerThanOrEqualFloat64(value, other float64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLowerEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanOrEqualFloat32 returns true if a given float32 value is lower than or equal to other float32 value
func (a *Assertion) LowerThanOrEqualFloat32(value, other float32, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLowerEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanOrEqualString returns true if a given string value is lower than or equal to other string value
func (a *Assertion) LowerThanOrEqualString(value, other string, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLowerEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}
