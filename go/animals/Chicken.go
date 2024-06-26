package animals

import "math/rand"

type Chicken struct {
	AnimalBreed, Product string
}

func (chicken *Chicken) Init() Chicken {
	return Chicken{
		AnimalBreed: "chickens",
		Product: "eggs",
	}
}

// Returns a random amount of products.
func (chicken *Chicken) GetOutputProduct() int {
	return rand.Intn(2)
}
