package main

import (
	"farm/farm"
	"fmt"
)

func main() {
	farm := new(farm.Farm)
	farm1 := farm.Init()
	farm1.AddAnimal("cows", 5)
	farm1.AddAnimal("chickens", 10)
	//fmt.Println(farm1.Animals)
	farm1.GetProduction()
	fmt.Println(farm1.ProductionPerWeek)
}
