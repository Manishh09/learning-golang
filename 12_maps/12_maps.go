package main

import "fmt"

func main() {
	// create a map
	sampleMap := make(map[string]string)

	// assign values

	sampleMap["name"] = "Will"
	sampleMap["age"] = "22"
	sampleMap["profession"] = "sw eng"

	fmt.Println(sampleMap)

	delete(sampleMap, "age")

	fmt.Println(sampleMap)

	for key, value := range sampleMap {
		fmt.Println(key, value)
	}
}
