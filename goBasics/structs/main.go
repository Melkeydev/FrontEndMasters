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

func (p *Person) changeName(name string) {
	p.Name = name
}

func main() {
	newPerson := NewPerson("melkey", 27)
	fmt.Printf("this is my person %+v\n", newPerson)

	newPerson.changeName("marc")
	fmt.Printf("this is my person %+v\n", newPerson)

	fmt.Println("this is the memory address of my person", &newPerson)
}
