package main

import "fmt"

func main() {
	// Initilize Empty Map
	newMap := make(map[string]string)

	newMap["firstName"] = "Sharad"
	newMap["lastName"] = "Shinde"

	fmt.Println("Map - ", newMap)
	fmt.Println("Map Length - ", len(newMap))

	// Create Map with Initial Values
	mapWithValues := map[string]int{"boys": 5, "girls": 5}
	fmt.Println("Map - ", mapWithValues)
}
