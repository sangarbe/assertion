package assertion

import (
	"fmt"
	"testing"
)

func TestAssertion_CompareMethodsReturnOk(t *testing.T) {
	data := []struct {
		method string
		okArgs []interface{}
	}{
		{"EqualBool", []interface{}{true, true}},
		{"True", []interface{}{true}},
		{"False", []interface{}{false}},
		{"EqualInt", []interface{}{1, 1}},
		{"EqualInt8", []interface{}{1, 1}},
		{"EqualInt16", []interface{}{1, 1}},
		{"EqualInt32", []interface{}{1, 1}},
		{"EqualInt64", []interface{}{1, 1}},
		{"EqualUint", []interface{}{1, 1}},
		{"EqualUint8", []interface{}{1, 1}},
		{"EqualUint16", []interface{}{1, 1}},
		{"EqualUint32", []interface{}{1, 1}},
		{"EqualUint64", []interface{}{1, 1}},
		{"EqualFloat32", []interface{}{1.5, 1.5}},
		{"EqualFloat64", []interface{}{1.5, 1.5}},
		{"EqualString", []interface{}{"a", "a"}},
		{"GreaterThanInt64", []interface{}{1, 0}},
		{"GreaterThanInt32", []interface{}{1, 0}},
		{"GreaterThanInt16", []interface{}{1, 0}},
		{"GreaterThanInt8", []interface{}{1, 0}},
		{"GreaterThanInt", []interface{}{1, 0}},
		{"GreaterThanUint64", []interface{}{1, 0}},
		{"GreaterThanUint32", []interface{}{1, 0}},
		{"GreaterThanUint16", []interface{}{1, 0}},
		{"GreaterThanUint8", []interface{}{1, 0}},
		{"GreaterThanUint", []interface{}{1, 0}},
		{"GreaterThanFloat64", []interface{}{1, 0}},
		{"GreaterThanFloat32", []interface{}{1, 0}},
		{"GreaterThanString", []interface{}{"b", "a"}},
		{"LowerThanInt64", []interface{}{0, 1}},
		{"LowerThanInt32", []interface{}{0, 1}},
		{"LowerThanInt16", []interface{}{0, 1}},
		{"LowerThanInt8", []interface{}{0, 1}},
		{"LowerThanInt", []interface{}{0, 1}},
		{"LowerThanUint64", []interface{}{0, 1}},
		{"LowerThanUint32", []interface{}{0, 1}},
		{"LowerThanUint16", []interface{}{0, 1}},
		{"LowerThanUint8", []interface{}{0, 1}},
		{"LowerThanUint", []interface{}{0, 1}},
		{"LowerThanFloat64", []interface{}{0, 1}},
		{"LowerThanFloat32", []interface{}{0, 1}},
		{"LowerThanString", []interface{}{"a", "b"}},
		{"GreaterThanOrEqualInt64", []interface{}{1, 0}},
		{"GreaterThanOrEqualInt64", []interface{}{0, 0}},
		{"GreaterThanOrEqualInt32", []interface{}{1, 0}},
		{"GreaterThanOrEqualInt32", []interface{}{0, 0}},
		{"GreaterThanOrEqualInt16", []interface{}{1, 0}},
		{"GreaterThanOrEqualInt16", []interface{}{0, 0}},
		{"GreaterThanOrEqualInt8", []interface{}{1, 0}},
		{"GreaterThanOrEqualInt8", []interface{}{0, 0}},
		{"GreaterThanOrEqualInt", []interface{}{1, 0}},
		{"GreaterThanOrEqualInt", []interface{}{0, 0}},
		{"GreaterThanOrEqualUint64", []interface{}{1, 0}},
		{"GreaterThanOrEqualUint64", []interface{}{0, 0}},
		{"GreaterThanOrEqualUint32", []interface{}{1, 0}},
		{"GreaterThanOrEqualUint32", []interface{}{0, 0}},
		{"GreaterThanOrEqualUint16", []interface{}{1, 0}},
		{"GreaterThanOrEqualUint16", []interface{}{0, 0}},
		{"GreaterThanOrEqualUint8", []interface{}{1, 0}},
		{"GreaterThanOrEqualUint8", []interface{}{0, 0}},
		{"GreaterThanOrEqualUint", []interface{}{1, 0}},
		{"GreaterThanOrEqualUint", []interface{}{0, 0}},
		{"GreaterThanOrEqualFloat64", []interface{}{1, 0}},
		{"GreaterThanOrEqualFloat64", []interface{}{0, 0}},
		{"GreaterThanOrEqualFloat32", []interface{}{1, 0}},
		{"GreaterThanOrEqualFloat32", []interface{}{0, 0}},
		{"GreaterThanOrEqualString", []interface{}{"b", "a"}},
		{"GreaterThanOrEqualString", []interface{}{"a", "a"}},
		{"LowerThanOrEqualInt64", []interface{}{0, 1}},
		{"LowerThanOrEqualInt64", []interface{}{1, 1}},
		{"LowerThanOrEqualInt32", []interface{}{0, 1}},
		{"LowerThanOrEqualInt32", []interface{}{1, 1}},
		{"LowerThanOrEqualInt16", []interface{}{0, 1}},
		{"LowerThanOrEqualInt16", []interface{}{1, 1}},
		{"LowerThanOrEqualInt8", []interface{}{0, 1}},
		{"LowerThanOrEqualInt8", []interface{}{1, 1}},
		{"LowerThanOrEqualInt", []interface{}{0, 1}},
		{"LowerThanOrEqualInt", []interface{}{1, 1}},
		{"LowerThanOrEqualUint64", []interface{}{0, 1}},
		{"LowerThanOrEqualUint64", []interface{}{1, 1}},
		{"LowerThanOrEqualUint32", []interface{}{0, 1}},
		{"LowerThanOrEqualUint32", []interface{}{1, 1}},
		{"LowerThanOrEqualUint16", []interface{}{0, 1}},
		{"LowerThanOrEqualUint16", []interface{}{1, 1}},
		{"LowerThanOrEqualUint8", []interface{}{0, 1}},
		{"LowerThanOrEqualUint8", []interface{}{1, 1}},
		{"LowerThanOrEqualUint", []interface{}{0, 1}},
		{"LowerThanOrEqualUint", []interface{}{1, 1}},
		{"LowerThanOrEqualFloat64", []interface{}{0, 1}},
		{"LowerThanOrEqualFloat64", []interface{}{1, 1}},
		{"LowerThanOrEqualFloat32", []interface{}{0, 1}},
		{"LowerThanOrEqualFloat32", []interface{}{1, 1}},
		{"LowerThanOrEqualString", []interface{}{"a", "b"}},
		{"LowerThanOrEqualString", []interface{}{"a", "a"}},
		{"BetweenInt64", []interface{}{0, 0, 1}},
		{"BetweenInt64", []interface{}{1, 0, 1}},
		{"BetweenInt32", []interface{}{0, 0, 1}},
		{"BetweenInt32", []interface{}{1, 0, 1}},
		{"BetweenInt16", []interface{}{0, 0, 1}},
		{"BetweenInt16", []interface{}{1, 0, 1}},
		{"BetweenInt8", []interface{}{0, 0, 1}},
		{"BetweenInt8", []interface{}{1, 0, 1}},
		{"BetweenInt", []interface{}{0, 0, 1}},
		{"BetweenInt", []interface{}{1, 0, 1}},
		{"BetweenUint64", []interface{}{0, 0, 1}},
		{"BetweenUint64", []interface{}{1, 0, 1}},
		{"BetweenUint32", []interface{}{0, 0, 1}},
		{"BetweenUint32", []interface{}{1, 0, 1}},
		{"BetweenUint16", []interface{}{0, 0, 1}},
		{"BetweenUint16", []interface{}{1, 0, 1}},
		{"BetweenUint8", []interface{}{0, 0, 1}},
		{"BetweenUint8", []interface{}{1, 0, 1}},
		{"BetweenUint", []interface{}{0, 0, 1}},
		{"BetweenUint", []interface{}{1, 0, 1}},
		{"BetweenFloat64", []interface{}{0, 0, 1}},
		{"BetweenFloat64", []interface{}{1, 0, 1}},
		{"BetweenFloat32", []interface{}{0, 0, 1}},
		{"BetweenFloat32", []interface{}{1, 0, 1}},
		{"BetweenExcludeInt64", []interface{}{1, 0, 2}},
		{"BetweenExcludeInt32", []interface{}{1, 0, 2}},
		{"BetweenExcludeInt16", []interface{}{1, 0, 2}},
		{"BetweenExcludeInt8", []interface{}{1, 0, 2}},
		{"BetweenExcludeInt", []interface{}{1, 0, 2}},
		{"BetweenExcludeUint64", []interface{}{1, 0, 2}},
		{"BetweenExcludeUint32", []interface{}{1, 0, 2}},
		{"BetweenExcludeUint16", []interface{}{1, 0, 2}},
		{"BetweenExcludeUint8", []interface{}{1, 0, 2}},
		{"BetweenExcludeUint", []interface{}{1, 0, 2}},
		{"BetweenExcludeFloat64", []interface{}{1, 0, 2}},
		{"BetweenExcludeFloat32", []interface{}{1, 0, 2}},
	}

	for _, i := range data {
		t.Run(fmt.Sprintf("%s %v", i.method, i.okArgs), func(t *testing.T) {
			assertMethodReturnsOk(t, i.method, i.okArgs)
		})
	}
}

