package farm

import (
	"farm/animals"
)

type Farm struct {
	Animals           map[string][]animals.Animal
	ProductionPerWeek map[string]int
}

func (farm *Farm) Init() *Farm {
	return &Farm{
		Animals:           make(map[string][]animals.Animal),
		ProductionPerWeek: make(map[string]int),
	}
}

func (farm *Farm) AddAnimal(animalName string, number int) {
	for i := 0; i < number; i++ {
		animalInstance := farm.CreateAnimal(animalName)
		farm.Animals[animalName] = append(farm.Animals[animalName], animalInstance)
	}
}

func (farm *Farm) CreateAnimal(animalName string) animals.Animal {
	if animalName == "cow" {
		obj := (*animals.Cow).Init(new(animals.Cow))
		resp := &obj
		return resp
	} else if animalName == "chickens" {
		obj := (*animals.Chicken).Init(new(animals.Chicken))
		resp := &obj
		return resp
	} else {
		obj := (*animals.Cow).Init(new(animals.Cow))
		resp := &obj
		return resp
	}
}

func (farm *Farm) GetProduction() {
	for i := 0; i <= 6; i++ {
		for breed, animals := range farm.Animals {
			for _, animal := range animals {
				production := "production"
				if breed == "cows" {
					production = "milk"
				} else if breed == "chickens" {
					production = "eggs"
				}
				farm.ProductionPerWeek[production] += animal.GetOutputProduct()
			}
		}
	}
}
