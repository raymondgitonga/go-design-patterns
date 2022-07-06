package bridge

import (
	"fmt"
)

// Allows separation of abstraction from its implementation

type Shape interface {
	Height() float32
	Perimeter() float32
	Color() string
}

type TriangularShape struct {
	height float32
	base   float32
	length float32
	color  string
}

type SquareShape struct {
	height float32
	length float32
	color  string
}

type ShapeManipulator struct {
	shape Shape
}

func (t *TriangularShape) Perimeter() float32 {
	return (t.height * t.base) / 2
}

func (t *TriangularShape) Height() float32 {
	return t.length
}

func (t *TriangularShape) Color() string {
	return t.color
}

func (s *SquareShape) Perimeter() float32 {
	return s.height * 4
}

func (s *SquareShape) Height() float32 {
	return s.height
}

func (s *SquareShape) Color() string {
	return s.color
}

//Implementation

func (v *ShapeManipulator) Volume() float32 {
	return v.shape.Perimeter() * v.shape.Height()
}

func (v *ShapeManipulator) Color() string {
	return fmt.Sprintf("This is a %s shape", v.shape.Color())
}
