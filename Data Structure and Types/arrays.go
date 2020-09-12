package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arr [5]int
	fmt.Println("Empty Array:", arr)

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(50)
	}

	fmt.Println("Modified Array:", arr)
}
