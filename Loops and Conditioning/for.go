package main

import "fmt"

func main() {
	i := 1

	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}
}