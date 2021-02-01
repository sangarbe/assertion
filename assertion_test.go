package assertion

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestAssertion_New(t *testing.T) {
	a := New()

	assert.False(t, a.HasErrors())
	assert.Equal(t, 0, a.CountErrors())
}

func TestAssertion_ErrorAt(t *testing.T) {
	a := New()
	a.GreaterThanInt(1, 1)
	a.GreaterThanInt(2, 2)

	assert.True(t, a.HasErrors())
	assert.Equal(t, 2, a.CountErrors())
	assert.EqualError(t, a.ErrorAt(0), "1 is not greater than 1")
	assert.EqualError(t, a.ErrorAt(-2), "1 is not greater than 1")
	assert.EqualError(t, a.ErrorAt(1), "2 is not greater than 2")
	assert.EqualError(t, a.ErrorAt(-1), "2 is not greater than 2")
	assert.Nil(t, a.ErrorAt(2))
	assert.Nil(t, a.ErrorAt(-3))
}

func TestAssertion_EqualBool(t *testing.T) {
	assertMethod(t, "EqualBool", []interface{}{true, true}, true)
	assertMethod(t, "EqualBool", []interface{}{false, true}, false, "false is not true")
	assertMethod(t, "EqualBool", []interface{}{false, true, "custom error"}, false, "custom error")
	assertMethod(t, "EqualBool", []interface{}{false, true, "%s", "custom error"}, false, "custom error")
}

func TestAssertion_True(t *testing.T) {
	assertMethod(t, "True", []interface{}{true}, true)
	assertMethod(t, "True", []interface{}{false}, false, "false is not true")
	assertMethod(t, "True", []interface{}{false, "custom error"}, false, "custom error")
	assertMethod(t, "True", []interface{}{false, "%s", "custom error"}, false, "custom error")
}

func TestAssertion_False(t *testing.T) {
	assertMethod(t, "False", []interface{}{false}, true)
	assertMethod(t, "False", []interface{}{true}, false, "true is not false")
	assertMethod(t, "False", []interface{}{true, "custom error"}, false, "custom error")
	assertMethod(t, "False", []interface{}{true, "%s", "custom error"}, false, "custom error")
}

func TestAssertion_GreaterThanInt(t *testing.T) {
	a := New()
	a.GreaterThanInt(1, 0)

	assert.False(t, a.HasErrors())
	assert.Equal(t, 0, a.CountErrors())

	a.GreaterThanInt(1, 1)

	assert.True(t, a.HasErrors())
	assert.Equal(t, 1, a.CountErrors())
}

func TestAssertion_EqualInt64(t *testing.T) {
	a := New()

	assert.True(t, a.EqualInt64(1, 1))
	assert.False(t, a.HasErrors())
	assert.Equal(t, 0, a.CountErrors())

	assert.False(t, a.EqualInt64(1, 2))
	assert.True(t, a.HasErrors())
	assert.Equal(t, 1, a.CountErrors())
	assert.EqualError(t, a.ErrorAt(0), "1 is not equal 2")
}

func TestAssertion_EqualInt32(t *testing.T) {
	a := New()

	assert.True(t, a.EqualInt32(1, 1))
	assert.False(t, a.HasErrors())
	assert.Equal(t, 0, a.CountErrors())

	assert.False(t, a.EqualInt32(1, 2))
	assert.True(t, a.HasErrors())
	assert.Equal(t, 1, a.CountErrors())
	assert.EqualError(t, a.ErrorAt(0), "1 is not equal 2")
}

func assertMethod(t *testing.T, method string, params []interface{}, valid bool, err ...string) {
	a := New()
	m := reflect.ValueOf(&a).MethodByName(method)

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	var result []reflect.Value
	result = m.Call(in)

	assert.Equal(t, valid, result[0].Bool())

	errcount := 0
	if !valid {
		errcount = 1
	}

	assert.Equal(t, !valid, a.HasErrors())
	assert.Equal(t, errcount, a.CountErrors())

	if !valid && len(err) > 0 {
		assert.EqualError(t, a.ErrorAt(0), err[0])
	}
}