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
	a.GreaterThan(1, 1)
	a.GreaterThan(2, 2)

	assert.True(t, a.HasErrors())
	assert.Equal(t, 2, a.CountErrors())
	assert.EqualError(t, a.ErrorAt(0), "1 is not greater than 1")
	assert.EqualError(t, a.ErrorAt(-2), "1 is not greater than 1")
	assert.EqualError(t, a.ErrorAt(1), "2 is not greater than 2")
	assert.EqualError(t, a.ErrorAt(-1), "2 is not greater than 2")
	assert.Nil(t, a.ErrorAt(2))
	assert.Nil(t, a.ErrorAt(-3))
}

func assertAllReturnsTrue(t *testing.T, data []MethodDataOK) {
	for _, i := range data {
		t.Run(fmt.Sprintf("%s %v", i.method, i.okArgs), func(t *testing.T) {
			assertMethodReturnsTrue(t, i.method, i.okArgs)
		})
	}
}

func assertAllReturnsFalse(t *testing.T, data []MethodDataKO) {
	for _, i := range data {
		t.Run(fmt.Sprintf("%s %v", i.method, i.koArgs), func(t *testing.T) {
			assertMethodReturnsFalse(t, i.method, i.koArgs, i.errMsg)
		})
	}
}

func assertMethodReturnsTrue(t *testing.T, method string, okArgs []interface{}) {
	assertMethodMeetsExpectations(t, method, okArgs, true)
}

func assertMethodReturnsFalse(t *testing.T, method string, koArgs []interface{}, errMsg string) {
	assertMethodMeetsExpectations(t, method, koArgs, false, errMsg)
	assertMethodMeetsExpectations(t, method, append(koArgs, "custom error"), false, "custom error")
	assertMethodMeetsExpectations(t, method, append(koArgs, "%s", "custom error"), false, "custom error")
}

func assertMethodMeetsExpectations(t *testing.T, method string, params []interface{}, valid bool, err ...string) {
	a := New()
	m := reflect.ValueOf(&a).MethodByName(method)
	f := m.Type()

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		v := reflect.ValueOf(param)
		if !v.IsValid() {
			v = reflect.ValueOf(((*int)(nil)))
		}
		if k >= f.NumIn()-1 {
			in[k] = v
			continue
		}

		if f.In(k).Kind() == reflect.TypeOf(param).Kind() {
			in[k] = v
			continue
		}

		in[k] = v.Convert(f.In(k))
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
