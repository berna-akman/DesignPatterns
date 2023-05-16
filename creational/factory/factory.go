package factory

import "fmt"

type IPizza interface {
	getName() string
	getCookingTime() uint
}

func (p *Pizza) getName() string {
	return p.Name
}

func (p *Pizza) getCookingTime() uint {
	return p.CookingTime
}

func GetPizza(pizzaType string) (IPizza, error) {
	if pizzaType == "Neapolitan" {
		return NewNeapolitan(), nil
	}

	if pizzaType == "New York Style" {
		return NewNYStyle(), nil
	}

	return nil, fmt.Errorf("wrong pizza type")
}
