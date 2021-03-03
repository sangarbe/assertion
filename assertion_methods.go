package assertion

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
	"unicode"
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
