package main

import (
	"design-patterns/creational/abstract_factory"
	"design-patterns/creational/builder"
	"design-patterns/creational/factory"
	"design-patterns/creational/prototype"
	"fmt"
)

const (
	Neapolitan = "Neapolitan"
	NYStyle    = "New York Style"
)

func main() {

	// FACTORY PATTERN
	nyStyle, _ := factory.GetPizza(NYStyle)
	neapolitan, _ := factory.GetPizza(Neapolitan)

	fmt.Println("pizza1:", nyStyle)
	fmt.Println("pizza2:", neapolitan)
	fmt.Println()

	// ABSTRACT FACTORY PATTERN
	// Having activities in Horror genre
	reading := &abstract_factory.Activity{}
	reading.HaveActivity(&abstract_factory.HorrorFactory{})

	fmt.Println()

	// Having activities in Sci-Fi genre
	watching := &abstract_factory.Activity{}
	watching.HaveActivity(&abstract_factory.SciFiFactory{})

	fmt.Println()

	// Having activities in Adventure genre
	playing := &abstract_factory.Activity{}
	playing.HaveActivity(&abstract_factory.AdventureFactory{})

	fmt.Println()

	// BUILDER PATTERN
	builder.BuildDifferentPizzas()

	// PROTOTYPE PATTERN
	prototype.GetClone()
}
