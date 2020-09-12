package main

import "fmt"

func multipleReturns() (int, int) {
	return 5, 10
}

func main() {
	a, b := multipleReturns()
	fmt.Println(a, b)

	_, d := multipleReturns()
	fmt.Println(d)
}
