package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuilder_AddHeight(t *testing.T) {
	fullHuman := NewHuman().AddHeight(170).Build()

	assert.Equal(t, 170, fullHuman.height)
}
