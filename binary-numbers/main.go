package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
	var num int64
	fmt.Scan(&num)
    
	s := strconv.FormatInt(num, 2)
	sa := strings.Split(s, "0")
	
	var max int
	for _, str := range sa {
	    if len(str) > max {
	        max = len(str)
	    }
	}
	fmt.Println(max)
}
