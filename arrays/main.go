package main

import (
    "fmt"
    "strings"
)

func reverse(ss []int) []int {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}

	return ss
}

func main() {
	var num int
	fmt.Scan(&num)

	s := make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Scan(&s[i])
	}
    
    s = reverse(s)
	fmt.Println(strings.Trim(fmt.Sprint(s), "[]"))
}