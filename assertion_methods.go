package assertion

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const (
	errMsgNot             = `%v is not %v`
	errMsgNotEqual        = `%v is not equal %v`
	errMsgNotValid        = `%v is not a valid %v`
	errMsgNotGreater      = `%v is not greater than %v`
	errMsgNotLower        = `%v is not lower than %v`
	errMsgNotGreaterEqual = `%v is not greater than or equal %v`
	errMsgNotLowerEqual   = `%v is not lower than or equal %v`
	errMsgNotDifferent    = `%v is not different %v`
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
	rexE164          = `^\+?[1-9]\d{1,14}$`
)

var (
	regexpEmail = regexp.MustCompile(rexEmail)
	regexpIpv4  = regexp.MustCompile(rexIPv4)
	regexpE164  = regexp.MustCompile(rexE164)
)

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

// Boolean returns true if a given string is one of the following accepted forms:
// true, false, TRUE, FALSE, t, f, 1, or 0
func (a *Assertion) Boolean(value string, msgArgs ...interface{}) bool {
	if _, err := strconv.ParseBool(value); err == nil {
		return true
	}

	a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "boolean string"), msgArgs...)
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

// GreaterThanInt returns true if a given int value is greater than other int value
func (a *Assertion) GreaterThanInt(value, other int64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreater, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// GreaterThanOrEqualInt returns true if a given int value is greater than or equal to a min value
func (a *Assertion) GreaterThanOrEqualInt(value, other int64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpGreaterEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanInt returns true if a given int value is lower than a max value
func (a *Assertion) LowerThanInt(value, other int64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLower, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// LowerThanOrEqualInt returns true if a given int value is lower than or equal to a min value
func (a *Assertion) LowerThanOrEqualInt(value, other int64, msgArgs ...interface{}) bool {
	ok, err := compare(cmpOpLowerEqual, value, other, msgArgs...)
	if !ok {
		a.addError(err)
	}

	return ok
}

// Email returns true if a given value is a valid email format. It allows local
// portion to be quoted text and ipv4 for the domain portion (between square brackets).
func (a *Assertion) Email(value string, msgArgs ...interface{}) bool {
	if !regexpEmail.MatchString(value) {
		a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "email"), msgArgs...)
		return false
	}

	splits := strings.Split(value, "@")
	domain := splits[len(splits)-1]
	if len(domain) > 255 {
		a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "email"), msgArgs...)
		return false
	}

	if regexpIpv4.MatchString(domain) {
		a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "email"), msgArgs...)
		return false
	}

	return true
}

// Ipv4 returns true if a given value is a valid ipv4 format
func (a *Assertion) Ipv4(value string, msgArgs ...interface{}) bool {
	if regexpIpv4.MatchString(value) {
		return true
	}

	a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "ipv4"), msgArgs...)
	return false
}

// Alfanum returns true if a given value only contains alfa-numeric runes.
func (a *Assertion) Alfanum(value string, msgArgs ...interface{}) bool {
	for _, r := range value {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			a.addErrorMsg(fmt.Sprintf(errMsgNot, value, "alfa-numeric"), msgArgs...)
			return false
		}
	}

	return true
}

// Base64 returns true if a given value ia a valid base64 encoded string
func (a *Assertion) Base64(value string, msgArgs ...interface{}) bool {
	_, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "base64 encoded value"), msgArgs...)
		return false
	}
	return true
}

// Phone returns true if a given value ia a valid e164 phone number
func (a *Assertion) Phone(value string, msgArgs ...interface{}) bool {
	if !regexpE164.MatchString(value) {
		a.addErrorMsg(fmt.Sprintf(errMsgNotValid, value, "phone"), msgArgs...)
		return false
	}

	return true
}
