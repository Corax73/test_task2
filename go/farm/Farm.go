package farm

import (
	"farm/animals"
	"fmt"

	"golang.org/x/exp/maps"
)

type Farm struct {
	Animals           map[string][]animals.Animal
	ProductionPerWeek map[string]int
}

// Initializes Farm instance fields with zero values.
func (farm *Farm) Init() *Farm {
	return &Farm{
		Animals:           make(map[string][]animals.Animal),
		ProductionPerWeek: make(map[string]int),
	}
}

// Adds an Animal instance to a Farm instance.
func (farm *Farm) AddAnimal(animalName string, number int) {
	for i := 0; i < number; i++ {
		animalInstance := farm.CreateAnimal(animalName)
		farm.Animals[animalName] = append(farm.Animals[animalName], animalInstance)
	}
}

// Creates an Animal instance.
func (farm *Farm) CreateAnimal(animalName string) animals.Animal {
	switch animalName {
	case "cow":
		obj := (*animals.Cow).Init(new(animals.Cow))
		resp := &obj
		return resp
	case "chickens":
		obj := (*animals.Chicken).Init(new(animals.Chicken))
		resp := &obj
		return resp
	default:
		obj := (*animals.Cow).Init(new(animals.Cow))
		resp := &obj
		return resp
	}
}

// Receives products from the animals available on the Farm within 7 days.
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
	for production, quantity := range farm.ProductionPerWeek {
		fmt.Println(production, "collected", quantity)
	}
}

// Displays in the console the number of animals in the Farm by type.
func (farm *Farm) GetCountAnimals() {
	typesAnimals := maps.Keys(farm.Animals)
	for _, animalType := range typesAnimals {
		fmt.Println("Count", animalType, len(farm.Animals[animalType]))
	}
}
