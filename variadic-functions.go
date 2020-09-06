package main

import "fmt"

func addition(nums ...int) int {
	sum := 0
	for _, el := range nums {
		sum += el
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4}

	sum := addition(nums...)
	fmt.Println(sum)
}
