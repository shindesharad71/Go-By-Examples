package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6}

	// Hint: Range works same as forin or forof loop in JS!
	for _, el := range arr {
		fmt.Println(el)
	}

	newMap := map[string]string{"firstName": "Sharad", "lastName": "Shinde"}

	for k, v := range newMap {
		fmt.Printf("%s => %s\n", k, v)
	}
}
