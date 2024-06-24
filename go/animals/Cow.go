package animals

import "math/rand"

type Cow struct {
	AnimalBreed, Product string
}

func (cow *Cow) Init() Cow {
	return Cow{
		AnimalBreed: "cows",
		Product:     "milk",
	}
}
func (cow *Cow) GetOutputProduct() int {
	return rand.Intn(5) + 8
}
