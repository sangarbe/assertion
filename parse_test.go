package assertion

import (
	"testing"
)

func TestAssertion_Boolean_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
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
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_Boolean_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"Boolean", []interface{}{"on"}, "on is not a valid boolean string"},
		{"Boolean", []interface{}{"yes"}, "yes is not a valid boolean string"},
		{"Boolean", []interface{}{"y"}, "y is not a valid boolean string"},
		{"Boolean", []interface{}{"off"}, "off is not a valid boolean string"},
		{"Boolean", []interface{}{"no"}, "no is not a valid boolean string"},
		{"Boolean", []interface{}{"n"}, "n is not a valid boolean string"},
		{"Boolean", []interface{}{"ok"}, "ok is not a valid boolean string"},
		{"Boolean", []interface{}{"ko"}, "ko is not a valid boolean string"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_Truthy_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"Truthy", []interface{}{"true"}},
		{"Truthy", []interface{}{"TRUE"}},
		{"Truthy", []interface{}{"t"}},
		{"Truthy", []interface{}{"T"}},
		{"Truthy", []interface{}{"1"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_Truthy_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"Truthy", []interface{}{"on"}, "on is not a valid truthy string"},
		{"Truthy", []interface{}{"yes"}, "yes is not a valid truthy string"},
		{"Truthy", []interface{}{"y"}, "y is not a valid truthy string"},
		{"Truthy", []interface{}{"ok"}, "ok is not a valid truthy string"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_Falsy_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"Falsy", []interface{}{"false"}},
		{"Falsy", []interface{}{"FALSE"}},
		{"Falsy", []interface{}{"f"}},
		{"Falsy", []interface{}{"F"}},
		{"Falsy", []interface{}{"0"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_Falsy_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"Falsy", []interface{}{"off"}, "off is not a valid falsy string"},
		{"Falsy", []interface{}{"off"}, "off is not a valid falsy string"},
		{"Falsy", []interface{}{"no"}, "no is not a valid falsy string"},
		{"Falsy", []interface{}{"n"}, "n is not a valid falsy string"},
		{"Falsy", []interface{}{"ko"}, "ko is not a valid falsy string"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_Integer_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"Integer", []interface{}{"16174552211"}},
		{"Integer", []interface{}{"110110101"}},
		{"Integer", []interface{}{"110_110_101"}},
		{"Integer", []interface{}{"0b110110101"}},
		{"Integer", []interface{}{"0o110110101"}},
		{"Integer", []interface{}{"0x110110101"}},
		{"Integer", []interface{}{"-16174552211"}},
		{"Integer", []interface{}{"-110110101"}},
		{"Integer", []interface{}{"-110_110_101"}},
		{"Integer", []interface{}{"-0b110110101"}},
		{"Integer", []interface{}{"-0o110110101"}},
		{"Integer", []interface{}{"-0x110110101"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_Integer_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"Integer", []interface{}{"0b110110102"}, "0b110110102 is not a valid integer"},
		{"Integer", []interface{}{"0x11011010G"}, "0x11011010G is not a valid integer"},
		{"Integer", []interface{}{"0o110110108"}, "0o110110108 is not a valid integer"},
		{"Integer", []interface{}{"110.110.108"}, "110.110.108 is not a valid integer"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_IntegerBinary_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"IntegerBinary", []interface{}{"110110101"}},
		{"IntegerBinary", []interface{}{"-110110101"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_IntegerBinary_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"IntegerBinary", []interface{}{"0b110110101"}, "0b110110101 is not a valid base-2 integer"},
		{"IntegerBinary", []interface{}{"-0b110110101"}, "-0b110110101 is not a valid base-2 integer"},
		{"IntegerBinary", []interface{}{"110110102"}, "110110102 is not a valid base-2 integer"},
		{"IntegerBinary", []interface{}{"-110110102"}, "-110110102 is not a valid base-2 integer"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_IntegerOctal_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"IntegerOctal", []interface{}{"0767"}},
		{"IntegerOctal", []interface{}{"-7605"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_IntegerOctal_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"IntegerOctal", []interface{}{"108"}, "108 is not a valid base-8 integer"},
		{"IntegerOctal", []interface{}{"-108"}, "-108 is not a valid base-8 integer"},
		{"IntegerOctal", []interface{}{"0o777"}, "0o777 is not a valid base-8 integer"},
		{"IntegerOctal", []interface{}{"-0o777"}, "-0o777 is not a valid base-8 integer"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_IntegerHexadecimal_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"IntegerHexadecimal", []interface{}{"00ff"}},
		{"IntegerHexadecimal", []interface{}{"00FF"}},
		{"IntegerHexadecimal", []interface{}{"-00ff"}},
		{"IntegerHexadecimal", []interface{}{"-00FF"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_IntegerHexadecimal_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"IntegerHexadecimal", []interface{}{"0x00ff"}, "0x00ff is not a valid base-16 integer"},
		{"IntegerHexadecimal", []interface{}{"-0x00ff"}, "-0x00ff is not a valid base-16 integer"},
		{"IntegerHexadecimal", []interface{}{"00fg"}, "00fg is not a valid base-16 integer"},
		{"IntegerHexadecimal", []interface{}{"-00fg"}, "-00fg is not a valid base-16 integer"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_IntegerDecimal_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"IntegerDecimal", []interface{}{"0123456789"}},
		{"IntegerDecimal", []interface{}{"-0123456789"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_IntegerDecimal_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"IntegerDecimal", []interface{}{"111_111_111"}, "111_111_111 is not a valid base-10 integer"},
		{"IntegerDecimal", []interface{}{"-111_111_111"}, "-111_111_111 is not a valid base-10 integer"},
		{"IntegerDecimal", []interface{}{"111.111.111"}, "111.111.111 is not a valid base-10 integer"},
		{"IntegerDecimal", []interface{}{"111 111 111"}, "111 111 111 is not a valid base-10 integer"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_Unsigned_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"Integer", []interface{}{"16174552211"}},
		{"Integer", []interface{}{"110110101"}},
		{"Integer", []interface{}{"110_110_101"}},
		{"Integer", []interface{}{"0b110110101"}},
		{"Integer", []interface{}{"0o110110101"}},
		{"Integer", []interface{}{"0x110110101"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_Unsigned_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"Integer", []interface{}{"0b110110102"}, "0b110110102 is not a valid integer"},
		{"Integer", []interface{}{"0x11011010G"}, "0x11011010G is not a valid integer"},
		{"Integer", []interface{}{"0o110110108"}, "0o110110108 is not a valid integer"},
		{"Integer", []interface{}{"110.110.108"}, "110.110.108 is not a valid integer"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_UnsignedBinary_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"UnsignedBinary", []interface{}{"110110101"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_UnsignedBinary_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"UnsignedBinary", []interface{}{"0b110110101"}, "0b110110101 is not a valid base-2 unsigned integer"},
		{"UnsignedBinary", []interface{}{"-0b110110101"}, "-0b110110101 is not a valid base-2 unsigned integer"},
		{"UnsignedBinary", []interface{}{"110110102"}, "110110102 is not a valid base-2 unsigned integer"},
		{"UnsignedBinary", []interface{}{"-110110102"}, "-110110102 is not a valid base-2 unsigned integer"},
		{"UnsignedBinary", []interface{}{"-110110101"}, "-110110101 is not a valid base-2 unsigned integer"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_UnsignedOctal_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"UnsignedOctal", []interface{}{"0767"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_UnsignedOctal_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"UnsignedOctal", []interface{}{"108"}, "108 is not a valid base-8 unsigned integer"},
		{"UnsignedOctal", []interface{}{"-108"}, "-108 is not a valid base-8 unsigned integer"},
		{"UnsignedOctal", []interface{}{"-0767"}, "-0767 is not a valid base-8 unsigned integer"},
		{"UnsignedOctal", []interface{}{"0o777"}, "0o777 is not a valid base-8 unsigned integer"},
		{"UnsignedOctal", []interface{}{"-0o777"}, "-0o777 is not a valid base-8 unsigned integer"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_UnsignedHexadecimal_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"UnsignedHexadecimal", []interface{}{"00ff"}},
		{"UnsignedHexadecimal", []interface{}{"00FF"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_UnsignedHexadecimal_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"UnsignedHexadecimal", []interface{}{"0x00ff"}, "0x00ff is not a valid base-16 unsigned integer"},
		{"UnsignedHexadecimal", []interface{}{"-0x00ff"}, "-0x00ff is not a valid base-16 unsigned integer"},
		{"UnsignedHexadecimal", []interface{}{"00fg"}, "00fg is not a valid base-16 unsigned integer"},
		{"UnsignedHexadecimal", []interface{}{"-00fg"}, "-00fg is not a valid base-16 unsigned integer"},
		{"UnsignedHexadecimal", []interface{}{"-00FF"}, "-00FF is not a valid base-16 unsigned integer"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_UnsignedDecimal_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"UnsignedDecimal", []interface{}{"0123456789"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_UnsignedDecimal_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"UnsignedDecimal", []interface{}{"111_111_111"}, "111_111_111 is not a valid base-10 unsigned integer"},
		{"UnsignedDecimal", []interface{}{"-111_111_111"}, "-111_111_111 is not a valid base-10 unsigned integer"},
		{"UnsignedDecimal", []interface{}{"111.111.111"}, "111.111.111 is not a valid base-10 unsigned integer"},
		{"UnsignedDecimal", []interface{}{"111 111 111"}, "111 111 111 is not a valid base-10 unsigned integer"},
		{"UnsignedDecimal", []interface{}{"-0123456789"}, "-0123456789 is not a valid base-10 unsigned integer"},
	}

	assertAllReturnsFalse(t, data)
}

func TestAssertion_Float_ReturnsTrue(t *testing.T) {
	data := []MethodDataOK{
		{"Float", []interface{}{"3141516"}},
		{"Float", []interface{}{"3.141516"}},
		{"Float", []interface{}{"-3.141516"}},
		{"Float", []interface{}{"1.234560e+02"}},
		{"Float", []interface{}{"-1.234560e+02"}},
	}

	assertAllReturnsTrue(t, data)
}

func TestAssertion_Float_ReturnsFalse(t *testing.T) {
	data := []MethodDataKO{
		{"Float", []interface{}{"3.14.15"}, "3.14.15 is not a valid float"},
		{"Float", []interface{}{"ffff"}, "ffff is not a valid float"},
	}

	assertAllReturnsFalse(t, data)
}
