package assertion

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
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

func TestAssertion_AllAssertMethodsReturnOk(t *testing.T) {
	data := []struct {
		method string
		okArgs []interface{}
	}{
		{"EqualBool", []interface{}{true, true}},
		{"True", []interface{}{true}},
		{"False", []interface{}{false}},
		{"Boolean", []interface{}{"true"}},
		{"Boolean", []interface{}{"TRUE"}},
		{"Boolean", []interface{}{"t"}},
		{"Boolean", []interface{}{"T"}},
		{"Boolean", []interface{}{"1"}},
		{"Boolean", []interface{}{"false"}},
		{"Boolean", []interface{}{"FALSE"}},
		{"Boolean", []interface{}{"f"}},
		{"Boolean", []interface{}{"F"}},
		{"Boolean", []interface{}{"0"}},
		{"Truthy", []interface{}{"true"}},
		{"Truthy", []interface{}{"TRUE"}},
		{"Truthy", []interface{}{"t"}},
		{"Truthy", []interface{}{"T"}},
		{"Truthy", []interface{}{"1"}},
		{"Falsy", []interface{}{"false"}},
		{"Falsy", []interface{}{"FALSE"}},
		{"Falsy", []interface{}{"f"}},
		{"Falsy", []interface{}{"F"}},
		{"Falsy", []interface{}{"0"}},
		{"EqualInt", []interface{}{1, 1}},
		{"EqualInt8", []interface{}{1, 1}},
		{"EqualInt16", []interface{}{1, 1}},
		{"EqualInt32", []interface{}{1, 1}},
		{"EqualInt64", []interface{}{1, 1}},
		{"EqualFloat32", []interface{}{1.5, 1.5}},
		{"EqualFloat64", []interface{}{1.5, 1.5}},
		{"GreaterThanInt", []interface{}{1, 0}},
		{"LowerThanInt", []interface{}{0, 1}},
		{"GreaterThanOrEqualInt", []interface{}{1, 0}},
		{"GreaterThanOrEqualInt", []interface{}{1, 1}},
		{"LowerThanOrEqualInt", []interface{}{0, 1}},
		{"LowerThanOrEqualInt", []interface{}{1, 1}},
		{"Email", []interface{}{"test@mail.com"}},
		{"Email", []interface{}{"test@subdomain.mail.com"}},
		{"Email", []interface{}{"0123456789@mail.com"}},
		{"Email", []interface{}{"first.last@mail.com"}},
		{"Email", []interface{}{"first+last@mail.com"}},
		{"Email", []interface{}{"first-last@mail.com"}},
		{"Email", []interface{}{"first_last@mail.com"}},
		{"Email", []interface{}{`"first last"@mail.com`}},
		{"Email", []interface{}{"test@[0.0.0.0]"}},
		{"Email", []interface{}{"email@111.222.333.44444"}},
		{"Email", []interface{}{"test@" + strings.Repeat("subd.", 50) + "com"}},
		{"Ipv4", []interface{}{"127.0.0.1"}},
		{"Ipv4", []interface{}{"255.255.255.255"}},
		{"Ipv4", []interface{}{"0.0.0.0"}},
		{"Ipv4", []interface{}{"199.160.1.10"}},
		{"Alfanum", []interface{}{"abc123"}},
		{"Alfanum", []interface{}{"ABC098"}},
		{"Alfanum", []interface{}{"España"}},
		{"Base64", []interface{}{"c29tZSBkYXRhIHdpdGggACBhbmQg77u/"}},
		{"Phone", []interface{}{"+33626525690"}},
		{"Phone", []interface{}{"33626525690"}},
		{"Phone", []interface{}{"+16174552211"}},
	}

	for _, i := range data {
		t.Run(fmt.Sprintf("%s %v", i.method, i.okArgs), func(t *testing.T) {
			assertMethodReturnsOk(t, i.method, i.okArgs)
		})
	}
}

