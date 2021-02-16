package assertion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MethodDataOK struct {
	method string
	okArgs []interface{}
}

type MethodDataKO struct {
	method string
	koArgs []interface{}
	errMsg string
}

func TestAssertion_Nil_ReturnsTrue(t *testing.T) {
	var (
		nilPtr *int
		nilSlc []int
		nilMap map[int]int
		nilChn chan int
		nilFun func()
		nilItf interface{}
	)

	data := []MethodDataOK{
		{"Nil", []interface{}{nilPtr}},
		{"Nil", []interface{}{nilMap}},
		{"Nil", []interface{}{nilChn}},
		{"Nil", []interface{}{nilFun}},
		{"Nil", []interface{}{nilSlc}},
		{"Nil", []interface{}{nilItf}},
	}

	assertAllReturnsTrue(t, data)

	a := New()
	assert.True(t, a.Nil(nil))
}

func TestAssertion_Nil_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"Nil", []interface{}{true}, "true is not <nil>"},
		{"Nil", []interface{}{0}, "0 is not <nil>"},
		{"Nil", []interface{}{""}, " is not <nil>"},
		{"Nil", []interface{}{struct{}{}}, "{} is not <nil>"},
		{"Nil", []interface{}{' '}, "32 is not <nil>"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_Equal_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"Equal", []interface{}{true, true}},
		{"Equal", []interface{}{false, false}},
		{"Equal", []interface{}{int64(1), int64(1)}},
		{"Equal", []interface{}{int32(1), int32(1)}},
		{"Equal", []interface{}{int16(1), int16(1)}},
		{"Equal", []interface{}{int8(1), int8(1)}},
		{"Equal", []interface{}{1, 1}},
		{"Equal", []interface{}{uint64(1), uint64(1)}},
		{"Equal", []interface{}{uint32(1), uint32(1)}},
		{"Equal", []interface{}{uint16(1), uint16(1)}},
		{"Equal", []interface{}{uint8(1), uint8(1)}},
		{"Equal", []interface{}{uint(1), uint(1)}},
		{"Equal", []interface{}{float32(1.5), float32(1.5)}},
		{"Equal", []interface{}{float64(1.5), float64(1.5)}},
		{"Equal", []interface{}{"a", "a"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_Equal_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"Equal", []interface{}{true, false}, "true is not equal false"},
		{"Equal", []interface{}{int64(1), int64(0)}, "1 is not equal 0"},
		{"Equal", []interface{}{int32(1), int32(0)}, "1 is not equal 0"},
		{"Equal", []interface{}{int16(1), int16(0)}, "1 is not equal 0"},
		{"Equal", []interface{}{int8(1), int8(0)}, "1 is not equal 0"},
		{"Equal", []interface{}{1, 0}, "1 is not equal 0"},
		{"Equal", []interface{}{uint64(1), uint64(0)}, "1 is not equal 0"},
		{"Equal", []interface{}{uint32(1), uint32(0)}, "1 is not equal 0"},
		{"Equal", []interface{}{uint16(1), uint16(0)}, "1 is not equal 0"},
		{"Equal", []interface{}{uint8(1), uint8(0)}, "1 is not equal 0"},
		{"Equal", []interface{}{uint(1), uint(0)}, "1 is not equal 0"},
		{"Equal", []interface{}{float32(1.5), float32(0.5)}, "1.5 is not equal 0.5"},
		{"Equal", []interface{}{float64(1.5), float64(0.5)}, "1.5 is not equal 0.5"},
		{"Equal", []interface{}{"a", "b"}, "a is not equal b"},
		{"Equal", []interface{}{"a", 1}, "a and 1 are not of the same type"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_True_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"True", []interface{}{true}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_True_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"True", []interface{}{false}, "false is not equal true"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_False_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"False", []interface{}{false}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_False_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"False", []interface{}{true}, "true is not equal false"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_GreaterThan_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"GreaterThan", []interface{}{1, 0}},
		{"GreaterThan", []interface{}{int64(1), int64(0)}},
		{"GreaterThan", []interface{}{int32(1), int32(0)}},
		{"GreaterThan", []interface{}{int16(1), int16(0)}},
		{"GreaterThan", []interface{}{int8(1), int8(0)}},
		{"GreaterThan", []interface{}{uint64(1), uint64(0)}},
		{"GreaterThan", []interface{}{uint32(1), uint32(0)}},
		{"GreaterThan", []interface{}{uint16(1), uint16(0)}},
		{"GreaterThan", []interface{}{uint8(1), uint8(0)}},
		{"GreaterThan", []interface{}{uint(1), uint(0)}},
		{"GreaterThan", []interface{}{float32(1), float32(0)}},
		{"GreaterThan", []interface{}{float64(1), float64(0)}},
		{"GreaterThan", []interface{}{"b", "a"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_GreaterThan_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"GreaterThan", []interface{}{1, 1}, "1 is not greater than 1"},
		{"GreaterThan", []interface{}{int64(1), int64(1)}, "1 is not greater than 1"},
		{"GreaterThan", []interface{}{int32(1), int32(1)}, "1 is not greater than 1"},
		{"GreaterThan", []interface{}{int16(1), int16(1)}, "1 is not greater than 1"},
		{"GreaterThan", []interface{}{int8(1), int8(1)}, "1 is not greater than 1"},
		{"GreaterThan", []interface{}{uint64(1), uint64(1)}, "1 is not greater than 1"},
		{"GreaterThan", []interface{}{uint32(1), uint32(1)}, "1 is not greater than 1"},
		{"GreaterThan", []interface{}{uint16(1), uint16(1)}, "1 is not greater than 1"},
		{"GreaterThan", []interface{}{uint8(1), uint8(1)}, "1 is not greater than 1"},
		{"GreaterThan", []interface{}{uint(1), uint(1)}, "1 is not greater than 1"},
		{"GreaterThan", []interface{}{float32(1), float32(1)}, "1 is not greater than 1"},
		{"GreaterThan", []interface{}{float64(1), float64(1)}, "1 is not greater than 1"},
		{"GreaterThan", []interface{}{0, 1}, "0 is not greater than 1"},
		{"GreaterThan", []interface{}{int64(0), int64(1)}, "0 is not greater than 1"},
		{"GreaterThan", []interface{}{int32(0), int32(1)}, "0 is not greater than 1"},
		{"GreaterThan", []interface{}{int16(0), int16(1)}, "0 is not greater than 1"},
		{"GreaterThan", []interface{}{int8(0), int8(1)}, "0 is not greater than 1"},
		{"GreaterThan", []interface{}{uint64(0), uint64(1)}, "0 is not greater than 1"},
		{"GreaterThan", []interface{}{uint32(0), uint32(1)}, "0 is not greater than 1"},
		{"GreaterThan", []interface{}{uint16(0), uint16(1)}, "0 is not greater than 1"},
		{"GreaterThan", []interface{}{uint8(0), uint8(1)}, "0 is not greater than 1"},
		{"GreaterThan", []interface{}{uint(0), uint(1)}, "0 is not greater than 1"},
		{"GreaterThan", []interface{}{float32(0), float32(1)}, "0 is not greater than 1"},
		{"GreaterThan", []interface{}{float64(0), float64(1)}, "0 is not greater than 1"},
		{"GreaterThan", []interface{}{"a", "a"}, "a is not greater than a"},
		{"GreaterThan", []interface{}{"a", "b"}, "a is not greater than b"},
		{"GreaterThan", []interface{}{"a", 1}, "a and 1 are not of the same type"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_LowerThan_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"LowerThan", []interface{}{0, 1}},
		{"LowerThan", []interface{}{int64(0), int64(1)}},
		{"LowerThan", []interface{}{int32(0), int32(1)}},
		{"LowerThan", []interface{}{int16(0), int16(1)}},
		{"LowerThan", []interface{}{int8(0), int8(1)}},
		{"LowerThan", []interface{}{uint64(0), uint64(1)}},
		{"LowerThan", []interface{}{uint32(0), uint32(1)}},
		{"LowerThan", []interface{}{uint16(0), uint16(1)}},
		{"LowerThan", []interface{}{uint8(0), uint8(1)}},
		{"LowerThan", []interface{}{uint(0), uint(1)}},
		{"LowerThan", []interface{}{float32(0), float32(1)}},
		{"LowerThan", []interface{}{float64(0), float64(1)}},
		{"LowerThan", []interface{}{"a", "b"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_LowerThan_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"LowerThan", []interface{}{1, 1}, "1 is not lower than 1"},
		{"LowerThan", []interface{}{int64(1), int64(1)}, "1 is not lower than 1"},
		{"LowerThan", []interface{}{int32(1), int32(1)}, "1 is not lower than 1"},
		{"LowerThan", []interface{}{int16(1), int16(1)}, "1 is not lower than 1"},
		{"LowerThan", []interface{}{int8(1), int8(1)}, "1 is not lower than 1"},
		{"LowerThan", []interface{}{uint64(1), uint64(1)}, "1 is not lower than 1"},
		{"LowerThan", []interface{}{uint32(1), uint32(1)}, "1 is not lower than 1"},
		{"LowerThan", []interface{}{uint16(1), uint16(1)}, "1 is not lower than 1"},
		{"LowerThan", []interface{}{uint8(1), uint8(1)}, "1 is not lower than 1"},
		{"LowerThan", []interface{}{uint(1), uint(1)}, "1 is not lower than 1"},
		{"LowerThan", []interface{}{float32(1), float32(1)}, "1 is not lower than 1"},
		{"LowerThan", []interface{}{float64(1), float64(1)}, "1 is not lower than 1"},
		{"LowerThan", []interface{}{1, 0}, "1 is not lower than 0"},
		{"LowerThan", []interface{}{int64(1), int64(0)}, "1 is not lower than 0"},
		{"LowerThan", []interface{}{int32(1), int32(0)}, "1 is not lower than 0"},
		{"LowerThan", []interface{}{int16(1), int16(0)}, "1 is not lower than 0"},
		{"LowerThan", []interface{}{int8(1), int8(0)}, "1 is not lower than 0"},
		{"LowerThan", []interface{}{uint64(1), uint64(0)}, "1 is not lower than 0"},
		{"LowerThan", []interface{}{uint32(1), uint32(0)}, "1 is not lower than 0"},
		{"LowerThan", []interface{}{uint16(1), uint16(0)}, "1 is not lower than 0"},
		{"LowerThan", []interface{}{uint8(1), uint8(0)}, "1 is not lower than 0"},
		{"LowerThan", []interface{}{uint(1), uint(0)}, "1 is not lower than 0"},
		{"LowerThan", []interface{}{float32(1), float32(0)}, "1 is not lower than 0"},
		{"LowerThan", []interface{}{float64(1), float64(0)}, "1 is not lower than 0"},
		{"LowerThan", []interface{}{"a", "a"}, "a is not lower than a"},
		{"LowerThan", []interface{}{"b", "a"}, "b is not lower than a"},
		{"LowerThan", []interface{}{"a", 1}, "a and 1 are not of the same type"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_GreaterThanOrEqual_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"GreaterThanOrEqual", []interface{}{1, 0}},
		{"GreaterThanOrEqual", []interface{}{int64(1), int64(0)}},
		{"GreaterThanOrEqual", []interface{}{int32(1), int32(0)}},
		{"GreaterThanOrEqual", []interface{}{int16(1), int16(0)}},
		{"GreaterThanOrEqual", []interface{}{int8(1), int8(0)}},
		{"GreaterThanOrEqual", []interface{}{uint64(1), uint64(0)}},
		{"GreaterThanOrEqual", []interface{}{uint32(1), uint32(0)}},
		{"GreaterThanOrEqual", []interface{}{uint16(1), uint16(0)}},
		{"GreaterThanOrEqual", []interface{}{uint8(1), uint8(0)}},
		{"GreaterThanOrEqual", []interface{}{uint(1), uint(0)}},
		{"GreaterThanOrEqual", []interface{}{float32(1), float32(0)}},
		{"GreaterThanOrEqual", []interface{}{float64(1), float64(0)}},
		{"GreaterThanOrEqual", []interface{}{"b", "a"}},
		{"GreaterThanOrEqual", []interface{}{1, 1}},
		{"GreaterThanOrEqual", []interface{}{int64(1), int64(1)}},
		{"GreaterThanOrEqual", []interface{}{int32(1), int32(1)}},
		{"GreaterThanOrEqual", []interface{}{int16(1), int16(1)}},
		{"GreaterThanOrEqual", []interface{}{int8(1), int8(1)}},
		{"GreaterThanOrEqual", []interface{}{uint64(1), uint64(1)}},
		{"GreaterThanOrEqual", []interface{}{uint32(1), uint32(1)}},
		{"GreaterThanOrEqual", []interface{}{uint16(1), uint16(1)}},
		{"GreaterThanOrEqual", []interface{}{uint8(1), uint8(1)}},
		{"GreaterThanOrEqual", []interface{}{uint(1), uint(1)}},
		{"GreaterThanOrEqual", []interface{}{float32(1), float32(1)}},
		{"GreaterThanOrEqual", []interface{}{float64(1), float64(1)}},
		{"GreaterThanOrEqual", []interface{}{"a", "a"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_GreaterThanOrEqual_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"GreaterThanOrEqual", []interface{}{0, 1}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqual", []interface{}{int64(0), int64(1)}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqual", []interface{}{int32(0), int32(1)}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqual", []interface{}{int16(0), int16(1)}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqual", []interface{}{int8(0), int8(1)}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqual", []interface{}{uint64(0), uint64(1)}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqual", []interface{}{uint32(0), uint32(1)}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqual", []interface{}{uint16(0), uint16(1)}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqual", []interface{}{uint8(0), uint8(1)}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqual", []interface{}{uint(0), uint(1)}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqual", []interface{}{float32(0), float32(1)}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqual", []interface{}{float64(0), float64(1)}, "0 is not greater than or equal 1"},
		{"GreaterThanOrEqual", []interface{}{"a", "b"}, "a is not greater than or equal b"},
		{"GreaterThanOrEqual", []interface{}{"a", 1}, "a and 1 are not of the same type"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_LowerThanOrEqual_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"LowerThanOrEqual", []interface{}{0, 1}},
		{"LowerThanOrEqual", []interface{}{int64(0), int64(1)}},
		{"LowerThanOrEqual", []interface{}{int32(0), int32(1)}},
		{"LowerThanOrEqual", []interface{}{int16(0), int16(1)}},
		{"LowerThanOrEqual", []interface{}{int8(0), int8(1)}},
		{"LowerThanOrEqual", []interface{}{uint64(0), uint64(1)}},
		{"LowerThanOrEqual", []interface{}{uint32(0), uint32(1)}},
		{"LowerThanOrEqual", []interface{}{uint16(0), uint16(1)}},
		{"LowerThanOrEqual", []interface{}{uint8(0), uint8(1)}},
		{"LowerThanOrEqual", []interface{}{uint(0), uint(1)}},
		{"LowerThanOrEqual", []interface{}{float32(0), float32(1)}},
		{"LowerThanOrEqual", []interface{}{float64(0), float64(1)}},
		{"LowerThanOrEqual", []interface{}{"a", "b"}},
		{"LowerThanOrEqual", []interface{}{1, 1}},
		{"LowerThanOrEqual", []interface{}{int64(1), int64(1)}},
		{"LowerThanOrEqual", []interface{}{int32(1), int32(1)}},
		{"LowerThanOrEqual", []interface{}{int16(1), int16(1)}},
		{"LowerThanOrEqual", []interface{}{int8(1), int8(1)}},
		{"LowerThanOrEqual", []interface{}{uint64(1), uint64(1)}},
		{"LowerThanOrEqual", []interface{}{uint32(1), uint32(1)}},
		{"LowerThanOrEqual", []interface{}{uint16(1), uint16(1)}},
		{"LowerThanOrEqual", []interface{}{uint8(1), uint8(1)}},
		{"LowerThanOrEqual", []interface{}{uint(1), uint(1)}},
		{"LowerThanOrEqual", []interface{}{float32(1), float32(1)}},
		{"LowerThanOrEqual", []interface{}{float64(1), float64(1)}},
		{"LowerThanOrEqual", []interface{}{"a", "a"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_LowerThanOrEqual_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"LowerThanOrEqual", []interface{}{1, 0}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqual", []interface{}{int64(1), int64(0)}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqual", []interface{}{int32(1), int32(0)}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqual", []interface{}{int16(1), int16(0)}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqual", []interface{}{int8(1), int8(0)}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqual", []interface{}{uint64(1), uint64(0)}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqual", []interface{}{uint32(1), uint32(0)}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqual", []interface{}{uint16(1), uint16(0)}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqual", []interface{}{uint8(1), uint8(0)}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqual", []interface{}{uint(1), uint(0)}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqual", []interface{}{float32(1), float32(0)}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqual", []interface{}{float64(1), float64(0)}, "1 is not lower than or equal 0"},
		{"LowerThanOrEqual", []interface{}{"b", "a"}, "b is not lower than or equal a"},
		{"LowerThanOrEqual", []interface{}{"a", 1}, "a and 1 are not of the same type"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_Between_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"Between", []interface{}{0, 0, 1}},
		{"Between", []interface{}{int64(0), int64(0), int64(1)}},
		{"Between", []interface{}{int32(0), int32(0), int32(1)}},
		{"Between", []interface{}{int16(0), int16(0), int16(1)}},
		{"Between", []interface{}{int8(0), int8(0), int8(1)}},
		{"Between", []interface{}{uint64(0), uint64(0), uint64(1)}},
		{"Between", []interface{}{uint32(0), uint32(0), uint32(1)}},
		{"Between", []interface{}{uint16(0), uint16(0), uint16(1)}},
		{"Between", []interface{}{uint8(0), uint8(0), uint8(1)}},
		{"Between", []interface{}{uint(0), uint(0), uint(1)}},
		{"Between", []interface{}{float32(0), float32(0), float32(1)}},
		{"Between", []interface{}{float64(0), float64(0), float64(1)}},
		{"Between", []interface{}{"a", "a", "b"}},
		{"Between", []interface{}{1, 0, 1}},
		{"Between", []interface{}{int64(1), int64(0), int64(1)}},
		{"Between", []interface{}{int32(1), int32(0), int32(1)}},
		{"Between", []interface{}{int16(1), int16(0), int16(1)}},
		{"Between", []interface{}{int8(1), int8(0), int8(1)}},
		{"Between", []interface{}{uint64(1), uint64(0), uint64(1)}},
		{"Between", []interface{}{uint32(1), uint32(0), uint32(1)}},
		{"Between", []interface{}{uint16(1), uint16(0), uint16(1)}},
		{"Between", []interface{}{uint8(1), uint8(0), uint8(1)}},
		{"Between", []interface{}{uint(1), uint(0), uint(1)}},
		{"Between", []interface{}{float32(1), float32(0), float32(1)}},
		{"Between", []interface{}{float64(1), float64(0), float64(1)}},
		{"Between", []interface{}{"b", "a", "b"}},
		{"Between", []interface{}{1, 0, 2}},
		{"Between", []interface{}{int64(1), int64(0), int64(2)}},
		{"Between", []interface{}{int32(1), int32(0), int32(2)}},
		{"Between", []interface{}{int16(1), int16(0), int16(2)}},
		{"Between", []interface{}{int8(1), int8(0), int8(2)}},
		{"Between", []interface{}{uint64(1), uint64(0), uint64(2)}},
		{"Between", []interface{}{uint32(1), uint32(0), uint32(2)}},
		{"Between", []interface{}{uint16(1), uint16(0), uint16(2)}},
		{"Between", []interface{}{uint8(1), uint8(0), uint8(2)}},
		{"Between", []interface{}{uint(1), uint(0), uint(2)}},
		{"Between", []interface{}{float32(1), float32(0), float32(2)}},
		{"Between", []interface{}{float64(1), float64(0), float64(2)}},
		{"Between", []interface{}{"b", "a", "c"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_Between_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"Between", []interface{}{1, 0, 0}, "1 is not between 0 and 0"},
		{"Between", []interface{}{int64(1), int64(0), int64(0)}, "1 is not between 0 and 0"},
		{"Between", []interface{}{int32(1), int32(0), int32(0)}, "1 is not between 0 and 0"},
		{"Between", []interface{}{int16(1), int16(0), int16(0)}, "1 is not between 0 and 0"},
		{"Between", []interface{}{int8(1), int8(0), int8(0)}, "1 is not between 0 and 0"},
		{"Between", []interface{}{uint64(1), uint64(0), uint64(0)}, "1 is not between 0 and 0"},
		{"Between", []interface{}{uint32(1), uint32(0), uint32(0)}, "1 is not between 0 and 0"},
		{"Between", []interface{}{uint16(1), uint16(0), uint16(0)}, "1 is not between 0 and 0"},
		{"Between", []interface{}{uint8(1), uint8(0), uint8(0)}, "1 is not between 0 and 0"},
		{"Between", []interface{}{uint(1), uint(0), uint(0)}, "1 is not between 0 and 0"},
		{"Between", []interface{}{float32(1), float32(0), float32(0)}, "1 is not between 0 and 0"},
		{"Between", []interface{}{float64(1), float64(0), float64(0)}, "1 is not between 0 and 0"},
		{"Between", []interface{}{"b", "a", "a"}, "b is not between a and a"},
		{"Between", []interface{}{"a", 1, 1}, "a is not between 1 and 1"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_BetweenExclude_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{

		{"BetweenExclude", []interface{}{1, 0, 2}},
		{"BetweenExclude", []interface{}{int64(1), int64(0), int64(2)}},
		{"BetweenExclude", []interface{}{int32(1), int32(0), int32(2)}},
		{"BetweenExclude", []interface{}{int16(1), int16(0), int16(2)}},
		{"BetweenExclude", []interface{}{int8(1), int8(0), int8(2)}},
		{"BetweenExclude", []interface{}{uint64(1), uint64(0), uint64(2)}},
		{"BetweenExclude", []interface{}{uint32(1), uint32(0), uint32(2)}},
		{"BetweenExclude", []interface{}{uint16(1), uint16(0), uint16(2)}},
		{"BetweenExclude", []interface{}{uint8(1), uint8(0), uint8(2)}},
		{"BetweenExclude", []interface{}{uint(1), uint(0), uint(2)}},
		{"BetweenExclude", []interface{}{float32(1), float32(0), float32(2)}},
		{"BetweenExclude", []interface{}{float64(1), float64(0), float64(2)}},
		{"BetweenExclude", []interface{}{"b", "a", "c"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_BetweenExclude_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"BetweenExclude", []interface{}{0, 0, 1}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{int64(0), int64(0), int64(1)}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{int32(0), int32(0), int32(1)}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{int16(0), int16(0), int16(1)}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{int8(0), int8(0), int8(1)}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{uint64(0), uint64(0), uint64(1)}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{uint32(0), uint32(0), uint32(1)}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{uint16(0), uint16(0), uint16(1)}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{uint8(0), uint8(0), uint8(1)}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{uint(0), uint(0), uint(1)}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{float32(0), float32(0), float32(1)}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{float64(0), float64(0), float64(1)}, "0 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{"a", "a", "b"}, "a is not between a and b both excluded"},
		{"BetweenExclude", []interface{}{1, 0, 1}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{int64(1), int64(0), int64(1)}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{int32(1), int32(0), int32(1)}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{int16(1), int16(0), int16(1)}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{int8(1), int8(0), int8(1)}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{uint64(1), uint64(0), uint64(1)}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{uint32(1), uint32(0), uint32(1)}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{uint16(1), uint16(0), uint16(1)}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{uint8(1), uint8(0), uint8(1)}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{uint(1), uint(0), uint(1)}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{float32(1), float32(0), float32(1)}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{float64(1), float64(0), float64(1)}, "1 is not between 0 and 1 both excluded"},
		{"BetweenExclude", []interface{}{"b", "a", "b"}, "b is not between a and b both excluded"},
		{"BetweenExclude", []interface{}{1, 0, 0}, "1 is not between 0 and 0 both excluded"},
		{"BetweenExclude", []interface{}{int64(1), int64(0), int64(0)}, "1 is not between 0 and 0 both excluded"},
		{"BetweenExclude", []interface{}{int32(1), int32(0), int32(0)}, "1 is not between 0 and 0 both excluded"},
		{"BetweenExclude", []interface{}{int16(1), int16(0), int16(0)}, "1 is not between 0 and 0 both excluded"},
		{"BetweenExclude", []interface{}{int8(1), int8(0), int8(0)}, "1 is not between 0 and 0 both excluded"},
		{"BetweenExclude", []interface{}{uint64(1), uint64(0), uint64(0)}, "1 is not between 0 and 0 both excluded"},
		{"BetweenExclude", []interface{}{uint32(1), uint32(0), uint32(0)}, "1 is not between 0 and 0 both excluded"},
		{"BetweenExclude", []interface{}{uint16(1), uint16(0), uint16(0)}, "1 is not between 0 and 0 both excluded"},
		{"BetweenExclude", []interface{}{uint8(1), uint8(0), uint8(0)}, "1 is not between 0 and 0 both excluded"},
		{"BetweenExclude", []interface{}{uint(1), uint(0), uint(0)}, "1 is not between 0 and 0 both excluded"},
		{"BetweenExclude", []interface{}{float32(1), float32(0), float32(0)}, "1 is not between 0 and 0 both excluded"},
		{"BetweenExclude", []interface{}{float64(1), float64(0), float64(0)}, "1 is not between 0 and 0 both excluded"},
		{"BetweenExclude", []interface{}{"b", "a", "a"}, "b is not between a and a both excluded"},
		{"BetweenExclude", []interface{}{"a", 1, 1}, "a is not between 1 and 1 both excluded"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_VariadicMethods_PanicIfInvalidArgumentCount(t *testing.T) {
	errString := "missing required arguments"

	a := New()
	assert.PanicsWithError(t, errString, func() {
		a.Nil()
	})
	assert.PanicsWithError(t, errString, func() {
		a.Equal(1)
	})
	assert.PanicsWithError(t, errString, func() {
		a.GreaterThan(1)
	})
	assert.PanicsWithError(t, errString, func() {
		a.LowerThan(1)
	})
	assert.PanicsWithError(t, errString, func() {
		a.GreaterThanOrEqual(1)
	})
	assert.PanicsWithError(t, errString, func() {
		a.LowerThanOrEqual(1)
	})
	assert.PanicsWithError(t, errString, func() {
		a.Between(1, 1)
	})
	assert.PanicsWithError(t, errString, func() {
		a.BetweenExclude(1, 1)
	})

}
