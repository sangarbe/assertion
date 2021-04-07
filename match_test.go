package assertion

import (
	"strings"
	"testing"
)

func TestAssertion_Alfanum_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"Alfanum", []interface{}{"abc123"}},
		{"Alfanum", []interface{}{"ABC098"}},
		{"Alfanum", []interface{}{"España"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_Alfanum_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"Alfanum", []interface{}{"abc 123"}, "abc 123 is not alfa-numeric"},
		{"Alfanum", []interface{}{"abc.123"}, "abc.123 is not alfa-numeric"},
		{"Alfanum", []interface{}{"abc#123"}, "abc#123 is not alfa-numeric"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_Digits_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"Digits", []interface{}{"0123456789"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_Digits_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"Digits", []interface{}{"abc123"}, "abc123 is not only digits"},
		{"Digits", []interface{}{"123.456"}, "123.456 is not only digits"},
		{"Digits", []interface{}{"123#"}, "123# is not only digits"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_Letters_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"Letters", []interface{}{"abcDEFáöß"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_Letters_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"Letters", []interface{}{"abc123"}, "abc123 is not only letters"},
		{"Letters", []interface{}{"abc.DEF"}, "abc.DEF is not only letters"},
		{"Letters", []interface{}{"abc DEF"}, "abc DEF is not only letters"},
		{"Letters", []interface{}{"abc#"}, "abc# is not only letters"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_Email_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
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
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_Email_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
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
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_Phone_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"Phone", []interface{}{"+33626525690"}},
		{"Phone", []interface{}{"33626525690"}},
		{"Phone", []interface{}{"+16174552211"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_Phone_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"Phone", []interface{}{"+3362652569e"}, "+3362652569e is not a valid phone"},
		{"Phone", []interface{}{"+3361231231232652569"}, "+3361231231232652569 is not a valid phone"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_Ipv4_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"Ipv4", []interface{}{"127.0.0.1"}},
		{"Ipv4", []interface{}{"255.255.255.255"}},
		{"Ipv4", []interface{}{"0.0.0.0"}},
		{"Ipv4", []interface{}{"199.160.1.10"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_Ipv4_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"Ipv4", []interface{}{"127.0.01"}, "127.0.01 is not a valid ipv4"},
		{"Ipv4", []interface{}{"256.0.0.1"}, "256.0.0.1 is not a valid ipv4"},
		{"Ipv4", []interface{}{"0.0.0.0.0"}, "0.0.0.0.0 is not a valid ipv4"},
		{"Ipv4", []interface{}{"0.0.0.1234"}, "0.0.0.1234 is not a valid ipv4"},
		{"Ipv4", []interface{}{"0-0-0-0"}, "0-0-0-0 is not a valid ipv4"},
	}

	assertAllReturnsFalse(t, data)
}
