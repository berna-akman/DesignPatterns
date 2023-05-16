package prototype

import "fmt"

type Cloneable interface {
	Clone() Cloneable
}

func (c *Car) Clone() Cloneable {
	return &Car{
		Brand:  c.Brand,
		Color:  c.Color,
		Engine: c.Engine,
	}
}

func GetClone() {
	prototype := &Car{
		Brand:  "Tesla",
		Color:  "Red",
		Engine: "Electric",
	}

	cloneCar := prototype.Clone().(*Car)
	cloneCar.Color = "Blue"
	fmt.Println("Prototype: ", prototype)
	fmt.Println("Clone:", cloneCar)
}
