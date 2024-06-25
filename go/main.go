package main

import (
	"farm/farm"
	"fmt"
)

func main() {
	farm := new(farm.Farm)
	farm = farm.Init()

	farm.AddAnimal("cows", 10)
	farm.AddAnimal("chickens", 20)

	farm.GetProduction()
	farm.GetCountAnimals()

	fmt.Println("Went to the market, bought animals")

	farm.AddAnimal("cows", 1)
	farm.AddAnimal("chickens", 5)

	farm.GetProduction()
	farm.GetCountAnimals()
}
