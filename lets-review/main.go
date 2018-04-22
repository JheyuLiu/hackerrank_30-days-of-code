package main

import (
	"fmt"
	"bufio"
	"os"
)

func print_string(str string) {
	var odd_str string
	for i, c := range(str) {
		if i%2 == 0 {
			fmt.Printf("%c", c)
		} else {
			odd_str += string(c)
		}
	}
    fmt.Println("", odd_str)
}

func main() {
	var num int
	fmt.Scan(&num)

    scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < num; i++ {
		scanner.Scan()
	    s := scanner.Text()
	    print_string(s)
	}
}