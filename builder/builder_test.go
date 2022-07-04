package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuilder_Build(t *testing.T) {
	fullHuman := NewHuman().AddHeight(170).AddAge(23).AddEyeColor("brown").AddName("John").Build()

	assert.Equal(t, "John is 23 years old, is 170cm tall and has brown eyes", fullHuman)
}
