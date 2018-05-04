package main

import (
	"fmt"
	"strconv"
)

func main() {
	var val string
	fmt.Scan(&val)

    num, err := strconv.ParseInt(val, 10, 0)
    if err != nil {
    	fmt.Println("Bad String")
    } else {
        fmt.Println(num)
    }
}