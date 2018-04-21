package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)

	if num%2 == 0 {
		if num >= 2 && num <= 5 {
			fmt.Println("Not Weird")
		} else if num >= 6 && num <= 20 {
			fmt.Println("Weird")
		} else if num > 20 {
			fmt.Println("Not Weird")
		}
	} else {
		fmt.Println("Weird")
	}
}
