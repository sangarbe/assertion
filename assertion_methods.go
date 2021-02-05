package assertion

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	errMsgNotEqual = `%v is not equal %v`
	errMsgNotValid = `%v is not a valid %v`
	errMsgNotGreater = `%v is not greater than %v`
	errMsgNotLower = `%v is not lower than %v`
	errMsgNotGreaterEqual = `%v is not greater than or equal %v`
	errMsgNotLowerEqual = `%v is not lower than or equal %v`
)

var (
	rexText          = "[0-9A-Za-z!#-'*+\\-/=?^_`{-~]"
	rexDottedString  = fmt.Sprintf(`(?:%s)+(\.(?:%s)+)*`, rexText, rexText)
	rexQuotedText    = `[ !#-\[\]-~]`
	rexQuotedPair    = `\\[ -~]`
	rexQuotedContent = fmt.Sprintf(`(?:%s)|(?:%s)`, rexQuotedText, rexQuotedPair)
	rexOctet         = `0|25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9][0-9]?`
	rexQuotedString  = fmt.Sprintf(`"(?:%s)*"`, rexQuotedContent)
	rexLocalPart     = fmt.Sprintf(`(?:%s)|(?:%s)`, rexDottedString, rexQuotedString)
	rexIPv4Octets    = fmt.Sprintf(`(?:%s)(\.(?:%s)){3}`, rexOctet, rexOctet)
	rexSubdomain     = `[0-9A-Za-z]([\-0-9A-Za-z]{0,61}[0-9A-Za-z])?`
	rexDomain        = fmt.Sprintf(`(?:%s)(?:\.(?:%s))+`, rexSubdomain, rexSubdomain)
	rexIPv4OrDomain  = fmt.Sprintf(`(?:\[%s\])|(?:%s)`, rexIPv4Octets, rexDomain)
	rexEmail         = fmt.Sprintf(`^(?:%s)@(?:%s)$`, rexLocalPart, rexIPv4OrDomain)
	rexIPv4          = fmt.Sprintf(`^%s$`, rexIPv4Octets)
)

var (
	regexpEmail = regexp.MustCompile(rexEmail)
	regexpIpv4  = regexp.MustCompile(rexIPv4)
)

// EqualBool returns true if a given bool value is equal to other bool value
func (a *Assertion) EqualBool(value, other bool, msgArgs ...interface{}) bool {
	if other == value {
		return true
	}

	a.appendError(fmt.Sprintf(errMsgNotEqual, value, other), msgArgs...)
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
	if _, err := strconv.ParseBool(value); err == nil {
		return true
	}

	a.appendError(fmt.Sprintf(errMsgNotValid, value, "boolean string"), msgArgs...)
	return false
}

// Truthy returns true if a given string is one of the following accepted forms:
// true, TRUE, t, or 1
func (a *Assertion) Truthy(value string, msgArgs ...interface{}) bool {
	if b, err := strconv.ParseBool(value); err == nil {
		return b
	}

	a.appendError(fmt.Sprintf(errMsgNotValid, value, "truthy string"), msgArgs...)
	return false
}

// Falsy returns true if a given string is one of the following accepted forms:
// true, false, TRUE, FALSE, t, f, 1, and 0
func (a *Assertion) Falsy(value string, msgArgs ...interface{}) bool {
	if b, err := strconv.ParseBool(value); err == nil {
		return !b
	}

	a.appendError(fmt.Sprintf(errMsgNotValid, value, "falsy string"), msgArgs...)
	return false
}

// EqualInt64 returns true if a given int64 value is equal to other int64 value
func (a *Assertion) EqualInt64(value, other int64, msgArgs ...interface{}) bool {
	if other == value {
		return true
	}

	a.appendError(fmt.Sprintf(errMsgNotEqual, value, other), msgArgs...)
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

	a.appendError(fmt.Sprintf(errMsgNotEqual, value, other), msgArgs...)
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

	a.appendError(fmt.Sprintf(errMsgNotGreater, value, min), msgArgs...)
	return false
}

// GreaterThanOrEqualInt returns true if a given int value is greater than or equal to a min value
func (a *Assertion) GreaterThanOrEqualInt(value, min int64, msgArgs ...interface{}) bool {
	if value >= min {
		return true
	}

	a.appendError(fmt.Sprintf(errMsgNotGreaterEqual, value, min), msgArgs...)
	return false
}

// LowerThanInt returns true if a given int value is lower than a max value
func (a *Assertion) LowerThanInt(value, max int64, msgArgs ...interface{}) bool {
	if value < max {
		return true
	}

	a.appendError(fmt.Sprintf(errMsgNotLower, value, max), msgArgs...)
	return false
}

// LowerThanOrEqualInt returns true if a given int value is lower than or equal to a min value
func (a *Assertion) LowerThanOrEqualInt(value, max int64, msgArgs ...interface{}) bool {
	if value <= max {
		return true
	}

	a.appendError(fmt.Sprintf(errMsgNotLowerEqual, value, max), msgArgs...)
	return false
}

// Email returns true if a given value is a valid email format. It allows local
// portion to be quoted text and ipv4 for the domain portion (between square brackets).
func (a *Assertion) Email(value string, msgArgs ...interface{}) bool {
	if !regexpEmail.MatchString(value) {
		a.appendError(fmt.Sprintf(errMsgNotValid, value, "email"), msgArgs...)
		return false
	}

	splits := strings.Split(value, "@")
	domain := splits[len(splits)-1]
	if len(domain) > 255 {
		a.appendError(fmt.Sprintf(errMsgNotValid, value, "email"), msgArgs...)
		return false
	}

	if regexpIpv4.MatchString(domain) {
		a.appendError(fmt.Sprintf(errMsgNotValid, value, "email"), msgArgs...)
		return false
	}

	return true
}
