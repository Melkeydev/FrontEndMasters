package main

import "fmt"

func main() {
	myMap := make(map[string]string)

	myMap["key"] = "value"

	for key, value := range myMap {
		fmt.Printf("this is key %s and this is value %s\n", key, value)
	}

	myNestedMap := map[string]map[string]string{}

	myNestedMap["firstMap"] = map[string]string{
		"nestedMap": "myFirstNestedMap",
	}

	fmt.Println(myNestedMap)

}
