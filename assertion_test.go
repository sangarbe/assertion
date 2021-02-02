package assertion

import (
	"fmt"
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

func TestAssertion_AllAssertMethods(t *testing.T) {
	data := []struct {
		method string
		okArgs []interface{}
		koArgs []interface{}
		errMsg string
	}{
		{"EqualBool", []interface{}{true, true}, []interface{}{true, false}, "true is not false"},
		{"True", []interface{}{true}, []interface{}{false}, "false is not true"},
		{"False", []interface{}{false}, []interface{}{true}, "true is not false"},
		{"Boolean", []interface{}{"true"}, []interface{}{"on"}, "on is not a valid boolean string"},
		{"Boolean", []interface{}{"TRUE"}, []interface{}{"on"}, "on is not a valid boolean string"},
		{"Boolean", []interface{}{"t"}, []interface{}{"on"}, "on is not a valid boolean string"},
		{"Boolean", []interface{}{"T"}, []interface{}{"on"}, "on is not a valid boolean string"},
		{"Boolean", []interface{}{"1"}, []interface{}{"on"}, "on is not a valid boolean string"},
		{"Boolean", []interface{}{"false"}, []interface{}{"off"}, "off is not a valid boolean string"},
		{"Boolean", []interface{}{"FALSE"}, []interface{}{"off"}, "off is not a valid boolean string"},
		{"Boolean", []interface{}{"f"}, []interface{}{"off"}, "off is not a valid boolean string"},
		{"Boolean", []interface{}{"F"}, []interface{}{"off"}, "off is not a valid boolean string"},
		{"Boolean", []interface{}{"0"}, []interface{}{"off"}, "off is not a valid boolean string"},
		{"Truthy", []interface{}{"true"}, []interface{}{"on"}, "on is not a valid truthy string"},
		{"Truthy", []interface{}{"TRUE"}, []interface{}{"on"}, "on is not a valid truthy string"},
		{"Truthy", []interface{}{"t"}, []interface{}{"on"}, "on is not a valid truthy string"},
		{"Truthy", []interface{}{"T"}, []interface{}{"on"}, "on is not a valid truthy string"},
		{"Truthy", []interface{}{"1"}, []interface{}{"on"}, "on is not a valid truthy string"},
		{"Falsy", []interface{}{"false"}, []interface{}{"off"}, "off is not a valid falsy string"},
		{"Falsy", []interface{}{"FALSE"}, []interface{}{"off"}, "off is not a valid falsy string"},
		{"Falsy", []interface{}{"f"}, []interface{}{"off"}, "off is not a valid falsy string"},
		{"Falsy", []interface{}{"F"}, []interface{}{"off"}, "off is not a valid falsy string"},
		{"Falsy", []interface{}{"0"}, []interface{}{"off"}, "off is not a valid falsy string"},
		{"EqualInt", []interface{}{1, 1}, []interface{}{1, 0}, "1 is not equal 0"},
		//{"EqualInt64", []interface{}{1, 1}, []interface{}{1, 0}, "1 is not equal 0"},
	}

	for _, i := range data {
		t.Run(fmt.Sprintf("%s %v%v", i.method, i.okArgs, i.koArgs), func(t *testing.T) {
			assertMethod(t, i.method, i.okArgs, i.koArgs, i.errMsg)
		})
	}
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

func assertMethod(t *testing.T, method string, okArgs []interface{}, koArgs []interface{}, errMsg string) {
	assertMethodMeetsExpectations(t, method, okArgs, true)
	assertMethodMeetsExpectations(t, method, koArgs, false, errMsg)
	assertMethodMeetsExpectations(t, method, append(koArgs, "custom error"), false, "custom error")
	assertMethodMeetsExpectations(t, method, append(koArgs, "%s", "custom error"), false, "custom error")
}

func assertMethodMeetsExpectations(t *testing.T, method string, params []interface{}, valid bool, err ...string) {
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
