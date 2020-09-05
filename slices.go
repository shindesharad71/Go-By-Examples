package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s := make([]int, 3)
	fmt.Println("Empty Slice - ", s)

	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(10)
	}

	fmt.Println("Modified Slice - ", s)
}
