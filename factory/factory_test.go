package factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactory_SimpleInterest(t *testing.T) {
	si := NewCalculator(1, 10000, 12).Calculate()
	assert.Equal(t, 1452.00, si)
}

func TestFactory_CompoundInterest(t *testing.T) {
	ci := int(NewCalculator(2, 1000, 12).Calculate())
	assert.Equal(t, 9223372036854775807, ci)
}
