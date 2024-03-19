package main

import (
	"fmt"
	"slices"
)

func main() {
	animals := []string{
		"Dog",
		"Cat",
		"Lynx",
	}

	animals = append(animals, "Fish")
	fmt.Println("This is my slice", animals)

	animals = slices.Delete(animals, 3, 4)
	fmt.Println("This is my new slice", animals)

	for index, animal := range animals {
		fmt.Printf("this is my animal %s at %d\n", animal, index)
	}

	for value := range 10 {
		fmt.Printf("this is value %d\n", value)
	}

	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}

}
