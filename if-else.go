package main

import "fmt"

func main() {
	number := 10

	for i := 1; i <= number; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
