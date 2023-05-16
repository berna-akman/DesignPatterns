package builder

import "fmt"

type IBuilder interface {
	setSize(size string)
	addCheese(add bool)
	addPepperoni(add bool)
	addMushrooms(add bool)
}

func (b *Builder) Build() *Pizza {
	return &Pizza{
		Size:      b.Size,
		Cheese:    b.Cheese,
		Pepperoni: b.Pepperoni,
		Mushrooms: b.Mushrooms,
	}
}

func (b *Builder) setSize(size string) *Builder {
	b.Size = size
	return b
}

func (b *Builder) addCheese(add bool) *Builder {
	b.Cheese = add
	return b
}

func (b *Builder) addPepperoni(add bool) *Builder {
	b.Pepperoni = add
	return b
}

func (b *Builder) addMushrooms(add bool) *Builder {
	b.Mushrooms = add
	return b
}

func BuildDifferentPizzas() {
	builder := Builder{}
	pizza1 := builder.setSize("Large").addCheese(true).addMushrooms(false).addPepperoni(true).Build()
	pizza2 := builder.setSize("Medium").addCheese(true).addMushrooms(true).addPepperoni(true).Build()
	fmt.Println("pizza1:", pizza1)
	fmt.Println("pizza2:", pizza2)
}
