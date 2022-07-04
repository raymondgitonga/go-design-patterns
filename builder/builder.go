package builder

import (
	"fmt"
)

type human struct {
	age      int
	height   int
	eyeColor string
	name     string
}

// HumanBuilder defines the fields we want to set on this builder
type HumanBuilder interface {
	AddAge(age int) HumanBuilder
	AddHeight(height int) HumanBuilder
	AddEyeColor(color string) HumanBuilder
	AddName(name string) HumanBuilder
	Build() string
}

// NewHuman We're going to use this so we don't leak the actual builder
func NewHuman() HumanBuilder {
	return &human{}
}

func (h *human) AddAge(age int) HumanBuilder {
	h.age = age
	return h
}

func (h *human) AddHeight(height int) HumanBuilder {
	h.height = height

	return h
}

func (h *human) AddEyeColor(color string) HumanBuilder {
	h.eyeColor = color

	return h
}

func (h *human) AddName(name string) HumanBuilder {
	h.name = name

	return h
}
func (h *human) Build() string {
	fullHuman := fmt.Sprintf("%s is %d years old, is %dcm tall and has %s eyes", h.name,
		h.age, h.height, h.eyeColor)

	return fullHuman
}
