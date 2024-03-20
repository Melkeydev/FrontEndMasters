package main

import (
	"fmt"
	"slices"
)

func main() {
	animals := []string{
		"dog",
		"mouse",
	}

	animals = append(animals, "cat")
	animals = slices.Delete(animals, 2, 3)
	fmt.Println(animals)

	for index, animal := range animals {
		fmt.Printf("This is animal %s at %d\n", animal, index)
	}

	for value := range 10 {
		fmt.Println(value)
	}

	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}

}
