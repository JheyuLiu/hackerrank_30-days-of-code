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
	si, _ := strconv.ParseUint(scanner.Text(), 10, 64)

	scanner.Scan()
	sd, _ := strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	ss := scanner.Text()

	fmt.Println(i+si)
	fmt.Printf("%.1f\n", d+sd)
	fmt.Println(s+ss)
}
