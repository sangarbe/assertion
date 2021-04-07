package assertion

import (
	"fmt"
	"strings"
)

// StartsWith returns true if a given string starts with the given needle substring
func (a *Assertion) StartsWith(value, needle string, msgArgs ...interface{}) bool {
	if !strings.HasPrefix(value, needle) {
		a.addErrorMsg(fmt.Sprintf(errMsgNotStartsWith, value, needle), msgArgs...)
		return false
	}

	return true
}

// EndsWith returns true if a given string ends with the given needle substring
func (a *Assertion) EndsWith(value, needle string, msgArgs ...interface{}) bool {
	if !strings.HasSuffix(value, needle) {
		a.addErrorMsg(fmt.Sprintf(errMsgNotEndsWith, value, needle), msgArgs...)
		return false
	}

	return true
}

// Contains returns true if a given string ends with the given needle substring
func (a *Assertion) Contains(value, needle string, msgArgs ...interface{}) bool {
	if !strings.Contains(value, needle) {
		a.addErrorMsg(fmt.Sprintf(errMsgNotContains, value, needle), msgArgs...)
		return false
	}

	return true
}
