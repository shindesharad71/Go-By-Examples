package main

import (
	"fmt"
	"time"
)

func main() {
	num := 10

	for i := 1; i < num; i++ {
		switch i {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		default:
			fmt.Println("Any Other Number")
		}
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its Weekend!")
	default:
		fmt.Println("Its Workday!")
	}

}
