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

func TestAssertion_IntGreaterThan(t *testing.T) {
	a := New()
	a.IntGreaterThan(1, 0)

	assert.False(t, a.HasErrors())
	assert.Equal(t, 0, a.CountErrors())

	a.IntGreaterThan(1, 1)

	assert.True(t, a.HasErrors())
	assert.Equal(t, 1, a.CountErrors())
}

func TestAssertion_ErrorAt(t *testing.T) {
	a := New()
	a.IntGreaterThan(1, 1)
	a.IntGreaterThan(2, 2)

	assert.True(t, a.HasErrors())
	assert.Equal(t, 2, a.CountErrors())
	assert.EqualError(t, a.ErrorAt(0), "1 is not greater than 1")
	assert.EqualError(t, a.ErrorAt(-2), "1 is not greater than 1")
	assert.EqualError(t, a.ErrorAt(1), "2 is not greater than 2")
	assert.EqualError(t, a.ErrorAt(-1), "2 is not greater than 2")
	assert.Nil(t, a.ErrorAt(2))
	assert.Nil(t, a.ErrorAt(-3))
}