func TestAssertion_CompareMethodsReturnKo(t *testing.T) {
	data := []struct {
		method string
		koArgs []interface{}
		errMsg string
	}{
		{"EqualBool", []interface{}{true, false}, "true is not equal false"},
		{"True", []interface{}{false}, "false is not equal true"},
		{"False", []interface{}{true}, "true is not equal false"},
		{"EqualInt", []interface{}{1, 0}, "1 is not equal 0"},
		{"EqualInt8", []interface{}{1, 0}, "1 is not equal 0"},
		{"EqualInt16", []interface{}{1, 0}, "1 is not equal 0"},
		{"EqualInt32", []interface{}{1, 0}, "1 is not equal 0"},
		{"EqualInt64", []interface{}{1, 0}, "1 is not equal 0"},
		{"EqualUint", []interface{}{1, 0}, "1 is not equal 0"},
		{"EqualUint8", []interface{}{1, 0}, "1 is not equal 0"},
		{"EqualUint16", []interface{}{1, 0}, "1 is not equal 0"},
		{"EqualUint32", []interface{}{1, 0}, "1 is not equal 0"},
		{"EqualUint64", []interface{}{1, 0}, "1 is not equal 0"},
		{"EqualFloat32", []interface{}{1.5, 0.5}, "1.5 is not equal 0.5"},
		{"EqualFloat64", []interface{}{1.5, 0.5}, "1.5 is not equal 0.5"},
		{"EqualString", []interface{}{"a", "b"}, "a is not equal b"},
		{"GreaterThanInt", []interface{}{1, 1}, "1 is not greater than 1"},
		{"GreaterThanInt8", []interface{}{1, 1}, "1 is not greater than 1"},
		{"GreaterThanInt16", []interface{}{1, 1}, "1 is not greater than 1"},
		{"GreaterThanInt32", []interface{}{1, 1}, "1 is not greater than 1"},
		{"GreaterThanInt64", []interface{}{1, 1}, "1 is not greater than 1"},
		{"GreaterThanUint", []interface{}{1, 1}, "1 is not greater than 1"},
		{"GreaterThanUint8", []interface{}{1, 1}, "1 is not greater than 1"},
		{"GreaterThanUint16", []interface{}{1, 1}, "1 is not greater than 1"},
		{"GreaterThanUint32", []interface{}{1, 1}, "1 is not greater than 1"},
		{"GreaterThanUint64", []interface{}{1, 1}, "1 is not greater than 1"},
		{"GreaterThanFloat32", []interface{}{1, 1}, "1 is not greater than 1"},
		{"GreaterThanFloat64", []interface{}{1, 1}, "1 is not greater than 1"},
		{"GreaterThanString", []interface{}{"a", "b"}, "a is not greater than b"},
		{"LowerThanInt", []interface{}{1, 1}, "1 is not lower than 1"},
		{"LowerThanInt8", []interface{}{1, 1}, "1 is not lower than 1"},
		{"LowerThanInt16", []interface{}{1, 1}, "1 is not lower than 1"},
		{"LowerThanInt32", []interface{}{1, 1}, "1 is not lower than 1"},
		{"LowerThanInt64", []interface{}{1, 1}, "1 is not lower than 1"},
		{"LowerThanUint", []interface{}{1, 1}, "1 is not lower than 1"},
		{"LowerThanUint8", []interface{}{1, 1}, "1 is not lower than 1"},
		{"LowerThanUint16", []interface{}{1, 1}, "1 is not lower than 1"},
		{"LowerThanUint32", []interface{}{1, 1}, "1 is not lower than 1"},
		{"LowerThanUint64", []interface{}{1, 1}, "1 is not lower than 1"},
		{"LowerThanFloat32", []interface{}{1, 1}, "1 is not lower than 1"},
		{"LowerThanFloat64", []interface{}{1, 1}, "1 is not lower than 1"},
		{"LowerThanString", []interface{}{"a", "a"}, "a is not lower than a"},
		{"GreaterThanOrEqualInt", []interface{}{0, 1}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqualInt8", []interface{}{0, 1}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqualInt16", []interface{}{0, 1}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqualInt32", []interface{}{0, 1}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqualInt64", []interface{}{0, 1}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqualUint", []interface{}{0, 1}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqualUint8", []interface{}{0, 1}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqualUint16", []interface{}{0, 1}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqualUint32", []interface{}{0, 1}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqualUint64", []interface{}{0, 1}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqualFloat32", []interface{}{0, 1}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqualFloat64", []interface{}{0, 1}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqualString", []interface{}{"a", "b"}, "a is not greater than or equal b"},
		{"LowerThanOrEqualInt", []interface{}{1, 0}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqualInt8", []interface{}{1, 0}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqualInt16", []interface{}{1, 0}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqualInt32", []interface{}{1, 0}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqualInt64", []interface{}{1, 0}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqualUint", []interface{}{1, 0}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqualUint8", []interface{}{1, 0}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqualUint16", []interface{}{1, 0}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqualUint32", []interface{}{1, 0}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqualUint64", []interface{}{1, 0}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqualFloat32", []interface{}{1, 0}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqualFloat64", []interface{}{1, 0}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqualString", []interface{}{"b", "a"}, "b is not lower than or equal a"},
		{"BetweenInt", []interface{}{1, 0, 0}, "1 is not between 0 and 0"},
		{"BetweenInt8", []interface{}{1, 0, 0}, "1 is not between 0 and 0"},
		{"BetweenInt16", []interface{}{1, 0, 0}, "1 is not between 0 and 0"},
		{"BetweenInt32", []interface{}{1, 0, 0}, "1 is not between 0 and 0"},
		{"BetweenInt64", []interface{}{1, 0, 0}, "1 is not between 0 and 0"},
		{"BetweenUint", []interface{}{1, 0, 0}, "1 is not between 0 and 0"},
		{"BetweenUint8", []interface{}{1, 0, 0}, "1 is not between 0 and 0"},
		{"BetweenUint16", []interface{}{1, 0, 0}, "1 is not between 0 and 0"},
		{"BetweenUint32", []interface{}{1, 0, 0}, "1 is not between 0 and 0"},
		{"BetweenUint64", []interface{}{1, 0, 0}, "1 is not between 0 and 0"},
		{"BetweenFloat32", []interface{}{1, 0, 0}, "1 is not between 0 and 0"},
		{"BetweenFloat64", []interface{}{1, 0, 0}, "1 is not between 0 and 0"},
		{"BetweenInt", []interface{}{0, 1, 1}, "0 is not between 1 and 1"},
		{"BetweenInt8", []interface{}{0, 1, 1}, "0 is not between 1 and 1"},
		{"BetweenInt16", []interface{}{0, 1, 1}, "0 is not between 1 and 1"},
		{"BetweenInt32", []interface{}{0, 1, 1}, "0 is not between 1 and 1"},
		{"BetweenInt64", []interface{}{0, 1, 1}, "0 is not between 1 and 1"},
		{"BetweenUint", []interface{}{0, 1, 1}, "0 is not between 1 and 1"},
		{"BetweenUint8", []interface{}{0, 1, 1}, "0 is not between 1 and 1"},
		{"BetweenUint16", []interface{}{0, 1, 1}, "0 is not between 1 and 1"},
		{"BetweenUint32", []interface{}{0, 1, 1}, "0 is not between 1 and 1"},
		{"BetweenUint64", []interface{}{0, 1, 1}, "0 is not between 1 and 1"},
		{"BetweenFloat32", []interface{}{0, 1, 1}, "0 is not between 1 and 1"},
		{"BetweenFloat64", []interface{}{0, 1, 1}, "0 is not between 1 and 1"},
		{"BetweenExcludeInt", []interface{}{1, 0, 1}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExcludeInt8", []interface{}{1, 0, 1}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExcludeInt16", []interface{}{1, 0, 1}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExcludeInt32", []interface{}{1, 0, 1}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExcludeInt64", []interface{}{1, 0, 1}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExcludeUint", []interface{}{1, 0, 1}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExcludeUint8", []interface{}{1, 0, 1}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExcludeUint16", []interface{}{1, 0, 1}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExcludeUint32", []interface{}{1, 0, 1}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExcludeUint64", []interface{}{1, 0, 1}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExcludeFloat32", []interface{}{1, 0, 1}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExcludeFloat64", []interface{}{1, 0, 1}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExcludeInt", []interface{}{0, 0, 1}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExcludeInt8", []interface{}{0, 0, 1}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExcludeInt16", []interface{}{0, 0, 1}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExcludeInt32", []interface{}{0, 0, 1}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExcludeInt64", []interface{}{0, 0, 1}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExcludeUint", []interface{}{0, 0, 1}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExcludeUint8", []interface{}{0, 0, 1}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExcludeUint16", []interface{}{0, 0, 1}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExcludeUint32", []interface{}{0, 0, 1}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExcludeUint64", []interface{}{0, 0, 1}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExcludeFloat32", []interface{}{0, 0, 1}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExcludeFloat64", []interface{}{0, 0, 1}, "0 is not between 0 and 1 both excluded"},
	}

	for _, i := range data {
		t.Run(fmt.Sprintf("%s %v", i.method, i.koArgs), func(t *testing.T) {
			assertMethodReturnsKo(t, i.method, i.koArgs, i.errMsg)
		})
	}
}
