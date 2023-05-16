package factory

type Pizza struct {
	Name        string
	CookingTime uint
}

type Neapolitan struct {
	Pizza
}

type NYStyle struct {
	Pizza
}

func NewNeapolitan() IPizza {
	return &Neapolitan{
		Pizza: Pizza{
			Name:        "Neapolitan",
			CookingTime: 17,
		},
	}
}

func NewNYStyle() IPizza {
	return &NYStyle{
		Pizza: Pizza{
			Name:        "New Yok Style Pizza",
			CookingTime: 19,
		},
	}
}
