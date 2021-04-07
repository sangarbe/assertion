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

func TestAssertion_StartsWithInsensitive_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"StartsWithInsensitive", []interface{}{"Hello world!", "Hello"}},
		{"StartsWithInsensitive", []interface{}{"Hello world!", "hello"}},
		{"StartsWithInsensitive", []interface{}{"esdrújula", "ESDRÚ"}},
		{"StartsWithInsensitive", []interface{}{"ESDRÚJULA", "esdrú"}},
		{"StartsWithInsensitive", []interface{}{"#especial", "#esp"}},
		{"StartsWithInsensitive", []interface{}{"Straße", "straß"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_StartsWithInsensitive_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"StartsWithInsensitive", []interface{}{"Hello world!", "ello"}, "Hello world! does not start with ello"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_EndsWithInsensitive_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"EndsWithInsensitive", []interface{}{"Hello world!", "world!"}},
		{"EndsWithInsensitive", []interface{}{"Hello WORLD!", "world!"}},
		{"EndsWithInsensitive", []interface{}{"Hello world!", "WORLD!"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_EndsWithInsensitive_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"EndsWithInsensitive", []interface{}{"Hello world!", "world"}, "Hello world! does not end with world"},
		{"EndsWithInsensitive", []interface{}{"Hello world!", "Hello"}, "Hello world! does not end with Hello"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_ContainsInsensitive_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"ContainsInsensitive", []interface{}{"Hello world!", "WORLD"}},
		{"ContainsInsensitive", []interface{}{"Hello WORLD!", "WORLD"}},
		{"ContainsInsensitive", []interface{}{"Hello world!", "HELLO"}},
		{"ContainsInsensitive", []interface{}{"HELLO world!", "hello"}},
		{"ContainsInsensitive", []interface{}{"Hello world!", "World"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_ContainsInsensitive_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"ContainsInsensitive", []interface{}{"Hello world!", "lloworld"}, "Hello world! does not contain lloworld"},
	}

	assertAllReturnsFalse(t, data)
}