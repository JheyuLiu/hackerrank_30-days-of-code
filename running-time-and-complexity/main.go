package main

import "fmt"

func checkPrime(num int) string {
	if num <= 1 {
		return "Not prime"
	}
	if num <= 3 {
		return "Prime"
	}
	if num%2 == 0 || num%3 == 0 {
		return "Not prime"
	}
	for i := 5; i*i <= num; i=i+6 {
		if num%i == 0 || num%(i+2) == 0 {
			return "Not prime"
		}
	}
	return "Prime"
}

func main() {
	var num, val int
	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		fmt.Scan(&val)
		fmt.Println(checkPrime(val))
	}
}