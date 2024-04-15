package main

import "fmt"

const (
	PI = 3.14
)

func main() {
	myName := "Melkey"
	myInt := 10
	myFloat := 10.0

	fmt.Printf("Hello my name is %s my int is %d my float is %f\n", myName, myInt, myFloat)

	var myFriendsName string
	var myBool bool
	var myOtherInt int

	myFriendsName = "Prime"

	fmt.Printf("my other friends name %s my bool %t and myother int %d\n", myFriendsName, myBool, myOtherInt)
}