func TestAssertion_AllAssertMethodsReturnKo(t *testing.T) {
	data := []struct {
		method string
		koArgs []interface{}
		errMsg string
	}{
		{"EqualBool", []interface{}{true, false}, "true is not equal false"},
		{"True", []interface{}{false}, "false is not equal true"},
		{"False", []interface{}{true}, "true is not equal false"},
		{"Boolean", []interface{}{"on"}, "on is not a valid boolean string"},
		{"Boolean", []interface{}{"on"}, "on is not a valid boolean string"},
		{"Boolean", []interface{}{"on"}, "on is not a valid boolean string"},
		{"Boolean", []interface{}{"on"}, "on is not a valid boolean string"},
		{"Boolean", []interface{}{"on"}, "on is not a valid boolean string"},
		{"Boolean", []interface{}{"off"}, "off is not a valid boolean string"},
		{"Boolean", []interface{}{"off"}, "off is not a valid boolean string"},
		{"Boolean", []interface{}{"off"}, "off is not a valid boolean string"},
		{"Boolean", []interface{}{"off"}, "off is not a valid boolean string"},
		{"Boolean", []interface{}{"off"}, "off is not a valid boolean string"},
		{"Truthy", []interface{}{"on"}, "on is not a valid truthy string"},
		{"Truthy", []interface{}{"on"}, "on is not a valid truthy string"},
		{"Truthy", []interface{}{"on"}, "on is not a valid truthy string"},
		{"Truthy", []interface{}{"on"}, "on is not a valid truthy string"},
		{"Truthy", []interface{}{"on"}, "on is not a valid truthy string"},
		{"Falsy", []interface{}{"off"}, "off is not a valid falsy string"},
		{"Falsy", []interface{}{"off"}, "off is not a valid falsy string"},
		{"Falsy", []interface{}{"off"}, "off is not a valid falsy string"},
		{"Falsy", []interface{}{"off"}, "off is not a valid falsy string"},
		{"Falsy", []interface{}{"off"}, "off is not a valid falsy string"},
		{"EqualInt", []interface{}{1, 0}, "1 is not equal 0"},
		{"EqualInt8", []interface{}{1, 0}, "1 is not equal 0"},
		{"EqualInt16", []interface{}{1, 0}, "1 is not equal 0"},
		{"EqualInt32", []interface{}{1, 0}, "1 is not equal 0"},
		{"EqualInt64", []interface{}{1, 0}, "1 is not equal 0"},
		{"EqualFloat32", []interface{}{1.5, 0.5}, "1.5 is not equal 0.5"},
		{"EqualFloat64", []interface{}{1.5, 0.5}, "1.5 is not equal 0.5"},
		{"GreaterThanInt", []interface{}{1, 1}, "1 is not greater than 1"},
		{"LowerThanInt", []interface{}{1, 1}, "1 is not lower than 1"},
		{"GreaterThanOrEqualInt", []interface{}{0, 1}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqualInt", []interface{}{0, 1}, "0 is not greater than or equal 1"},
		{"LowerThanOrEqualInt", []interface{}{1, 0}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqualInt", []interface{}{1, 0}, "1 is not lower than or equal 0"},
		{"Email", []interface{}{"plainaddress"}, "plainaddress is not a valid email"},
		{"Email", []interface{}{"#@%^%#$@#$@#.com"}, "#@%^%#$@#$@#.com is not a valid email"},
		{"Email", []interface{}{"@example.com"}, "@example.com is not a valid email"},
		{"Email", []interface{}{"Joe Smith <email@example.com>"}, "Joe Smith <email@example.com> is not a valid email"},
		{"Email", []interface{}{"email.example.com"}, "email.example.com is not a valid email"},
		{"Email", []interface{}{"email@example@example.com"}, "email@example@example.com is not a valid email"},
		{"Email", []interface{}{".email@example.com"}, ".email@example.com is not a valid email"},
		{"Email", []interface{}{"email.@example.com"}, "email.@example.com is not a valid email"},
		{"Email", []interface{}{"email..email@example.com"}, "email..email@example.com is not a valid email"},
		{"Email", []interface{}{"あいうえお@example.com"}, "あいうえお@example.com is not a valid email"},
		{"Email", []interface{}{"email@example.com (Joe Smith)"}, "email@example.com (Joe Smith) is not a valid email"},
		{"Email", []interface{}{"email@example"}, "email@example is not a valid email"},
		{"Email", []interface{}{"email@-example.com"}, "email@-example.com is not a valid email"},
		{"Email", []interface{}{"email@example-.com"}, "email@example-.com is not a valid email"},
		{"Email", []interface{}{"email@example..com"}, "email@example..com is not a valid email"},
		{"Email", []interface{}{"Abc..123@example.com"}, "Abc..123@example.com is not a valid email"},
		{"Email", []interface{}{"first last@example.com"}, "first last@example.com is not a valid email"},
		{"Email", []interface{}{"test@" + strings.Repeat("subd.", 50) + "com.es"}, "test@" + strings.Repeat("subd.", 50) + "com.es is not a valid email"},
		{"Email", []interface{}{"test@0.0.0.0"}, "test@0.0.0.0 is not a valid email"},
		{"Ipv4", []interface{}{"127.0.01"}, "127.0.01 is not a valid ipv4"},
		{"Ipv4", []interface{}{"256.0.0.1"}, "256.0.0.1 is not a valid ipv4"},
		{"Ipv4", []interface{}{"0.0.0.0.0"}, "0.0.0.0.0 is not a valid ipv4"},
		{"Ipv4", []interface{}{"0.0.0.1234"}, "0.0.0.1234 is not a valid ipv4"},
		{"Ipv4", []interface{}{"0-0-0-0"}, "0-0-0-0 is not a valid ipv4"},
		{"Alfanum", []interface{}{"abc 123"}, "abc 123 is not alfa-numeric"},
		{"Alfanum", []interface{}{"abc.123"}, "abc.123 is not alfa-numeric"},
		{"Alfanum", []interface{}{"abc#123"}, "abc#123 is not alfa-numeric"},
		{"Base64", []interface{}{"c29tZSBkYXRhIHdpdGggACBhbmQg77u"}, "c29tZSBkYXRhIHdpdGggACBhbmQg77u is not a valid base64 encoded value"},
		{"Base64", []interface{}{"c29tZSBkYXRhIHdpdGggACBhbmQg77u/,"}, "c29tZSBkYXRhIHdpdGggACBhbmQg77u/, is not a valid base64 encoded value"},
		{"Phone", []interface{}{"+3362652569e"}, "+3362652569e is not a valid phone"},
		{"Phone", []interface{}{"+3361231231232652569"}, "+3361231231232652569 is not a valid phone"},
	}

	for _, i := range data {
		t.Run(fmt.Sprintf("%s %v", i.method, i.koArgs), func(t *testing.T) {
			assertMethodReturnsKo(t, i.method, i.koArgs, i.errMsg)
		})
	}
}

func assertMethodReturnsOk(t *testing.T, method string, okArgs []interface{}) {
	assertMethodMeetsExpectations(t, method, okArgs, true)
}

func assertMethodReturnsKo(t *testing.T, method string, koArgs []interface{}, errMsg string) {
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
