package main

import "fmt"

func nextNumber(num int) func() int {
	return func() int {
		fmt.Println(num + 1)
		return num + 1
	}
}

func main() {
	initialNumber := 1

	newNumber := nextNumber(initialNumber)
	fmt.Println(newNumber)
}
