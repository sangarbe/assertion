package assertion

import (
	"fmt"
	"strconv"
)

// Boolean returns true if a given string is one of the following accepted forms:
// true, false, TRUE, FALSE, t, f, 1, or 0
func (a *Assertion) Boolean(value string, msgArgs ...interface{}) bool {
	if _, err := strconv.ParseBool(value); err == nil {
		return true
	}

	a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "boolean string"), msgArgs...)
	return false
}

// Truthy returns true if a given string is one of the following accepted forms:
// true, TRUE, t, or 1
func (a *Assertion) Truthy(value string, msgArgs ...interface{}) bool {
	if b, err := strconv.ParseBool(value); err == nil {
		return b
	}

	a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "truthy string"), msgArgs...)
	return false
}

// Falsy returns true if a given string is one of the following accepted forms:
// true, false, TRUE, FALSE, t, f, 1, and 0
func (a *Assertion) Falsy(value string, msgArgs ...interface{}) bool {
	if b, err := strconv.ParseBool(value); err == nil {
		return !b
	}

	a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "falsy string"), msgArgs...)
	return false
}

// Integer returns true if a given string can be parsed as a valid integer value
func (a *Assertion) Integer(value string, msgArgs ...interface{}) bool {
	if _, err := strconv.ParseInt(value, 0, 64); err == nil {
		return true
	}

	a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "integer"), msgArgs...)
	return false
}

// IntegerBinary returns true if a given string can be parsed as a valid integer value
// in base 2
func (a *Assertion) IntegerBinary(value string, msgArgs ...interface{}) bool {
	if _, err := strconv.ParseInt(value, 2, 64); err == nil {
		return true
	}

	a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "base-2 integer"), msgArgs...)
	return false
}

// IntegerOctal returns true if a given string can be parsed as a valid integer value
// in base 8
func (a *Assertion) IntegerOctal(value string, msgArgs ...interface{}) bool {
	if _, err := strconv.ParseInt(value, 8, 64); err == nil {
		return true
	}

	a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "base-8 integer"), msgArgs...)
	return false
}

// IntegerHexadecimal returns true if a given string can be parsed as a valid integer value
// in base 16
func (a *Assertion) IntegerHexadecimal(value string, msgArgs ...interface{}) bool {
	if _, err := strconv.ParseInt(value, 16, 64); err == nil {
		return true
	}

	a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "base-16 integer"), msgArgs...)
	return false
}

// IntegerDecimal returns true if a given string can be parsed as a valid integer value
// in base 10
func (a *Assertion) IntegerDecimal(value string, msgArgs ...interface{}) bool {
	if _, err := strconv.ParseInt(value, 10, 64); err == nil {
		return true
	}

	a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "base-10 integer"), msgArgs...)
	return false
}

// Unsigned returns true if a given string can be parsed as a valid unsigned integer value
func (a *Assertion) Unsigned(value string, msgArgs ...interface{}) bool {
	if _, err := strconv.ParseUint(value, 0, 64); err == nil {
		return true
	}

	a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "unsigned integer"), msgArgs...)
	return false
}

// UnsignedBinary returns true if a given string can be parsed as a valid unsigned integer value
// in base 2
func (a *Assertion) UnsignedBinary(value string, msgArgs ...interface{}) bool {
	if _, err := strconv.ParseUint(value, 2, 64); err == nil {
		return true
	}

	a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "base-2 unsigned integer"), msgArgs...)
	return false
}

// UnsignedOctal returns true if a given string can be parsed as a valid unsigned integer value
// in base 8
func (a *Assertion) UnsignedOctal(value string, msgArgs ...interface{}) bool {
	if _, err := strconv.ParseUint(value, 8, 64); err == nil {
		return true
	}

	a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "base-8 unsigned integer"), msgArgs...)
	return false
}

// UnsignedHexadecimal returns true if a given string can be parsed as a valid unsigned integer value
// in base 16
func (a *Assertion) UnsignedHexadecimal(value string, msgArgs ...interface{}) bool {
	if _, err := strconv.ParseUint(value, 16, 64); err == nil {
		return true
	}

	a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "base-16 unsigned integer"), msgArgs...)
	return false
}

// UnsignedDecimal returns true if a given string can be parsed as a valid unsigned integer value
// in base 10
func (a *Assertion) UnsignedDecimal(value string, msgArgs ...interface{}) bool {
	if _, err := strconv.ParseUint(value, 10, 64); err == nil {
		return true
	}

	a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "base-10 unsigned integer"), msgArgs...)
	return false
}

// Float returns true if a given string can be parsed as a valid float value
func (a *Assertion) Float(value string, msgArgs ...interface{}) bool {
	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return true
	}

	a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "float"), msgArgs...)
	return false
}


