package bridge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactory_TriangleVolume(t *testing.T) {
	triangle := TriangularShape{
		height: float32(12),
		base:   float32(5.00),
		length: float32(6.00),
		color:  "blue",
	}

	shape := ShapeManipulator{
		shape: &triangle,
	}

	volume := shape.Volume()

	assert.Equal(t, float32(180), volume)
}

func TestFactory_SquareVolume(t *testing.T) {
	square := SquareShape{
		height: 4,
		length: 4,
		color:  "red",
	}

	shape := ShapeManipulator{
		&square,
	}

	volume := shape.Volume()

	assert.Equal(t, float32(64), volume)
}

func TestFactory_TriangleColor(t *testing.T) {
	triangle := TriangularShape{
		height: float32(12),
		base:   float32(5.00),
		length: float32(6.00),
		color:  "blue",
	}

	calc := ShapeManipulator{
		shape: &triangle,
	}

	color := calc.Color()

	assert.Equal(t, "This is a blue shape", color)
}

func TestFactory_SqaureColor(t *testing.T) {
	square := SquareShape{
		height: 4,
		length: 4,
		color:  "red",
	}

	shape := ShapeManipulator{
		&square,
	}

	volume := shape.Color()

	assert.Equal(t, "This is a red shape", volume)
}
