package assertion

import (
	"fmt"
	"reflect"
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

// StartsWithInsensitive returns true if a given string starts with the given
// needle substring with insensitive case
func (a *Assertion) StartsWithInsensitive(value, needle string, msgArgs ...interface{}) bool {
	if !strings.HasPrefix(strings.ToLower(value), strings.ToLower(needle)) {
		a.addErrorMsg(fmt.Sprintf(errMsgNotStartsWith, value, needle), msgArgs...)
		return false
	}

	return true
}

// EndsWithInsensitive returns true if a given string ends with the given needle substring
func (a *Assertion) EndsWithInsensitive(value, needle string, msgArgs ...interface{}) bool {
	if !strings.HasSuffix(strings.ToLower(value), strings.ToLower(needle)) {
		a.addErrorMsg(fmt.Sprintf(errMsgNotEndsWith, value, needle), msgArgs...)
		return false
	}

	return true
}

// ContainsInsensitive returns true if a given string ends with the given needle substring
func (a *Assertion) ContainsInsensitive(value, needle string, msgArgs ...interface{}) bool {
	if !strings.Contains(strings.ToLower(value), strings.ToLower(needle)) {
		a.addErrorMsg(fmt.Sprintf(errMsgNotContains, value, needle), msgArgs...)
		return false
	}

	return true
}

// HasKey returns true if a given key exists on the a given map
func (a *Assertion) HasKey(value interface{}, key interface{}, msgArgs ...interface{}) bool {

	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			if reflect.DeepEqual(k.Interface(), key){
				return true
			}
		}
	}

	a.addErrorMsg(fmt.Sprintf(errMsgNotHasKey, value, key), msgArgs...)
	return false
}
