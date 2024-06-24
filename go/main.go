package main

import (
	"farm/farm"
)

func main() {
	farm := new(farm.Farm)
	farm1 := farm.Init()
	farm1.AddAnimal("cow", 5)
	farm1.AddAnimal("chicken", 10)
	//fmt.Println(farm1.Animals)
	farm1.GetProduction()
}
