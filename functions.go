package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func multiply(num1 int, num2 int) int {
	return num1 * num2
}

func main() {
	addition := sum(5, 6)
	fmt.Println("5+6 = ", addition)

	multiplication := multiply(5, 6)
	fmt.Println("5*6 = ", multiplication)
}
