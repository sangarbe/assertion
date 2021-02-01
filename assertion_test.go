package assertion

import (
	"github.com/stretchr/testify/assert"
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
	a := New()

	assert.True(t, a.EqualBool(true, true))
	assert.False(t, a.HasErrors())
	assert.Equal(t, 0, a.CountErrors())

	assert.False(t, a.EqualBool(false, true))
	assert.True(t, a.HasErrors())
	assert.Equal(t, 1, a.CountErrors())
	assert.EqualError(t, a.ErrorAt(0), "false is not true")

	assert.False(t, a.EqualBool(false, true, "custom error"))
	assert.True(t, a.HasErrors())
	assert.Equal(t, 2, a.CountErrors())
	assert.EqualError(t, a.ErrorAt(1), "custom error")

	assert.False(t, a.EqualBool(false, true, "custom error %s", "dummy"))
	assert.True(t, a.HasErrors())
	assert.Equal(t, 3, a.CountErrors())
	assert.EqualError(t, a.ErrorAt(2), "custom error dummy")
}

func TestAssertion_True(t *testing.T) {
	a := New()

	assert.True(t, a.True(true))
	assert.False(t, a.HasErrors())
	assert.Equal(t, 0, a.CountErrors())

	assert.False(t, a.True(false))
	assert.True(t, a.HasErrors())
	assert.Equal(t, 1, a.CountErrors())
	assert.EqualError(t, a.ErrorAt(0), "false is not true")
}

func TestAssertion_False(t *testing.T) {
	a := New()

	assert.True(t, a.False(false))
	assert.False(t, a.HasErrors())
	assert.Equal(t, 0, a.CountErrors())

	assert.False(t, a.False(true))
	assert.True(t, a.HasErrors())
	assert.Equal(t, 1, a.CountErrors())
	assert.EqualError(t, a.ErrorAt(0), "true is not false")
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
