package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var si uint64
	var sd float64

	fmt.Scan(&si)
	fmt.Scan(&sd)

    scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ss := scanner.Text()

	fmt.Println(i + si)
	fmt.Printf("%.1f\n", d+sd)
	fmt.Println(s + ss)
}
