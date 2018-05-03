package main

import (
    "fmt"
    "errors"
    "math"
)

func check_number(n, p float64) (float64, error) {
    if (n < 0 || p < 0) {
    	return -1, errors.New("n and p should be non-negative")
    } else {
    	return math.Pow(n, p), nil
    }
}

func main() {
	var num int
	fmt.Scan(&num)

    var n, p float64
	for num > 0 {
		fmt.Scan(&n)
		fmt.Scan(&p)

        result, err := check_number(n, p)
        if err != nil {
        	fmt.Println(err)
        } else {
        	fmt.Println(result)
        }
        num--
	}
}