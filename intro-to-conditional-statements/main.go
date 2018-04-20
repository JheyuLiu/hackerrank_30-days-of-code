package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num, _ := strconv.ParseInt(scanner.Text(), 10, 64)

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