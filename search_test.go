package assertion

import (
	"testing"
)

func TestAssertion_StartsWith_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"StartsWith", []interface{}{"Hello world!", "Hello"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_StartsWith_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"StartsWith", []interface{}{"Hello world!", "hello"}, "Hello world! does not start with hello"},
		{"StartsWith", []interface{}{"Hello world!", "ello"}, "Hello world! does not start with ello"},
		{"StartsWith", []interface{}{"Hello world!", "world!"}, "Hello world! does not start with world!"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_EndsWith_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"EndsWith", []interface{}{"Hello world!", "world!"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_EndsWith_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"EndsWith", []interface{}{"Hello world!", "World!"}, "Hello world! does not end with World!"},
		{"EndsWith", []interface{}{"Hello world!", "world"}, "Hello world! does not end with world"},
		{"EndsWith", []interface{}{"Hello world!", "Hello"}, "Hello world! does not end with Hello"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_Contains_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"Contains", []interface{}{"Hello world!", "world"}},
		{"Contains", []interface{}{"Hello world!", "Hello"}},
		{"Contains", []interface{}{"Hello world!", "world!"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_Contains_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"Contains", []interface{}{"Hello world!", "World"}, "Hello world! does not contain World"},
		{"Contains", []interface{}{"Hello world!", "lloworld"}, "Hello world! does not contain lloworld"},
		{"Contains", []interface{}{"Hello world!", "hello"}, "Hello world! does not contain hello"},
	}

	assertAllReturnsFalse(t, data)
}
