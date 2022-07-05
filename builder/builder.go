package builder

type Human struct {
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
	Build() Human
}

// NewHuman We're going to use this so we don't leak the actual builder
func NewHuman() HumanBuilder {
	return &Human{}
}

func (h *Human) AddAge(age int) HumanBuilder {
	h.age = age
	return h
}

func (h *Human) AddHeight(height int) HumanBuilder {
	h.height = height

	return h
}

func (h *Human) AddEyeColor(color string) HumanBuilder {
	h.eyeColor = color

	return h
}

func (h *Human) AddName(name string) HumanBuilder {
	h.name = name

	return h
}
func (h *Human) Build() Human {
	return Human{
		age:      h.age,
		height:   h.height,
		eyeColor: h.eyeColor,
		name:     h.name,
	}
}
