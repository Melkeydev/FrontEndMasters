package main

import "fmt"

func main() {
	myMap := make(map[string]string)
	myMap["uuid"] = "1234"

	for key, value := range myMap {
		fmt.Printf("key: %s and value :%s\n", key, value)
	}

	myNestedMap := make(map[string]map[string]string)

	myNestedMap["firstMap"] = map[string]string{
		"nestedMap": "myFirstNestedMap",
	}

	fmt.Println(myNestedMap)
}
