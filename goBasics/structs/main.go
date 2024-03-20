package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func (p *Person) changeName(newName string) {
	p.Name = newName
}

func main() {
	myPerson := NewPerson("Melkey", 26)
	myPerson.changeName("Marc")
	fmt.Printf("Ths is my person %+v\n", myPerson)

	mySlice := []int{1, 2, 3}

	for i := range mySlice {
		mySlice[i]++
	}

	fmt.Println("this is the slice", mySlice)
}
