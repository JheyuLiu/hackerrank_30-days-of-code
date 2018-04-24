package main

import "fmt"

func factorial(num int) (ans int) {
	if num < 1 {
		return 1
	}

	ans = num * factorial(num-1)
	return
}

func main() {
	var num int
	fmt.Scan(&num)

	fmt.Println(factorial(num))
}