package assertion

import (
	"fmt"
	"regexp"
	"strconv"
)

var emailRegex = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

// EqualBool returns true if a given bool value is equal to other bool value
func (a *Assertion) EqualBool(value, other bool, msgArgs ...interface{}) bool {
	if other == value {
		return true
	}

	a.appendError(fmt.Sprintf("%t is not %t", value, other), msgArgs...)
	return false
}

// True returns true if a given bool value is true
func (a *Assertion) True(value bool, msgArgs ...interface{}) bool {
	return a.EqualBool(value, true, msgArgs...)
}

// False returns true if a given bool value is false
func (a *Assertion) False(value bool, msgArgs ...interface{}) bool {
	return a.EqualBool(value, false, msgArgs...)
}

// Boolean returns true if a given string is one of the following accepted forms:
// true, false, TRUE, FALSE, t, f, 1, or 0
func (a *Assertion) Boolean(value string, msgArgs ...interface{}) bool {
	if _, err := strconv.ParseBool(value); err == nil{
		return true
	}

	a.appendError(fmt.Sprintf("%s is not a valid boolean string", value), msgArgs...)
	return false
}

// Truthy returns true if a given string is one of the following accepted forms:
// true, TRUE, t, or 1
func (a *Assertion) Truthy(value string, msgArgs ...interface{}) bool {
	if b, err := strconv.ParseBool(value); err == nil{
		return b
	}

	a.appendError(fmt.Sprintf("%s is not a valid truthy string", value), msgArgs...)
	return false
}

// Falsy returns true if a given string is one of the following accepted forms:
// true, false, TRUE, FALSE, t, f, 1, and 0
func (a *Assertion) Falsy(value string, msgArgs ...interface{}) bool {
	if b, err := strconv.ParseBool(value); err == nil{
		return !b
	}

	a.appendError(fmt.Sprintf("%s is not a valid falsy string", value), msgArgs...)
	return false
}

// EqualInt64 returns true if a given int64 value is equal to other int64 value
func (a *Assertion) EqualInt64(value, other int64, msgArgs ...interface{}) bool {
	if other == value {
		return true
	}

	a.appendError(fmt.Sprintf("%d is not equal %d", value, other), msgArgs...)
	return false
}

// EqualInt32 returns true if a given int32 value is equal to other int32 value
func (a *Assertion) EqualInt32(value, other int32, msgArgs ...interface{}) bool {
	return a.EqualInt64(int64(value), int64(other), msgArgs...)
}

// EqualInt16 returns true if a given int16 value is equal to other int16 value
func (a *Assertion) EqualInt16(value, other int16, msgArgs ...interface{}) bool {
	return a.EqualInt64(int64(value), int64(other), msgArgs...)
}

// EqualInt8 returns true if a given int8 value is equal to other int8 value
func (a *Assertion) EqualInt8(value, other int8, msgArgs ...interface{}) bool {
	return a.EqualInt64(int64(value), int64(other), msgArgs...)
}

// EqualInt returns true if a given int value is equal to other int value
func (a *Assertion) EqualInt(value, other int, msgArgs ...interface{}) bool {
	return a.EqualInt64(int64(value), int64(other), msgArgs...)
}

// EqualFloat64 returns true if a given float64 value is equal to other float64 value
func (a *Assertion) EqualFloat64(value, other float64, msgArgs ...interface{}) bool {
	if other == value {
		return true
	}

	a.appendError(fmt.Sprintf("%.2f is not equal %.2f", value, other), msgArgs...)
	return false
}

// EqualFloat32 returns true if a given float32 value is equal to other float32 value
func (a *Assertion) EqualFloat32(value, other float32, msgArgs ...interface{}) bool {
	return a.EqualFloat64(float64(value), float64(other), msgArgs...)
}

// GreaterThanInt returns true if a given int value is greater than a min value
func (a *Assertion) GreaterThanInt(value, min int64, msgArgs ...interface{}) bool {
	if value > min {
		return true
	}

	a.appendError(fmt.Sprintf("%d is not greater than %d", value, min), msgArgs...)
	return false
}

// GreaterThanOrEqualInt returns true if a given int value is greater than or equal to a min value
func (a *Assertion) GreaterThanOrEqualInt(value, min int64, msgArgs ...interface{}) bool {
	if value >= min {
		return true
	}

	a.appendError(fmt.Sprintf("%d is not greater than or equal %d", value, min), msgArgs...)
	return false
}

// LowerThanInt returns true if a given int value is lower than a max value
func (a *Assertion) LowerThanInt(value, max int64, msgArgs ...interface{}) bool {
	if value < max {
		return true
	}

	a.appendError(fmt.Sprintf("%d is not lower than %d", value, max), msgArgs...)
	return false
}

// LowerThanOrEqualInt returns true if a given int value is lower than or equal to a min value
func (a *Assertion) LowerThanOrEqualInt(value, max int64, msgArgs ...interface{}) bool {
	if value <= max {
		return true
	}

	a.appendError(fmt.Sprintf("%d is not lower than or equal %d", value, max), msgArgs...)
	return false
}

// Email returns true if a given value is a valid email format
func (a *Assertion) Email(value string, msgArgs ...interface{}) bool {
	if emailRegex.MatchString(value) {
		return true
	}

	a.appendError(fmt.Sprintf("email %s is invalid", value), msgArgs...)
	return false
}
